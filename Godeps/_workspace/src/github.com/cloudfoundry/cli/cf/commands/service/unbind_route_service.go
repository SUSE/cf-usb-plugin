package service

import (
	"github.com/cloudfoundry/cli/cf/api"
	"github.com/cloudfoundry/cli/cf/command_registry"
	"github.com/cloudfoundry/cli/cf/configuration/core_config"
	"github.com/cloudfoundry/cli/cf/errors"
	. "github.com/cloudfoundry/cli/cf/i18n"
	"github.com/cloudfoundry/cli/cf/models"
	"github.com/cloudfoundry/cli/cf/requirements"
	"github.com/cloudfoundry/cli/cf/terminal"
	"github.com/cloudfoundry/cli/flags"
	"github.com/cloudfoundry/cli/flags/flag"
)

type RouteServiceUnbinder interface {
	UnbindRoute(route models.Route, serviceInstance models.ServiceInstance) error
}

type UnbindRouteService struct {
	ui                      terminal.UI
	config                  core_config.Reader
	routeRepo               api.RouteRepository
	routeServiceBindingRepo api.RouteServiceBindingRepository
	domainReq               requirements.DomainRequirement
	serviceInstanceReq      requirements.ServiceInstanceRequirement
}

func init() {
	command_registry.Register(&UnbindRouteService{})
}

func (cmd *UnbindRouteService) MetaData() command_registry.CommandMetadata {
	fs := make(map[string]flags.FlagSet)
	fs["n"] = &cliFlags.StringFlag{ShortName: "n", Usage: T("Hostname used in combination with DOMAIN to specify the route to unbind")}
	fs["f"] = &cliFlags.BoolFlag{ShortName: "f", Usage: T("Force unbinding without confirmation")}

	return command_registry.CommandMetadata{
		Name:        "unbind-route-service",
		ShortName:   "urs",
		Description: T("Unbind a service instance from a route"),
		Usage: T(`CF_NAME unbind-route-service DOMAIN SERVICE_INSTANCE [-n HOST] [-f]
		
EXAMPLE:
	CF_NAME unbind-route-service 10.244.0.34.xip.io myratelimiter -n spring-music`),
		Flags: fs,
	}
}

func (cmd *UnbindRouteService) Requirements(requirementsFactory requirements.Factory, fc flags.FlagContext) ([]requirements.Requirement, error) {
	if len(fc.Args()) != 2 {
		cmd.ui.Failed(T("Incorrect Usage. Requires DOMAIN and SERVICE_INSTANCE as arguments\n\n") + command_registry.Commands.CommandUsage("unbind-route-service"))
	}

	serviceName := fc.Args()[1]
	cmd.serviceInstanceReq = requirementsFactory.NewServiceInstanceRequirement(serviceName)

	domainName := fc.Args()[0]
	cmd.domainReq = requirementsFactory.NewDomainRequirement(domainName)

	return []requirements.Requirement{
		requirementsFactory.NewLoginRequirement(),
		cmd.domainReq,
		cmd.serviceInstanceReq,
	}, nil
}

func (cmd *UnbindRouteService) SetDependency(deps command_registry.Dependency, pluginCall bool) command_registry.Command {
	cmd.ui = deps.Ui
	cmd.config = deps.Config
	cmd.routeRepo = deps.RepoLocator.GetRouteRepository()
	cmd.routeServiceBindingRepo = deps.RepoLocator.GetRouteServiceBindingRepository()
	return cmd
}

func (cmd *UnbindRouteService) Execute(c flags.FlagContext) {
	host := c.String("n")
	domain := cmd.domainReq.GetDomain()
	path := "" // path is not currently supported

	route, err := cmd.routeRepo.Find(host, domain, path)
	if err != nil {
		cmd.ui.Failed(err.Error())
	}

	serviceInstance := cmd.serviceInstanceReq.GetServiceInstance()
	confirmed := c.Bool("f")
	if !confirmed {
		confirmed = cmd.ui.Confirm(T("Unbinding may leave apps mapped to route {{.URL}} vulnerable; e.g. if service instance {{.ServiceInstanceName}} provides authentication. Do you want to proceed?",
			map[string]interface{}{
				"URL": route.URL(),
				"ServiceInstanceName": serviceInstance.Name,
			}))

		if !confirmed {
			cmd.ui.Warn(T("Unbind cancelled"))
			return
		}
	}

	cmd.ui.Say(T("Unbinding route {{.URL}} from service instance {{.ServiceInstanceName}} in org {{.OrgName}} / space {{.SpaceName}} as {{.CurrentUser}}...",
		map[string]interface{}{
			"ServiceInstanceName": terminal.EntityNameColor(serviceInstance.Name),
			"URL":         terminal.EntityNameColor(route.URL()),
			"OrgName":     terminal.EntityNameColor(cmd.config.OrganizationFields().Name),
			"SpaceName":   terminal.EntityNameColor(cmd.config.SpaceFields().Name),
			"CurrentUser": terminal.EntityNameColor(cmd.config.Username()),
		}))

	err = cmd.UnbindRoute(route, serviceInstance)
	if err != nil {
		httpError, ok := err.(errors.HttpError)
		if ok && httpError.ErrorCode() == errors.ROUTE_WAS_NOT_BOUND {
			cmd.ui.Warn(T("Route {{.Route}} was not bound to service instance {{.ServiceInstance}}.", map[string]interface{}{"Route": route.URL(), "ServiceInstance": serviceInstance.Name}))
		} else {
			cmd.ui.Failed(err.Error())
		}
	}

	cmd.ui.Ok()
}

func (cmd *UnbindRouteService) UnbindRoute(route models.Route, serviceInstance models.ServiceInstance) error {
	return cmd.routeServiceBindingRepo.Unbind(serviceInstance.Guid, route.Guid, serviceInstance.IsUserProvided())
}
