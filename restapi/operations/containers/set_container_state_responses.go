// Code generated by go-swagger; DO NOT EDIT.

package containers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/jlowellwofford/imageapi/models"
)

// SetContainerStateOKCode is the HTTP code returned for type SetContainerStateOK
const SetContainerStateOKCode int = 200

/*SetContainerStateOK Container state changed

swagger:response setContainerStateOK
*/
type SetContainerStateOK struct {

	/*
	  In: Body
	*/
	Payload *models.Container `json:"body,omitempty"`
}

// NewSetContainerStateOK creates SetContainerStateOK with default headers values
func NewSetContainerStateOK() *SetContainerStateOK {

	return &SetContainerStateOK{}
}

// WithPayload adds the payload to the set container state o k response
func (o *SetContainerStateOK) WithPayload(payload *models.Container) *SetContainerStateOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the set container state o k response
func (o *SetContainerStateOK) SetPayload(payload *models.Container) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SetContainerStateOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*SetContainerStateDefault error

swagger:response setContainerStateDefault
*/
type SetContainerStateDefault struct {
	_statusCode int
}

// NewSetContainerStateDefault creates SetContainerStateDefault with default headers values
func NewSetContainerStateDefault(code int) *SetContainerStateDefault {
	if code <= 0 {
		code = 500
	}

	return &SetContainerStateDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the set container state default response
func (o *SetContainerStateDefault) WithStatusCode(code int) *SetContainerStateDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the set container state default response
func (o *SetContainerStateDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WriteResponse to the client
func (o *SetContainerStateDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(o._statusCode)
}
