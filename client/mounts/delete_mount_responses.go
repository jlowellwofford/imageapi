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

// DeleteMountReader is a Reader for the DeleteMount structure.
type DeleteMountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteMountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteMountNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteMountDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteMountNoContent creates a DeleteMountNoContent with default headers values
func NewDeleteMountNoContent() *DeleteMountNoContent {
	return &DeleteMountNoContent{}
}

/*DeleteMountNoContent handles this case with default header values.

Unmount succeeded
*/
type DeleteMountNoContent struct {
	Payload *models.Mount
}

func (o *DeleteMountNoContent) Error() string {
	return fmt.Sprintf("[DELETE /mount][%d] deleteMountNoContent  %+v", 204, o.Payload)
}

func (o *DeleteMountNoContent) GetPayload() *models.Mount {
	return o.Payload
}

func (o *DeleteMountNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Mount)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteMountDefault creates a DeleteMountDefault with default headers values
func NewDeleteMountDefault(code int) *DeleteMountDefault {
	return &DeleteMountDefault{
		_statusCode: code,
	}
}

/*DeleteMountDefault handles this case with default header values.

Unmount failed
*/
type DeleteMountDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete mount default response
func (o *DeleteMountDefault) Code() int {
	return o._statusCode
}

func (o *DeleteMountDefault) Error() string {
	return fmt.Sprintf("[DELETE /mount][%d] DeleteMount default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteMountDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteMountDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
