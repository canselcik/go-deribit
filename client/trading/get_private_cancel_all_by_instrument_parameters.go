// Code generated by go-swagger; DO NOT EDIT.

package trading

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

// NewGetPrivateCancelAllByInstrumentParams creates a new GetPrivateCancelAllByInstrumentParams object
// with the default values initialized.
func NewGetPrivateCancelAllByInstrumentParams() *GetPrivateCancelAllByInstrumentParams {
	var ()
	return &GetPrivateCancelAllByInstrumentParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPrivateCancelAllByInstrumentParamsWithTimeout creates a new GetPrivateCancelAllByInstrumentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPrivateCancelAllByInstrumentParamsWithTimeout(timeout time.Duration) *GetPrivateCancelAllByInstrumentParams {
	var ()
	return &GetPrivateCancelAllByInstrumentParams{

		timeout: timeout,
	}
}

// NewGetPrivateCancelAllByInstrumentParamsWithContext creates a new GetPrivateCancelAllByInstrumentParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPrivateCancelAllByInstrumentParamsWithContext(ctx context.Context) *GetPrivateCancelAllByInstrumentParams {
	var ()
	return &GetPrivateCancelAllByInstrumentParams{

		Context: ctx,
	}
}

// NewGetPrivateCancelAllByInstrumentParamsWithHTTPClient creates a new GetPrivateCancelAllByInstrumentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPrivateCancelAllByInstrumentParamsWithHTTPClient(client *http.Client) *GetPrivateCancelAllByInstrumentParams {
	var ()
	return &GetPrivateCancelAllByInstrumentParams{
		HTTPClient: client,
	}
}

/*GetPrivateCancelAllByInstrumentParams contains all the parameters to send to the API endpoint
for the get private cancel all by instrument operation typically these are written to a http.Request
*/
type GetPrivateCancelAllByInstrumentParams struct {

	/*InstrumentName
	  Instrument name

	*/
	InstrumentName string
	/*Type
	  Order type - limit, stop or all, default - `all`

	*/
	Type *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get private cancel all by instrument params
func (o *GetPrivateCancelAllByInstrumentParams) WithTimeout(timeout time.Duration) *GetPrivateCancelAllByInstrumentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get private cancel all by instrument params
func (o *GetPrivateCancelAllByInstrumentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get private cancel all by instrument params
func (o *GetPrivateCancelAllByInstrumentParams) WithContext(ctx context.Context) *GetPrivateCancelAllByInstrumentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get private cancel all by instrument params
func (o *GetPrivateCancelAllByInstrumentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get private cancel all by instrument params
func (o *GetPrivateCancelAllByInstrumentParams) WithHTTPClient(client *http.Client) *GetPrivateCancelAllByInstrumentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get private cancel all by instrument params
func (o *GetPrivateCancelAllByInstrumentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInstrumentName adds the instrumentName to the get private cancel all by instrument params
func (o *GetPrivateCancelAllByInstrumentParams) WithInstrumentName(instrumentName string) *GetPrivateCancelAllByInstrumentParams {
	o.SetInstrumentName(instrumentName)
	return o
}

// SetInstrumentName adds the instrumentName to the get private cancel all by instrument params
func (o *GetPrivateCancelAllByInstrumentParams) SetInstrumentName(instrumentName string) {
	o.InstrumentName = instrumentName
}

// WithType adds the typeVar to the get private cancel all by instrument params
func (o *GetPrivateCancelAllByInstrumentParams) WithType(typeVar *string) *GetPrivateCancelAllByInstrumentParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the get private cancel all by instrument params
func (o *GetPrivateCancelAllByInstrumentParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *GetPrivateCancelAllByInstrumentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param instrument_name
	qrInstrumentName := o.InstrumentName
	qInstrumentName := qrInstrumentName
	if qInstrumentName != "" {
		if err := r.SetQueryParam("instrument_name", qInstrumentName); err != nil {
			return err
		}
	}

	if o.Type != nil {

		// query param type
		var qrType string
		if o.Type != nil {
			qrType = *o.Type
		}
		qType := qrType
		if qType != "" {
			if err := r.SetQueryParam("type", qType); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
