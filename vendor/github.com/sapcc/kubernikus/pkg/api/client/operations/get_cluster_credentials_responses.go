// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/sapcc/kubernikus/pkg/api/models"
)

// GetClusterCredentialsReader is a Reader for the GetClusterCredentials structure.
type GetClusterCredentialsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClusterCredentialsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetClusterCredentialsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetClusterCredentialsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetClusterCredentialsOK creates a GetClusterCredentialsOK with default headers values
func NewGetClusterCredentialsOK() *GetClusterCredentialsOK {
	return &GetClusterCredentialsOK{}
}

/*GetClusterCredentialsOK handles this case with default header values.

OK
*/
type GetClusterCredentialsOK struct {
	Payload *models.Credentials
}

func (o *GetClusterCredentialsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/clusters/{name}/credentials][%d] getClusterCredentialsOK  %+v", 200, o.Payload)
}

func (o *GetClusterCredentialsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Credentials)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClusterCredentialsDefault creates a GetClusterCredentialsDefault with default headers values
func NewGetClusterCredentialsDefault(code int) *GetClusterCredentialsDefault {
	return &GetClusterCredentialsDefault{
		_statusCode: code,
	}
}

/*GetClusterCredentialsDefault handles this case with default header values.

Error
*/
type GetClusterCredentialsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get cluster credentials default response
func (o *GetClusterCredentialsDefault) Code() int {
	return o._statusCode
}

func (o *GetClusterCredentialsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/clusters/{name}/credentials][%d] GetClusterCredentials default  %+v", o._statusCode, o.Payload)
}

func (o *GetClusterCredentialsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
