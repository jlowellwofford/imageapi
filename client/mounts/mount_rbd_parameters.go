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

	"github.com/jlowellwofford/imageapi/models"
)

// NewMountRbdParams creates a new MountRbdParams object
// with the default values initialized.
func NewMountRbdParams() *MountRbdParams {
	var ()
	return &MountRbdParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewMountRbdParamsWithTimeout creates a new MountRbdParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewMountRbdParamsWithTimeout(timeout time.Duration) *MountRbdParams {
	var ()
	return &MountRbdParams{

		timeout: timeout,
	}
}

// NewMountRbdParamsWithContext creates a new MountRbdParams object
// with the default values initialized, and the ability to set a context for a request
func NewMountRbdParamsWithContext(ctx context.Context) *MountRbdParams {
	var ()
	return &MountRbdParams{

		Context: ctx,
	}
}

// NewMountRbdParamsWithHTTPClient creates a new MountRbdParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewMountRbdParamsWithHTTPClient(client *http.Client) *MountRbdParams {
	var ()
	return &MountRbdParams{
		HTTPClient: client,
	}
}

/*MountRbdParams contains all the parameters to send to the API endpoint
for the mount rbd operation typically these are written to a http.Request
*/
type MountRbdParams struct {

	/*Mount*/
	Mount *models.MountRbd

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the mount rbd params
func (o *MountRbdParams) WithTimeout(timeout time.Duration) *MountRbdParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the mount rbd params
func (o *MountRbdParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the mount rbd params
func (o *MountRbdParams) WithContext(ctx context.Context) *MountRbdParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the mount rbd params
func (o *MountRbdParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the mount rbd params
func (o *MountRbdParams) WithHTTPClient(client *http.Client) *MountRbdParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the mount rbd params
func (o *MountRbdParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMount adds the mount to the mount rbd params
func (o *MountRbdParams) WithMount(mount *models.MountRbd) *MountRbdParams {
	o.SetMount(mount)
	return o
}

// SetMount adds the mount to the mount rbd params
func (o *MountRbdParams) SetMount(mount *models.MountRbd) {
	o.Mount = mount
}

// WriteToRequest writes these params to a swagger request
func (o *MountRbdParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Mount != nil {
		if err := r.SetBodyParam(o.Mount); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}