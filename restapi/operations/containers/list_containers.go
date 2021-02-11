// Code generated by go-swagger; DO NOT EDIT.

package containers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ListContainersHandlerFunc turns a function with the right signature into a list containers handler
type ListContainersHandlerFunc func(ListContainersParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListContainersHandlerFunc) Handle(params ListContainersParams) middleware.Responder {
	return fn(params)
}

// ListContainersHandler interface for that can handle valid list containers params
type ListContainersHandler interface {
	Handle(ListContainersParams) middleware.Responder
}

// NewListContainers creates a new http.Handler for the list containers operation
func NewListContainers(ctx *middleware.Context, handler ListContainersHandler) *ListContainers {
	return &ListContainers{Context: ctx, Handler: handler}
}

/*ListContainers swagger:route GET /container containers listContainers

Get a list of containers

*/
type ListContainers struct {
	Context *middleware.Context
	Handler ListContainersHandler
}

func (o *ListContainers) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListContainersParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}