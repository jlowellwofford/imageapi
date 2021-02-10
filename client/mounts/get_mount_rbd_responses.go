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

// GetMountRbdReader is a Reader for the GetMountRbd structure.
type GetMountRbdReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMountRbdReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMountRbdOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetMountRbdDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMountRbdOK creates a GetMountRbdOK with default headers values
func NewGetMountRbdOK() *GetMountRbdOK {
	return &GetMountRbdOK{}
}

/*GetMountRbdOK handles this case with default header values.

RBD mount entry
*/
type GetMountRbdOK struct {
	Payload *models.MountRbd
}

func (o *GetMountRbdOK) Error() string {
	return fmt.Sprintf("[GET /mount/rbd/{id}][%d] getMountRbdOK  %+v", 200, o.Payload)
}

func (o *GetMountRbdOK) GetPayload() *models.MountRbd {
	return o.Payload
}

func (o *GetMountRbdOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MountRbd)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMountRbdDefault creates a GetMountRbdDefault with default headers values
func NewGetMountRbdDefault(code int) *GetMountRbdDefault {
	return &GetMountRbdDefault{
		_statusCode: code,
	}
}

/*GetMountRbdDefault handles this case with default header values.

error
*/
type GetMountRbdDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get mount rbd default response
func (o *GetMountRbdDefault) Code() int {
	return o._statusCode
}

func (o *GetMountRbdDefault) Error() string {
	return fmt.Sprintf("[GET /mount/rbd/{id}][%d] get_mount_rbd default  %+v", o._statusCode, o.Payload)
}

func (o *GetMountRbdDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetMountRbdDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
