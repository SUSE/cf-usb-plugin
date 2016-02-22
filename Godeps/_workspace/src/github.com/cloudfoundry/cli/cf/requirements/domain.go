package requirements

import (
	"github.com/cloudfoundry/cli/cf/api"
	"github.com/cloudfoundry/cli/cf/configuration/core_config"
	"github.com/cloudfoundry/cli/cf/models"
	"github.com/cloudfoundry/cli/cf/terminal"
)

//go:generate counterfeiter -o fakes/fake_domain_requirement.go . DomainRequirement
type DomainRequirement interface {
	Requirement
	GetDomain() models.DomainFields
}

type domainApiRequirement struct {
	name       string
	ui         terminal.UI
	config     core_config.Reader
	domainRepo api.DomainRepository
	domain     models.DomainFields
}

func NewDomainRequirement(name string, ui terminal.UI, config core_config.Reader, domainRepo api.DomainRepository) (req *domainApiRequirement) {
	req = new(domainApiRequirement)
	req.name = name
	req.ui = ui
	req.config = config
	req.domainRepo = domainRepo
	return
}

func (req *domainApiRequirement) Execute() bool {
	var apiErr error
	req.domain, apiErr = req.domainRepo.FindByNameInOrg(req.name, req.config.OrganizationFields().Guid)

	if apiErr != nil {
		req.ui.Failed(apiErr.Error())
		return false
	}

	return true
}

func (req *domainApiRequirement) GetDomain() models.DomainFields {
	return req.domain
}
