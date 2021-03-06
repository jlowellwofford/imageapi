// Code generated by go-swagger; DO NOT EDIT.

package mounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// MountOverlayHandlerFunc turns a function with the right signature into a mount overlay handler
type MountOverlayHandlerFunc func(MountOverlayParams) middleware.Responder

// Handle executing the request and returning a response
func (fn MountOverlayHandlerFunc) Handle(params MountOverlayParams) middleware.Responder {
	return fn(params)
}

// MountOverlayHandler interface for that can handle valid mount overlay params
type MountOverlayHandler interface {
	Handle(MountOverlayParams) middleware.Responder
}

// NewMountOverlay creates a new http.Handler for the mount overlay operation
func NewMountOverlay(ctx *middleware.Context, handler MountOverlayHandler) *MountOverlay {
	return &MountOverlay{Context: ctx, Handler: handler}
}

/*MountOverlay swagger:route POST /mount/overlay mounts mountOverlay

MountOverlay mount overlay API

*/
type MountOverlay struct {
	Context *middleware.Context
	Handler MountOverlayHandler
}

func (o *MountOverlay) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewMountOverlayParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
