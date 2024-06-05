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

// NewGetObjectsParams creates a new GetObjectsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetObjectsParams() *GetObjectsParams {
	return &GetObjectsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetObjectsParamsWithTimeout creates a new GetObjectsParams object
// with the ability to set a timeout on a request.
func NewGetObjectsParamsWithTimeout(timeout time.Duration) *GetObjectsParams {
	return &GetObjectsParams{
		timeout: timeout,
	}
}

// NewGetObjectsParamsWithContext creates a new GetObjectsParams object
// with the ability to set a context for a request.
func NewGetObjectsParamsWithContext(ctx context.Context) *GetObjectsParams {
	return &GetObjectsParams{
		Context: ctx,
	}
}

// NewGetObjectsParamsWithHTTPClient creates a new GetObjectsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetObjectsParamsWithHTTPClient(client *http.Client) *GetObjectsParams {
	return &GetObjectsParams{
		HTTPClient: client,
	}
}

/*
GetObjectsParams contains all the parameters to send to the API endpoint

	for the get objects operation.

	Typically these are written to a http.Request.
*/
type GetObjectsParams struct {

	/* Query.

	   query string
	*/
	Query string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get objects params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetObjectsParams) WithDefaults() *GetObjectsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get objects params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetObjectsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get objects params
func (o *GetObjectsParams) WithTimeout(timeout time.Duration) *GetObjectsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get objects params
func (o *GetObjectsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get objects params
func (o *GetObjectsParams) WithContext(ctx context.Context) *GetObjectsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get objects params
func (o *GetObjectsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get objects params
func (o *GetObjectsParams) WithHTTPClient(client *http.Client) *GetObjectsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get objects params
func (o *GetObjectsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithQuery adds the query to the get objects params
func (o *GetObjectsParams) WithQuery(query string) *GetObjectsParams {
	o.SetQuery(query)
	return o
}

// SetQuery adds the query to the get objects params
func (o *GetObjectsParams) SetQuery(query string) {
	o.Query = query
}

// WriteToRequest writes these params to a swagger request
func (o *GetObjectsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param query
	qrQuery := o.Query
	qQuery := qrQuery
	if qQuery != "" {

		if err := r.SetQueryParam("query", qQuery); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
