// Code generated by go-swagger; DO NOT EDIT.

package containers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/jlowellwofford/imageapi/models"
)

// GetContainerOKCode is the HTTP code returned for type GetContainerOK
const GetContainerOKCode int = 200

/*GetContainerOK Container entry

swagger:response getContainerOK
*/
type GetContainerOK struct {

	/*
	  In: Body
	*/
	Payload *models.Container `json:"body,omitempty"`
}

// NewGetContainerOK creates GetContainerOK with default headers values
func NewGetContainerOK() *GetContainerOK {

	return &GetContainerOK{}
}

// WithPayload adds the payload to the get container o k response
func (o *GetContainerOK) WithPayload(payload *models.Container) *GetContainerOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get container o k response
func (o *GetContainerOK) SetPayload(payload *models.Container) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetContainerOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetContainerDefault error

swagger:response getContainerDefault
*/
type GetContainerDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetContainerDefault creates GetContainerDefault with default headers values
func NewGetContainerDefault(code int) *GetContainerDefault {
	if code <= 0 {
		code = 500
	}

	return &GetContainerDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get container default response
func (o *GetContainerDefault) WithStatusCode(code int) *GetContainerDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get container default response
func (o *GetContainerDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get container default response
func (o *GetContainerDefault) WithPayload(payload *models.Error) *GetContainerDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get container default response
func (o *GetContainerDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetContainerDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
