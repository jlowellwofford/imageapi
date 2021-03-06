// Code generated by go-swagger; DO NOT EDIT.

package mounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kraken-hpc/imageapi/models"
)

// MountRbdReader is a Reader for the MountRbd structure.
type MountRbdReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MountRbdReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewMountRbdCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewMountRbdDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewMountRbdCreated creates a MountRbdCreated with default headers values
func NewMountRbdCreated() *MountRbdCreated {
	return &MountRbdCreated{}
}

/*MountRbdCreated handles this case with default header values.

RBD mount succeed
*/
type MountRbdCreated struct {
	Payload *models.MountRbd
}

func (o *MountRbdCreated) Error() string {
	return fmt.Sprintf("[POST /mount/rbd][%d] mountRbdCreated  %+v", 201, o.Payload)
}

func (o *MountRbdCreated) GetPayload() *models.MountRbd {
	return o.Payload
}

func (o *MountRbdCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MountRbd)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMountRbdDefault creates a MountRbdDefault with default headers values
func NewMountRbdDefault(code int) *MountRbdDefault {
	return &MountRbdDefault{
		_statusCode: code,
	}
}

/*MountRbdDefault handles this case with default header values.

error
*/
type MountRbdDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the mount rbd default response
func (o *MountRbdDefault) Code() int {
	return o._statusCode
}

func (o *MountRbdDefault) Error() string {
	return fmt.Sprintf("[POST /mount/rbd][%d] mount_rbd default  %+v", o._statusCode, o.Payload)
}

func (o *MountRbdDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *MountRbdDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
