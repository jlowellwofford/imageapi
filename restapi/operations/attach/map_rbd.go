// Code generated by go-swagger; DO NOT EDIT.

package attach

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// MapRbdHandlerFunc turns a function with the right signature into a map rbd handler
type MapRbdHandlerFunc func(MapRbdParams) middleware.Responder

// Handle executing the request and returning a response
func (fn MapRbdHandlerFunc) Handle(params MapRbdParams) middleware.Responder {
	return fn(params)
}

// MapRbdHandler interface for that can handle valid map rbd params
type MapRbdHandler interface {
	Handle(MapRbdParams) middleware.Responder
}

// NewMapRbd creates a new http.Handler for the map rbd operation
func NewMapRbd(ctx *middleware.Context, handler MapRbdHandler) *MapRbd {
	return &MapRbd{Context: ctx, Handler: handler}
}

/*MapRbd swagger:route POST /attach/rbd attach mapRbd

MapRbd map rbd API

*/
type MapRbd struct {
	Context *middleware.Context
	Handler MapRbdHandler
}

func (o *MapRbd) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewMapRbdParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
