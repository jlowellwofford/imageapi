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
)

// NewListMountsParams creates a new ListMountsParams object
// with the default values initialized.
func NewListMountsParams() *ListMountsParams {

	return &ListMountsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListMountsParamsWithTimeout creates a new ListMountsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListMountsParamsWithTimeout(timeout time.Duration) *ListMountsParams {

	return &ListMountsParams{

		timeout: timeout,
	}
}

// NewListMountsParamsWithContext creates a new ListMountsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListMountsParamsWithContext(ctx context.Context) *ListMountsParams {

	return &ListMountsParams{

		Context: ctx,
	}
}

// NewListMountsParamsWithHTTPClient creates a new ListMountsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListMountsParamsWithHTTPClient(client *http.Client) *ListMountsParams {

	return &ListMountsParams{
		HTTPClient: client,
	}
}

/*ListMountsParams contains all the parameters to send to the API endpoint
for the list mounts operation typically these are written to a http.Request
*/
type ListMountsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list mounts params
func (o *ListMountsParams) WithTimeout(timeout time.Duration) *ListMountsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list mounts params
func (o *ListMountsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list mounts params
func (o *ListMountsParams) WithContext(ctx context.Context) *ListMountsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list mounts params
func (o *ListMountsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list mounts params
func (o *ListMountsParams) WithHTTPClient(client *http.Client) *ListMountsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list mounts params
func (o *ListMountsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListMountsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
