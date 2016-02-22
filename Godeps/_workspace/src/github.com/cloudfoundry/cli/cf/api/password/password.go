package password

import (
	"fmt"
	"strings"

	"github.com/cloudfoundry/cli/cf/configuration/core_config"
	"github.com/cloudfoundry/cli/cf/errors"
	. "github.com/cloudfoundry/cli/cf/i18n"
	"github.com/cloudfoundry/cli/cf/net"
)

type PasswordRepository interface {
	UpdatePassword(old string, new string) error
}

type CloudControllerPasswordRepository struct {
	config  core_config.Reader
	gateway net.Gateway
}

func NewCloudControllerPasswordRepository(config core_config.Reader, gateway net.Gateway) (repo CloudControllerPasswordRepository) {
	repo.config = config
	repo.gateway = gateway
	return
}

func (repo CloudControllerPasswordRepository) UpdatePassword(old string, new string) error {
	uaaEndpoint := repo.config.UaaEndpoint()
	if uaaEndpoint == "" {
		return errors.New(T("UAA endpoint missing from config file"))
	}

	url := fmt.Sprintf("/Users/%s/password", repo.config.UserGuid())
	body := fmt.Sprintf(`{"password":"%s","oldPassword":"%s"}`, new, old)

	return repo.gateway.UpdateResource(uaaEndpoint, url, strings.NewReader(body))
}
