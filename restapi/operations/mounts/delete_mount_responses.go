// Code generated by go-swagger; DO NOT EDIT.

package mounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/jlowellwofford/imageapi/models"
)

// DeleteMountNoContentCode is the HTTP code returned for type DeleteMountNoContent
const DeleteMountNoContentCode int = 204

/*DeleteMountNoContent Unmount succeeded

swagger:response deleteMountNoContent
*/
type DeleteMountNoContent struct {

	/*
	  In: Body
	*/
	Payload *models.Mount `json:"body,omitempty"`
}

// NewDeleteMountNoContent creates DeleteMountNoContent with default headers values
func NewDeleteMountNoContent() *DeleteMountNoContent {

	return &DeleteMountNoContent{}
}

// WithPayload adds the payload to the delete mount no content response
func (o *DeleteMountNoContent) WithPayload(payload *models.Mount) *DeleteMountNoContent {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete mount no content response
func (o *DeleteMountNoContent) SetPayload(payload *models.Mount) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteMountNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(204)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*DeleteMountDefault Unmount failed

swagger:response deleteMountDefault
*/
type DeleteMountDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteMountDefault creates DeleteMountDefault with default headers values
func NewDeleteMountDefault(code int) *DeleteMountDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteMountDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete mount default response
func (o *DeleteMountDefault) WithStatusCode(code int) *DeleteMountDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete mount default response
func (o *DeleteMountDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the delete mount default response
func (o *DeleteMountDefault) WithPayload(payload *models.Error) *DeleteMountDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete mount default response
func (o *DeleteMountDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteMountDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
