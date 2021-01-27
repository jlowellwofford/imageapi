// Code generated by go-swagger; DO NOT EDIT.

package attach

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/jlowellwofford/imageapi/models"
)

// MapRbdCreatedCode is the HTTP code returned for type MapRbdCreated
const MapRbdCreatedCode int = 201

/*MapRbdCreated RBD attach succeed

swagger:response mapRbdCreated
*/
type MapRbdCreated struct {

	/*
	  In: Body
	*/
	Payload *models.Rbd `json:"body,omitempty"`
}

// NewMapRbdCreated creates MapRbdCreated with default headers values
func NewMapRbdCreated() *MapRbdCreated {

	return &MapRbdCreated{}
}

// WithPayload adds the payload to the map rbd created response
func (o *MapRbdCreated) WithPayload(payload *models.Rbd) *MapRbdCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the map rbd created response
func (o *MapRbdCreated) SetPayload(payload *models.Rbd) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *MapRbdCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*MapRbdDefault error

swagger:response mapRbdDefault
*/
type MapRbdDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewMapRbdDefault creates MapRbdDefault with default headers values
func NewMapRbdDefault(code int) *MapRbdDefault {
	if code <= 0 {
		code = 500
	}

	return &MapRbdDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the map rbd default response
func (o *MapRbdDefault) WithStatusCode(code int) *MapRbdDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the map rbd default response
func (o *MapRbdDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the map rbd default response
func (o *MapRbdDefault) WithPayload(payload *models.Error) *MapRbdDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the map rbd default response
func (o *MapRbdDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *MapRbdDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
