package defaults

import (
	"fmt"

	"github.com/cloudfoundry/cli/cf/api/resources"
	"github.com/cloudfoundry/cli/cf/configuration/core_config"
	"github.com/cloudfoundry/cli/cf/models"
	"github.com/cloudfoundry/cli/cf/net"
)

type DefaultSecurityGroupsRepoBase struct {
	ConfigRepo core_config.Reader
	Gateway    net.Gateway
}

func (repo *DefaultSecurityGroupsRepoBase) Bind(groupGuid string, path string) error {
	updatedPath := fmt.Sprintf("%s/%s", path, groupGuid)
	return repo.Gateway.UpdateResourceFromStruct(repo.ConfigRepo.ApiEndpoint(), updatedPath, "")
}

func (repo *DefaultSecurityGroupsRepoBase) List(path string) ([]models.SecurityGroupFields, error) {
	groups := []models.SecurityGroupFields{}

	err := repo.Gateway.ListPaginatedResources(
		repo.ConfigRepo.ApiEndpoint(),
		path,
		resources.SecurityGroupResource{},
		func(resource interface{}) bool {
			if securityGroupResource, ok := resource.(resources.SecurityGroupResource); ok {
				groups = append(groups, securityGroupResource.ToFields())
			}

			return true
		},
	)

	return groups, err
}

func (repo *DefaultSecurityGroupsRepoBase) Delete(groupGuid string, path string) error {
	updatedPath := fmt.Sprintf("%s/%s", path, groupGuid)
	return repo.Gateway.DeleteResource(repo.ConfigRepo.ApiEndpoint(), updatedPath)
}
