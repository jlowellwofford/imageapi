// Code generated by go-swagger; DO NOT EDIT.

package mounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteMountHandlerFunc turns a function with the right signature into a delete mount handler
type DeleteMountHandlerFunc func(DeleteMountParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteMountHandlerFunc) Handle(params DeleteMountParams) middleware.Responder {
	return fn(params)
}

// DeleteMountHandler interface for that can handle valid delete mount params
type DeleteMountHandler interface {
	Handle(DeleteMountParams) middleware.Responder
}

// NewDeleteMount creates a new http.Handler for the delete mount operation
func NewDeleteMount(ctx *middleware.Context, handler DeleteMountHandler) *DeleteMount {
	return &DeleteMount{Context: ctx, Handler: handler}
}

/*DeleteMount swagger:route DELETE /mount mounts deleteMount

Unmount a specified mount.  Note that mount reference IDs must be specified.

*/
type DeleteMount struct {
	Context *middleware.Context
	Handler DeleteMountHandler
}

func (o *DeleteMount) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteMountParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
