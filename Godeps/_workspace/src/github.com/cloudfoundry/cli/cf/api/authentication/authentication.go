package authentication

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	. "github.com/cloudfoundry/cli/cf/i18n"

	"github.com/cloudfoundry/cli/cf/configuration/core_config"
	"github.com/cloudfoundry/cli/cf/errors"
	"github.com/cloudfoundry/cli/cf/net"
)

type TokenRefresher interface {
	RefreshAuthToken() (updatedToken string, apiErr error)
}

type AuthenticationRepository interface {
	RefreshAuthToken() (updatedToken string, apiErr error)
	Authenticate(credentials map[string]string) (apiErr error)
	GetLoginPromptsAndSaveUAAServerURL() (map[string]core_config.AuthPrompt, error)
}

type UAAAuthenticationRepository struct {
	config  core_config.ReadWriter
	gateway net.Gateway
}

func NewUAAAuthenticationRepository(gateway net.Gateway, config core_config.ReadWriter) (uaa UAAAuthenticationRepository) {
	uaa.gateway = gateway
	uaa.config = config
	return
}

func (uaa UAAAuthenticationRepository) Authenticate(credentials map[string]string) error {
	data := url.Values{
		"grant_type": {"password"},
		"scope":      {""},
	}
	for key, val := range credentials {
		data[key] = []string{val}
	}

	err := uaa.getAuthToken(data)
	if err != nil {
		httpError, ok := err.(errors.HttpError)
		if ok {
			switch {
			case httpError.StatusCode() == http.StatusUnauthorized:
				return errors.New(T("Credentials were rejected, please try again."))
			case httpError.StatusCode() >= http.StatusInternalServerError:
				return errors.New(T("The targeted API endpoint could not be reached."))
			}
		}

		return err
	}

	return nil
}

type LoginResource struct {
	Prompts map[string][]string
	Links   map[string]string
}

var knownAuthPromptTypes = map[string]core_config.AuthPromptType{
	"text":     core_config.AuthPromptTypeText,
	"password": core_config.AuthPromptTypePassword,
}

func (r *LoginResource) parsePrompts() (prompts map[string]core_config.AuthPrompt) {
	prompts = make(map[string]core_config.AuthPrompt)
	for key, val := range r.Prompts {
		prompts[key] = core_config.AuthPrompt{
			Type:        knownAuthPromptTypes[val[0]],
			DisplayName: val[1],
		}
	}
	return
}

func (uaa UAAAuthenticationRepository) GetLoginPromptsAndSaveUAAServerURL() (prompts map[string]core_config.AuthPrompt, apiErr error) {
	url := fmt.Sprintf("%s/login", uaa.config.AuthenticationEndpoint())
	resource := &LoginResource{}
	apiErr = uaa.gateway.GetResource(url, resource)

	prompts = resource.parsePrompts()
	if resource.Links["uaa"] == "" {
		uaa.config.SetUaaEndpoint(uaa.config.AuthenticationEndpoint())
	} else {
		uaa.config.SetUaaEndpoint(resource.Links["uaa"])
	}
	return
}

func (uaa UAAAuthenticationRepository) RefreshAuthToken() (string, error) {
	data := url.Values{
		"refresh_token": {uaa.config.RefreshToken()},
		"grant_type":    {"refresh_token"},
		"scope":         {""},
	}

	apiErr := uaa.getAuthToken(data)
	updatedToken := uaa.config.AccessToken()

	return updatedToken, apiErr
}

func (uaa UAAAuthenticationRepository) getAuthToken(data url.Values) error {
	type uaaErrorResponse struct {
		Code        string `json:"error"`
		Description string `json:"error_description"`
	}

	type AuthenticationResponse struct {
		AccessToken  string           `json:"access_token"`
		TokenType    string           `json:"token_type"`
		RefreshToken string           `json:"refresh_token"`
		Error        uaaErrorResponse `json:"error"`
	}

	path := fmt.Sprintf("%s/oauth/token", uaa.config.AuthenticationEndpoint())
	request, err := uaa.gateway.NewRequest("POST", path, "Basic "+base64.StdEncoding.EncodeToString([]byte("cf:")), strings.NewReader(data.Encode()))
	if err != nil {
		return fmt.Errorf("%s: %s", T("Failed to start oauth request"), err.Error())
	}
	request.HttpReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response := new(AuthenticationResponse)
	_, err = uaa.gateway.PerformRequestForJSONResponse(request, &response)

	switch err.(type) {
	case nil:
	case errors.HttpError:
		return err
	case *errors.InvalidTokenError:
		return errors.New(T("Authentication has expired.  Please log back in to re-authenticate.\n\nTIP: Use `cf login -a <endpoint> -u <user> -o <org> -s <space>` to log back in and re-authenticate."))
	default:
		return fmt.Errorf("%s: %s", T("auth request failed"), err.Error())
	}

	// TODO: get the actual status code
	if response.Error.Code != "" {
		return errors.NewHttpError(0, response.Error.Code, response.Error.Description)
	}

	uaa.config.SetAccessToken(fmt.Sprintf("%s %s", response.TokenType, response.AccessToken))
	uaa.config.SetRefreshToken(response.RefreshToken)

	return nil
}
