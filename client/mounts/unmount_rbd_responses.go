// Code generated by go-swagger; DO NOT EDIT.

package mounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/jlowellwofford/imageapi/models"
)

// UnmountRbdReader is a Reader for the UnmountRbd structure.
type UnmountRbdReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UnmountRbdReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUnmountRbdNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUnmountRbdDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUnmountRbdNoContent creates a UnmountRbdNoContent with default headers values
func NewUnmountRbdNoContent() *UnmountRbdNoContent {
	return &UnmountRbdNoContent{}
}

/*UnmountRbdNoContent handles this case with default header values.

Unmounted
*/
type UnmountRbdNoContent struct {
}

func (o *UnmountRbdNoContent) Error() string {
	return fmt.Sprintf("[DELETE /mount/rbd/{id}][%d] unmountRbdNoContent ", 204)
}

func (o *UnmountRbdNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUnmountRbdDefault creates a UnmountRbdDefault with default headers values
func NewUnmountRbdDefault(code int) *UnmountRbdDefault {
	return &UnmountRbdDefault{
		_statusCode: code,
	}
}

/*UnmountRbdDefault handles this case with default header values.

error
*/
type UnmountRbdDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the unmount rbd default response
func (o *UnmountRbdDefault) Code() int {
	return o._statusCode
}

func (o *UnmountRbdDefault) Error() string {
	return fmt.Sprintf("[DELETE /mount/rbd/{id}][%d] unmount_rbd default  %+v", o._statusCode, o.Payload)
}

func (o *UnmountRbdDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *UnmountRbdDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}