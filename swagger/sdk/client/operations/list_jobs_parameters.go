// Code generated by go-swagger; DO NOT EDIT.

package operations

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

// NewListJobsParams creates a new ListJobsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListJobsParams() *ListJobsParams {
	return &ListJobsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListJobsParamsWithTimeout creates a new ListJobsParams object
// with the ability to set a timeout on a request.
func NewListJobsParamsWithTimeout(timeout time.Duration) *ListJobsParams {
	return &ListJobsParams{
		timeout: timeout,
	}
}

// NewListJobsParamsWithContext creates a new ListJobsParams object
// with the ability to set a context for a request.
func NewListJobsParamsWithContext(ctx context.Context) *ListJobsParams {
	return &ListJobsParams{
		Context: ctx,
	}
}

// NewListJobsParamsWithHTTPClient creates a new ListJobsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListJobsParamsWithHTTPClient(client *http.Client) *ListJobsParams {
	return &ListJobsParams{
		HTTPClient: client,
	}
}

/* ListJobsParams contains all the parameters to send to the API endpoint
   for the list jobs operation.

   Typically these are written to a http.Request.
*/
type ListJobsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list jobs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListJobsParams) WithDefaults() *ListJobsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list jobs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListJobsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list jobs params
func (o *ListJobsParams) WithTimeout(timeout time.Duration) *ListJobsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list jobs params
func (o *ListJobsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list jobs params
func (o *ListJobsParams) WithContext(ctx context.Context) *ListJobsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list jobs params
func (o *ListJobsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list jobs params
func (o *ListJobsParams) WithHTTPClient(client *http.Client) *ListJobsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list jobs params
func (o *ListJobsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListJobsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
