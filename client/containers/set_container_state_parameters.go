// Code generated by go-swagger; DO NOT EDIT.

package containers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewSetContainerStateParams creates a new SetContainerStateParams object
// with the default values initialized.
func NewSetContainerStateParams() *SetContainerStateParams {
	var ()
	return &SetContainerStateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSetContainerStateParamsWithTimeout creates a new SetContainerStateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSetContainerStateParamsWithTimeout(timeout time.Duration) *SetContainerStateParams {
	var ()
	return &SetContainerStateParams{

		timeout: timeout,
	}
}

// NewSetContainerStateParamsWithContext creates a new SetContainerStateParams object
// with the default values initialized, and the ability to set a context for a request
func NewSetContainerStateParamsWithContext(ctx context.Context) *SetContainerStateParams {
	var ()
	return &SetContainerStateParams{

		Context: ctx,
	}
}

// NewSetContainerStateParamsWithHTTPClient creates a new SetContainerStateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSetContainerStateParamsWithHTTPClient(client *http.Client) *SetContainerStateParams {
	var ()
	return &SetContainerStateParams{
		HTTPClient: client,
	}
}

/*SetContainerStateParams contains all the parameters to send to the API endpoint
for the set container state operation typically these are written to a http.Request
*/
type SetContainerStateParams struct {

	/*ID*/
	ID int64
	/*State*/
	State string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the set container state params
func (o *SetContainerStateParams) WithTimeout(timeout time.Duration) *SetContainerStateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set container state params
func (o *SetContainerStateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set container state params
func (o *SetContainerStateParams) WithContext(ctx context.Context) *SetContainerStateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set container state params
func (o *SetContainerStateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set container state params
func (o *SetContainerStateParams) WithHTTPClient(client *http.Client) *SetContainerStateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set container state params
func (o *SetContainerStateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the set container state params
func (o *SetContainerStateParams) WithID(id int64) *SetContainerStateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the set container state params
func (o *SetContainerStateParams) SetID(id int64) {
	o.ID = id
}

// WithState adds the state to the set container state params
func (o *SetContainerStateParams) WithState(state string) *SetContainerStateParams {
	o.SetState(state)
	return o
}

// SetState adds the state to the set container state params
func (o *SetContainerStateParams) SetState(state string) {
	o.State = state
}

// WriteToRequest writes these params to a swagger request
func (o *SetContainerStateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	// path param state
	if err := r.SetPathParam("state", o.State); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}