package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetDriverEndpointParams creates a new GetDriverEndpointParams object
// with the default values initialized.
func NewGetDriverEndpointParams() *GetDriverEndpointParams {
	var ()
	return &GetDriverEndpointParams{}
}

/*GetDriverEndpointParams contains all the parameters to send to the API endpoint
for the get driver endpoint operation typically these are written to a http.Request
*/
type GetDriverEndpointParams struct {

	/*DriverEndpointID
	  Driver Endpoint ID

	*/
	DriverEndpointID string
}

// WithDriverEndpointID adds the driverEndpointId to the get driver endpoint params
func (o *GetDriverEndpointParams) WithDriverEndpointID(DriverEndpointID string) *GetDriverEndpointParams {
	o.DriverEndpointID = DriverEndpointID
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetDriverEndpointParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	var res []error

	// path param driver_endpoint_id
	if err := r.SetPathParam("driver_endpoint_id", o.DriverEndpointID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}