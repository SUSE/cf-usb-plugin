package quotas

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/cloudfoundry/cli/cf/api/resources"
	"github.com/cloudfoundry/cli/cf/configuration/core_config"
	"github.com/cloudfoundry/cli/cf/errors"
	"github.com/cloudfoundry/cli/cf/models"
	"github.com/cloudfoundry/cli/cf/net"
)

type QuotaRepository interface {
	FindAll() (quotas []models.QuotaFields, apiErr error)
	FindByName(name string) (quota models.QuotaFields, apiErr error)

	AssignQuotaToOrg(orgGuid, quotaGuid string) error

	// CRUD ahoy
	Create(quota models.QuotaFields) error
	Update(quota models.QuotaFields) error
	Delete(quotaGuid string) error
}

type CloudControllerQuotaRepository struct {
	config  core_config.Reader
	gateway net.Gateway
}

func NewCloudControllerQuotaRepository(config core_config.Reader, gateway net.Gateway) (repo CloudControllerQuotaRepository) {
	repo.config = config
	repo.gateway = gateway
	return
}

func (repo CloudControllerQuotaRepository) findAllWithPath(path string) ([]models.QuotaFields, error) {
	var quotas []models.QuotaFields
	apiErr := repo.gateway.ListPaginatedResources(
		repo.config.ApiEndpoint(),
		path,
		resources.QuotaResource{},
		func(resource interface{}) bool {
			if qr, ok := resource.(resources.QuotaResource); ok {
				quotas = append(quotas, qr.ToFields())
			}
			return true
		})
	return quotas, apiErr
}

func (repo CloudControllerQuotaRepository) FindAll() (quotas []models.QuotaFields, apiErr error) {
	return repo.findAllWithPath("/v2/quota_definitions")
}

func (repo CloudControllerQuotaRepository) FindByName(name string) (quota models.QuotaFields, apiErr error) {
	path := fmt.Sprintf("/v2/quota_definitions?q=%s", url.QueryEscape("name:"+name))
	quotas, apiErr := repo.findAllWithPath(path)
	if apiErr != nil {
		return
	}

	if len(quotas) == 0 {
		apiErr = errors.NewModelNotFoundError("Quota", name)
		return
	}

	quota = quotas[0]
	return
}

func (repo CloudControllerQuotaRepository) Create(quota models.QuotaFields) error {
	return repo.gateway.CreateResourceFromStruct(repo.config.ApiEndpoint(), "/v2/quota_definitions", quota)
}

func (repo CloudControllerQuotaRepository) Update(quota models.QuotaFields) error {
	path := fmt.Sprintf("/v2/quota_definitions/%s", quota.Guid)
	return repo.gateway.UpdateResourceFromStruct(repo.config.ApiEndpoint(), path, quota)
}

func (repo CloudControllerQuotaRepository) AssignQuotaToOrg(orgGuid, quotaGuid string) (apiErr error) {
	path := fmt.Sprintf("/v2/organizations/%s", orgGuid)
	data := fmt.Sprintf(`{"quota_definition_guid":"%s"}`, quotaGuid)
	return repo.gateway.UpdateResource(repo.config.ApiEndpoint(), path, strings.NewReader(data))
}

func (repo CloudControllerQuotaRepository) Delete(quotaGuid string) (apiErr error) {
	path := fmt.Sprintf("/v2/quota_definitions/%s", quotaGuid)
	return repo.gateway.DeleteResource(repo.config.ApiEndpoint(), path)
}
