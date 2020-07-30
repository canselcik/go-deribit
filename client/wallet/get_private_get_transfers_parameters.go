// Code generated by go-swagger; DO NOT EDIT.

package wallet

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

// NewGetPrivateGetTransfersParams creates a new GetPrivateGetTransfersParams object
// with the default values initialized.
func NewGetPrivateGetTransfersParams() *GetPrivateGetTransfersParams {
	var ()
	return &GetPrivateGetTransfersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPrivateGetTransfersParamsWithTimeout creates a new GetPrivateGetTransfersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPrivateGetTransfersParamsWithTimeout(timeout time.Duration) *GetPrivateGetTransfersParams {
	var ()
	return &GetPrivateGetTransfersParams{

		timeout: timeout,
	}
}

// NewGetPrivateGetTransfersParamsWithContext creates a new GetPrivateGetTransfersParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPrivateGetTransfersParamsWithContext(ctx context.Context) *GetPrivateGetTransfersParams {
	var ()
	return &GetPrivateGetTransfersParams{

		Context: ctx,
	}
}

// NewGetPrivateGetTransfersParamsWithHTTPClient creates a new GetPrivateGetTransfersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPrivateGetTransfersParamsWithHTTPClient(client *http.Client) *GetPrivateGetTransfersParams {
	var ()
	return &GetPrivateGetTransfersParams{
		HTTPClient: client,
	}
}

/*GetPrivateGetTransfersParams contains all the parameters to send to the API endpoint
for the get private get transfers operation typically these are written to a http.Request
*/
type GetPrivateGetTransfersParams struct {

	/*Count
	  Number of requested items, default - `10`

	*/
	Count *int64
	/*Currency
	  The currency symbol

	*/
	Currency string
	/*Offset
	  The offset for pagination, default - `0`

	*/
	Offset *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get private get transfers params
func (o *GetPrivateGetTransfersParams) WithTimeout(timeout time.Duration) *GetPrivateGetTransfersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get private get transfers params
func (o *GetPrivateGetTransfersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get private get transfers params
func (o *GetPrivateGetTransfersParams) WithContext(ctx context.Context) *GetPrivateGetTransfersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get private get transfers params
func (o *GetPrivateGetTransfersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get private get transfers params
func (o *GetPrivateGetTransfersParams) WithHTTPClient(client *http.Client) *GetPrivateGetTransfersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get private get transfers params
func (o *GetPrivateGetTransfersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCount adds the count to the get private get transfers params
func (o *GetPrivateGetTransfersParams) WithCount(count *int64) *GetPrivateGetTransfersParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the get private get transfers params
func (o *GetPrivateGetTransfersParams) SetCount(count *int64) {
	o.Count = count
}

// WithCurrency adds the currency to the get private get transfers params
func (o *GetPrivateGetTransfersParams) WithCurrency(currency string) *GetPrivateGetTransfersParams {
	o.SetCurrency(currency)
	return o
}

// SetCurrency adds the currency to the get private get transfers params
func (o *GetPrivateGetTransfersParams) SetCurrency(currency string) {
	o.Currency = currency
}

// WithOffset adds the offset to the get private get transfers params
func (o *GetPrivateGetTransfersParams) WithOffset(offset *int64) *GetPrivateGetTransfersParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get private get transfers params
func (o *GetPrivateGetTransfersParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WriteToRequest writes these params to a swagger request
func (o *GetPrivateGetTransfersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// query param currency
	qrCurrency := o.Currency
	qCurrency := qrCurrency
	if qCurrency != "" {
		if err := r.SetQueryParam("currency", qCurrency); err != nil {
			return err
		}
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
