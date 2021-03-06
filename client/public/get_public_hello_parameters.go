// Code generated by go-swagger; DO NOT EDIT.

package public

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

// NewGetPublicHelloParams creates a new GetPublicHelloParams object
// with the default values initialized.
func NewGetPublicHelloParams() *GetPublicHelloParams {
	var ()
	return &GetPublicHelloParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPublicHelloParamsWithTimeout creates a new GetPublicHelloParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPublicHelloParamsWithTimeout(timeout time.Duration) *GetPublicHelloParams {
	var ()
	return &GetPublicHelloParams{

		timeout: timeout,
	}
}

// NewGetPublicHelloParamsWithContext creates a new GetPublicHelloParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPublicHelloParamsWithContext(ctx context.Context) *GetPublicHelloParams {
	var ()
	return &GetPublicHelloParams{

		Context: ctx,
	}
}

// NewGetPublicHelloParamsWithHTTPClient creates a new GetPublicHelloParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPublicHelloParamsWithHTTPClient(client *http.Client) *GetPublicHelloParams {
	var ()
	return &GetPublicHelloParams{
		HTTPClient: client,
	}
}

/*GetPublicHelloParams contains all the parameters to send to the API endpoint
for the get public hello operation typically these are written to a http.Request
*/
type GetPublicHelloParams struct {

	/*ClientName
	  Client software name

	*/
	ClientName string
	/*ClientVersion
	  Client software version

	*/
	ClientVersion string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get public hello params
func (o *GetPublicHelloParams) WithTimeout(timeout time.Duration) *GetPublicHelloParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get public hello params
func (o *GetPublicHelloParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get public hello params
func (o *GetPublicHelloParams) WithContext(ctx context.Context) *GetPublicHelloParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get public hello params
func (o *GetPublicHelloParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get public hello params
func (o *GetPublicHelloParams) WithHTTPClient(client *http.Client) *GetPublicHelloParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get public hello params
func (o *GetPublicHelloParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClientName adds the clientName to the get public hello params
func (o *GetPublicHelloParams) WithClientName(clientName string) *GetPublicHelloParams {
	o.SetClientName(clientName)
	return o
}

// SetClientName adds the clientName to the get public hello params
func (o *GetPublicHelloParams) SetClientName(clientName string) {
	o.ClientName = clientName
}

// WithClientVersion adds the clientVersion to the get public hello params
func (o *GetPublicHelloParams) WithClientVersion(clientVersion string) *GetPublicHelloParams {
	o.SetClientVersion(clientVersion)
	return o
}

// SetClientVersion adds the clientVersion to the get public hello params
func (o *GetPublicHelloParams) SetClientVersion(clientVersion string) {
	o.ClientVersion = clientVersion
}

// WriteToRequest writes these params to a swagger request
func (o *GetPublicHelloParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param client_name
	qrClientName := o.ClientName
	qClientName := qrClientName
	if qClientName != "" {
		if err := r.SetQueryParam("client_name", qClientName); err != nil {
			return err
		}
	}

	// query param client_version
	qrClientVersion := o.ClientVersion
	qClientVersion := qrClientVersion
	if qClientVersion != "" {
		if err := r.SetQueryParam("client_version", qClientVersion); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
