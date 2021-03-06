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
	"github.com/go-openapi/swag"

	"github.com/dhis2-sre/im-job/swagger/sdk/models"
)

// NewJobLogsParams creates a new JobLogsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewJobLogsParams() *JobLogsParams {
	return &JobLogsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewJobLogsParamsWithTimeout creates a new JobLogsParams object
// with the ability to set a timeout on a request.
func NewJobLogsParamsWithTimeout(timeout time.Duration) *JobLogsParams {
	return &JobLogsParams{
		timeout: timeout,
	}
}

// NewJobLogsParamsWithContext creates a new JobLogsParams object
// with the ability to set a context for a request.
func NewJobLogsParamsWithContext(ctx context.Context) *JobLogsParams {
	return &JobLogsParams{
		Context: ctx,
	}
}

// NewJobLogsParamsWithHTTPClient creates a new JobLogsParams object
// with the ability to set a custom HTTPClient for a request.
func NewJobLogsParamsWithHTTPClient(client *http.Client) *JobLogsParams {
	return &JobLogsParams{
		HTTPClient: client,
	}
}

/* JobLogsParams contains all the parameters to send to the API endpoint
   for the job logs operation.

   Typically these are written to a http.Request.
*/
type JobLogsParams struct {

	/* Body.

	   Logs request body parameter
	*/
	Body *models.LogsRequest

	// RunID.
	//
	// Format: uint64
	ID uint64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the job logs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *JobLogsParams) WithDefaults() *JobLogsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the job logs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *JobLogsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the job logs params
func (o *JobLogsParams) WithTimeout(timeout time.Duration) *JobLogsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the job logs params
func (o *JobLogsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the job logs params
func (o *JobLogsParams) WithContext(ctx context.Context) *JobLogsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the job logs params
func (o *JobLogsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the job logs params
func (o *JobLogsParams) WithHTTPClient(client *http.Client) *JobLogsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the job logs params
func (o *JobLogsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the job logs params
func (o *JobLogsParams) WithBody(body *models.LogsRequest) *JobLogsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the job logs params
func (o *JobLogsParams) SetBody(body *models.LogsRequest) {
	o.Body = body
}

// WithID adds the runID to the job logs params
func (o *JobLogsParams) WithID(runID uint64) *JobLogsParams {
	o.SetID(runID)
	return o
}

// SetID adds the runId to the job logs params
func (o *JobLogsParams) SetID(runID uint64) {
	o.ID = runID
}

// WriteToRequest writes these params to a swagger request
func (o *JobLogsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param runId
	if err := r.SetPathParam("runId", swag.FormatUint64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
