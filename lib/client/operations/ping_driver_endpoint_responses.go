package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PingDriverEndpointReader is a Reader for the PingDriverEndpoint structure.
type PingDriverEndpointReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *PingDriverEndpointReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPingDriverEndpointOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewPingDriverEndpointNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPingDriverEndpointInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPingDriverEndpointOK creates a PingDriverEndpointOK with default headers values
func NewPingDriverEndpointOK() *PingDriverEndpointOK {
	return &PingDriverEndpointOK{}
}

/*PingDriverEndpointOK handles this case with default header values.

OK
*/
type PingDriverEndpointOK struct {
}

func (o *PingDriverEndpointOK) Error() string {
	return fmt.Sprintf("[GET /driver_endpoint/{driver_endpoint_id}/ping][%d] pingDriverEndpointOK ", 200)
}

func (o *PingDriverEndpointOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPingDriverEndpointNotFound creates a PingDriverEndpointNotFound with default headers values
func NewPingDriverEndpointNotFound() *PingDriverEndpointNotFound {
	return &PingDriverEndpointNotFound{}
}

/*PingDriverEndpointNotFound handles this case with default header values.

Not Found
*/
type PingDriverEndpointNotFound struct {
}

func (o *PingDriverEndpointNotFound) Error() string {
	return fmt.Sprintf("[GET /driver_endpoint/{driver_endpoint_id}/ping][%d] pingDriverEndpointNotFound ", 404)
}

func (o *PingDriverEndpointNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPingDriverEndpointInternalServerError creates a PingDriverEndpointInternalServerError with default headers values
func NewPingDriverEndpointInternalServerError() *PingDriverEndpointInternalServerError {
	return &PingDriverEndpointInternalServerError{}
}

/*PingDriverEndpointInternalServerError handles this case with default header values.

Unexpected error
*/
type PingDriverEndpointInternalServerError struct {
	Payload string
}

func (o *PingDriverEndpointInternalServerError) Error() string {
	return fmt.Sprintf("[GET /driver_endpoint/{driver_endpoint_id}/ping][%d] pingDriverEndpointInternalServerError  %+v", 500, o.Payload)
}

func (o *PingDriverEndpointInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}