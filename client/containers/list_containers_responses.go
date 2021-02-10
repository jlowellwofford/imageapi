// Code generated by go-swagger; DO NOT EDIT.

package containers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/jlowellwofford/imageapi/models"
)

// ListContainersReader is a Reader for the ListContainers structure.
type ListContainersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListContainersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListContainersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListContainersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListContainersOK creates a ListContainersOK with default headers values
func NewListContainersOK() *ListContainersOK {
	return &ListContainersOK{}
}

/*ListContainersOK handles this case with default header values.

List of containers
*/
type ListContainersOK struct {
	Payload []*models.Container
}

func (o *ListContainersOK) Error() string {
	return fmt.Sprintf("[GET /container][%d] listContainersOK  %+v", 200, o.Payload)
}

func (o *ListContainersOK) GetPayload() []*models.Container {
	return o.Payload
}

func (o *ListContainersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListContainersDefault creates a ListContainersDefault with default headers values
func NewListContainersDefault(code int) *ListContainersDefault {
	return &ListContainersDefault{
		_statusCode: code,
	}
}

/*ListContainersDefault handles this case with default header values.

error
*/
type ListContainersDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the list containers default response
func (o *ListContainersDefault) Code() int {
	return o._statusCode
}

func (o *ListContainersDefault) Error() string {
	return fmt.Sprintf("[GET /container][%d] list_containers default  %+v", o._statusCode, o.Payload)
}

func (o *ListContainersDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListContainersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
