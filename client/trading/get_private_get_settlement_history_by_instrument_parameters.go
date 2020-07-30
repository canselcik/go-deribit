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
	"github.com/go-openapi/swag"
)

// NewGetPrivateGetSettlementHistoryByInstrumentParams creates a new GetPrivateGetSettlementHistoryByInstrumentParams object
// with the default values initialized.
func NewGetPrivateGetSettlementHistoryByInstrumentParams() *GetPrivateGetSettlementHistoryByInstrumentParams {
	var ()
	return &GetPrivateGetSettlementHistoryByInstrumentParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPrivateGetSettlementHistoryByInstrumentParamsWithTimeout creates a new GetPrivateGetSettlementHistoryByInstrumentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPrivateGetSettlementHistoryByInstrumentParamsWithTimeout(timeout time.Duration) *GetPrivateGetSettlementHistoryByInstrumentParams {
	var ()
	return &GetPrivateGetSettlementHistoryByInstrumentParams{

		timeout: timeout,
	}
}

// NewGetPrivateGetSettlementHistoryByInstrumentParamsWithContext creates a new GetPrivateGetSettlementHistoryByInstrumentParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPrivateGetSettlementHistoryByInstrumentParamsWithContext(ctx context.Context) *GetPrivateGetSettlementHistoryByInstrumentParams {
	var ()
	return &GetPrivateGetSettlementHistoryByInstrumentParams{

		Context: ctx,
	}
}

// NewGetPrivateGetSettlementHistoryByInstrumentParamsWithHTTPClient creates a new GetPrivateGetSettlementHistoryByInstrumentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPrivateGetSettlementHistoryByInstrumentParamsWithHTTPClient(client *http.Client) *GetPrivateGetSettlementHistoryByInstrumentParams {
	var ()
	return &GetPrivateGetSettlementHistoryByInstrumentParams{
		HTTPClient: client,
	}
}

/*GetPrivateGetSettlementHistoryByInstrumentParams contains all the parameters to send to the API endpoint
for the get private get settlement history by instrument operation typically these are written to a http.Request
*/
type GetPrivateGetSettlementHistoryByInstrumentParams struct {

	/*Count
	  Number of requested items, default - `20`

	*/
	Count *int64
	/*InstrumentName
	  Instrument name

	*/
	InstrumentName string
	/*Type
	  Settlement type

	*/
	Type *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get private get settlement history by instrument params
func (o *GetPrivateGetSettlementHistoryByInstrumentParams) WithTimeout(timeout time.Duration) *GetPrivateGetSettlementHistoryByInstrumentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get private get settlement history by instrument params
func (o *GetPrivateGetSettlementHistoryByInstrumentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get private get settlement history by instrument params
func (o *GetPrivateGetSettlementHistoryByInstrumentParams) WithContext(ctx context.Context) *GetPrivateGetSettlementHistoryByInstrumentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get private get settlement history by instrument params
func (o *GetPrivateGetSettlementHistoryByInstrumentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get private get settlement history by instrument params
func (o *GetPrivateGetSettlementHistoryByInstrumentParams) WithHTTPClient(client *http.Client) *GetPrivateGetSettlementHistoryByInstrumentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get private get settlement history by instrument params
func (o *GetPrivateGetSettlementHistoryByInstrumentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCount adds the count to the get private get settlement history by instrument params
func (o *GetPrivateGetSettlementHistoryByInstrumentParams) WithCount(count *int64) *GetPrivateGetSettlementHistoryByInstrumentParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the get private get settlement history by instrument params
func (o *GetPrivateGetSettlementHistoryByInstrumentParams) SetCount(count *int64) {
	o.Count = count
}

// WithInstrumentName adds the instrumentName to the get private get settlement history by instrument params
func (o *GetPrivateGetSettlementHistoryByInstrumentParams) WithInstrumentName(instrumentName string) *GetPrivateGetSettlementHistoryByInstrumentParams {
	o.SetInstrumentName(instrumentName)
	return o
}

// SetInstrumentName adds the instrumentName to the get private get settlement history by instrument params
func (o *GetPrivateGetSettlementHistoryByInstrumentParams) SetInstrumentName(instrumentName string) {
	o.InstrumentName = instrumentName
}

// WithType adds the typeVar to the get private get settlement history by instrument params
func (o *GetPrivateGetSettlementHistoryByInstrumentParams) WithType(typeVar *string) *GetPrivateGetSettlementHistoryByInstrumentParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the get private get settlement history by instrument params
func (o *GetPrivateGetSettlementHistoryByInstrumentParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *GetPrivateGetSettlementHistoryByInstrumentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Count != nil {

		// query param count
		var qrCount int64
		if o.Count != nil {
			qrCount = *o.Count
		}
		qCount := swag.FormatInt64(qrCount)
		if qCount != "" {
			if err := r.SetQueryParam("count", qCount); err != nil {
				return err
			}
		}

	}

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
