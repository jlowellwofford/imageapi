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

// NewGetContainerParams creates a new GetContainerParams object
// with the default values initialized.
func NewGetContainerParams() *GetContainerParams {
	var ()
	return &GetContainerParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetContainerParamsWithTimeout creates a new GetContainerParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetContainerParamsWithTimeout(timeout time.Duration) *GetContainerParams {
	var ()
	return &GetContainerParams{

		timeout: timeout,
	}
}

// NewGetContainerParamsWithContext creates a new GetContainerParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetContainerParamsWithContext(ctx context.Context) *GetContainerParams {
	var ()
	return &GetContainerParams{

		Context: ctx,
	}
}

// NewGetContainerParamsWithHTTPClient creates a new GetContainerParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetContainerParamsWithHTTPClient(client *http.Client) *GetContainerParams {
	var ()
	return &GetContainerParams{
		HTTPClient: client,
	}
}

/*GetContainerParams contains all the parameters to send to the API endpoint
for the get container operation typically these are written to a http.Request
*/
type GetContainerParams struct {

	/*ID*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get container params
func (o *GetContainerParams) WithTimeout(timeout time.Duration) *GetContainerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get container params
func (o *GetContainerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get container params
func (o *GetContainerParams) WithContext(ctx context.Context) *GetContainerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get container params
func (o *GetContainerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get container params
func (o *GetContainerParams) WithHTTPClient(client *http.Client) *GetContainerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get container params
func (o *GetContainerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get container params
func (o *GetContainerParams) WithID(id int64) *GetContainerParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get container params
func (o *GetContainerParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetContainerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
