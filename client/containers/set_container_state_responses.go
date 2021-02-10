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

// SetContainerStateReader is a Reader for the SetContainerState structure.
type SetContainerStateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetContainerStateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSetContainerStateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSetContainerStateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSetContainerStateOK creates a SetContainerStateOK with default headers values
func NewSetContainerStateOK() *SetContainerStateOK {
	return &SetContainerStateOK{}
}

/*SetContainerStateOK handles this case with default header values.

Container state changed
*/
type SetContainerStateOK struct {
	Payload *models.Container
}

func (o *SetContainerStateOK) Error() string {
	return fmt.Sprintf("[GET /container/{id}/{state}][%d] setContainerStateOK  %+v", 200, o.Payload)
}

func (o *SetContainerStateOK) GetPayload() *models.Container {
	return o.Payload
}

func (o *SetContainerStateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Container)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetContainerStateDefault creates a SetContainerStateDefault with default headers values
func NewSetContainerStateDefault(code int) *SetContainerStateDefault {
	return &SetContainerStateDefault{
		_statusCode: code,
	}
}

/*SetContainerStateDefault handles this case with default header values.

error
*/
type SetContainerStateDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the set container state default response
func (o *SetContainerStateDefault) Code() int {
	return o._statusCode
}

func (o *SetContainerStateDefault) Error() string {
	return fmt.Sprintf("[GET /container/{id}/{state}][%d] set_container_state default  %+v", o._statusCode, o.Payload)
}

func (o *SetContainerStateDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *SetContainerStateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
