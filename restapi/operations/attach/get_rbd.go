// Code generated by go-swagger; DO NOT EDIT.

package attach

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetRbdHandlerFunc turns a function with the right signature into a get rbd handler
type GetRbdHandlerFunc func(GetRbdParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetRbdHandlerFunc) Handle(params GetRbdParams) middleware.Responder {
	return fn(params)
}

// GetRbdHandler interface for that can handle valid get rbd params
type GetRbdHandler interface {
	Handle(GetRbdParams) middleware.Responder
}

// NewGetRbd creates a new http.Handler for the get rbd operation
func NewGetRbd(ctx *middleware.Context, handler GetRbdHandler) *GetRbd {
	return &GetRbd{Context: ctx, Handler: handler}
}

/*GetRbd swagger:route GET /attach/rbd/{id} attach getRbd

GetRbd get rbd API

*/
type GetRbd struct {
	Context *middleware.Context
	Handler GetRbdHandler
}

func (o *GetRbd) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetRbdParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
