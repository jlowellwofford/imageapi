// Code generated by go-swagger; DO NOT EDIT.

package mounts

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

// NewGetMountOverlayParams creates a new GetMountOverlayParams object
// with the default values initialized.
func NewGetMountOverlayParams() *GetMountOverlayParams {
	var ()
	return &GetMountOverlayParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetMountOverlayParamsWithTimeout creates a new GetMountOverlayParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetMountOverlayParamsWithTimeout(timeout time.Duration) *GetMountOverlayParams {
	var ()
	return &GetMountOverlayParams{

		timeout: timeout,
	}
}

// NewGetMountOverlayParamsWithContext creates a new GetMountOverlayParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetMountOverlayParamsWithContext(ctx context.Context) *GetMountOverlayParams {
	var ()
	return &GetMountOverlayParams{

		Context: ctx,
	}
}

// NewGetMountOverlayParamsWithHTTPClient creates a new GetMountOverlayParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetMountOverlayParamsWithHTTPClient(client *http.Client) *GetMountOverlayParams {
	var ()
	return &GetMountOverlayParams{
		HTTPClient: client,
	}
}

/*GetMountOverlayParams contains all the parameters to send to the API endpoint
for the get mount overlay operation typically these are written to a http.Request
*/
type GetMountOverlayParams struct {

	/*ID*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get mount overlay params
func (o *GetMountOverlayParams) WithTimeout(timeout time.Duration) *GetMountOverlayParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get mount overlay params
func (o *GetMountOverlayParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get mount overlay params
func (o *GetMountOverlayParams) WithContext(ctx context.Context) *GetMountOverlayParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get mount overlay params
func (o *GetMountOverlayParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get mount overlay params
func (o *GetMountOverlayParams) WithHTTPClient(client *http.Client) *GetMountOverlayParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get mount overlay params
func (o *GetMountOverlayParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get mount overlay params
func (o *GetMountOverlayParams) WithID(id int64) *GetMountOverlayParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get mount overlay params
func (o *GetMountOverlayParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetMountOverlayParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
