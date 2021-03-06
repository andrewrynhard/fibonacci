// Code generated by go-swagger; DO NOT EDIT.

package sequence

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetSequenceParams creates a new GetSequenceParams object
// with the default values initialized.
func NewGetSequenceParams() *GetSequenceParams {
	var ()
	return &GetSequenceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSequenceParamsWithTimeout creates a new GetSequenceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSequenceParamsWithTimeout(timeout time.Duration) *GetSequenceParams {
	var ()
	return &GetSequenceParams{

		timeout: timeout,
	}
}

// NewGetSequenceParamsWithContext creates a new GetSequenceParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSequenceParamsWithContext(ctx context.Context) *GetSequenceParams {
	var ()
	return &GetSequenceParams{

		Context: ctx,
	}
}

// NewGetSequenceParamsWithHTTPClient creates a new GetSequenceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSequenceParamsWithHTTPClient(client *http.Client) *GetSequenceParams {
	var ()
	return &GetSequenceParams{
		HTTPClient: client,
	}
}

/*GetSequenceParams contains all the parameters to send to the API endpoint
for the get sequence operation typically these are written to a http.Request
*/
type GetSequenceParams struct {

	/*N
	  Specifies which number the sequence should stop at (inclusively)

	*/
	N int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get sequence params
func (o *GetSequenceParams) WithTimeout(timeout time.Duration) *GetSequenceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get sequence params
func (o *GetSequenceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get sequence params
func (o *GetSequenceParams) WithContext(ctx context.Context) *GetSequenceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get sequence params
func (o *GetSequenceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get sequence params
func (o *GetSequenceParams) WithHTTPClient(client *http.Client) *GetSequenceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get sequence params
func (o *GetSequenceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithN adds the n to the get sequence params
func (o *GetSequenceParams) WithN(n int64) *GetSequenceParams {
	o.SetN(n)
	return o
}

// SetN adds the n to the get sequence params
func (o *GetSequenceParams) SetN(n int64) {
	o.N = n
}

// WriteToRequest writes these params to a swagger request
func (o *GetSequenceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param n
	if err := r.SetPathParam("n", swag.FormatInt64(o.N)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
