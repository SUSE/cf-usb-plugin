package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hpcloud/cf-plugin-usb/lib/models"
)

// NewUpdateDriverEndpointParams creates a new UpdateDriverEndpointParams object
// with the default values initialized.
func NewUpdateDriverEndpointParams() *UpdateDriverEndpointParams {
	var ()
	return &UpdateDriverEndpointParams{}
}

/*UpdateDriverEndpointParams contains all the parameters to send to the API endpoint
for the update driver endpoint operation typically these are written to a http.Request
*/
type UpdateDriverEndpointParams struct {

	/*DriverEndpoint
	  Updated information for an already registered driver endpoint


	*/
	DriverEndpoint *models.DriverEndpoint
	/*DriverEndpointID
	  Driver Endpoint ID


	*/
	DriverEndpointID string
}

// WithDriverEndpoint adds the driverEndpoint to the update driver endpoint params
func (o *UpdateDriverEndpointParams) WithDriverEndpoint(DriverEndpoint *models.DriverEndpoint) *UpdateDriverEndpointParams {
	o.DriverEndpoint = DriverEndpoint
	return o
}

// WithDriverEndpointID adds the driverEndpointId to the update driver endpoint params
func (o *UpdateDriverEndpointParams) WithDriverEndpointID(DriverEndpointID string) *UpdateDriverEndpointParams {
	o.DriverEndpointID = DriverEndpointID
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateDriverEndpointParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	var res []error

	if o.DriverEndpoint == nil {
		o.DriverEndpoint = new(models.DriverEndpoint)
	}

	if err := r.SetBodyParam(o.DriverEndpoint); err != nil {
		return err
	}

	// path param driver_endpoint_id
	if err := r.SetPathParam("driver_endpoint_id", o.DriverEndpointID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
