package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*DriverEndpoint driver endpoint

swagger:model driverEndpoint
*/
type DriverEndpoint struct {

	/* An authentication key used by the USB when communicating with the
	driver endpoint.

	*/
	AuthenticationKey string `json:"authenticationKey,omitempty"`

	/* The certificate used to issue the certificate providing TLS

	 */
	CaCertificate string `json:"caCertificate,omitempty"`

	/* URL for the driver endpoint. Used by the USB to create service
	instances, generate credentials, discover plans and schemas.

	*/
	EndpointURL string `json:"endpointURL,omitempty"`

	/* USB generated ID for the driver endpoint.

	 */
	ID string `json:"id,omitempty"`

	/* metadata
	 */
	Metadata *EndpointMetadata `json:"metadata,omitempty"`

	/* The name of the driver endpoint. It's displayed by the Cloud Foundry
	CLI when the user lists available service offerings.


	Required: true
	*/
	Name *string `json:"name"`

	/* Indicates if SSL validation is skiped for a specified driver endpoint

	 */
	SkipSSLValidation *bool `json:"skipSSLValidation,omitempty"`
}

// Validate validates this driver endpoint
func (m *DriverEndpoint) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMetadata(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DriverEndpoint) validateMetadata(formats strfmt.Registry) error {

	if swag.IsZero(m.Metadata) { // not required
		return nil
	}

	if m.Metadata != nil {

		if err := m.Metadata.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *DriverEndpoint) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}
