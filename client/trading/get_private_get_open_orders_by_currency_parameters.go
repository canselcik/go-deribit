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

// NewGetPrivateGetOpenOrdersByCurrencyParams creates a new GetPrivateGetOpenOrdersByCurrencyParams object
// with the default values initialized.
func NewGetPrivateGetOpenOrdersByCurrencyParams() *GetPrivateGetOpenOrdersByCurrencyParams {
	var ()
	return &GetPrivateGetOpenOrdersByCurrencyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPrivateGetOpenOrdersByCurrencyParamsWithTimeout creates a new GetPrivateGetOpenOrdersByCurrencyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPrivateGetOpenOrdersByCurrencyParamsWithTimeout(timeout time.Duration) *GetPrivateGetOpenOrdersByCurrencyParams {
	var ()
	return &GetPrivateGetOpenOrdersByCurrencyParams{

		timeout: timeout,
	}
}

// NewGetPrivateGetOpenOrdersByCurrencyParamsWithContext creates a new GetPrivateGetOpenOrdersByCurrencyParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPrivateGetOpenOrdersByCurrencyParamsWithContext(ctx context.Context) *GetPrivateGetOpenOrdersByCurrencyParams {
	var ()
	return &GetPrivateGetOpenOrdersByCurrencyParams{

		Context: ctx,
	}
}

// NewGetPrivateGetOpenOrdersByCurrencyParamsWithHTTPClient creates a new GetPrivateGetOpenOrdersByCurrencyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPrivateGetOpenOrdersByCurrencyParamsWithHTTPClient(client *http.Client) *GetPrivateGetOpenOrdersByCurrencyParams {
	var ()
	return &GetPrivateGetOpenOrdersByCurrencyParams{
		HTTPClient: client,
	}
}

/*GetPrivateGetOpenOrdersByCurrencyParams contains all the parameters to send to the API endpoint
for the get private get open orders by currency operation typically these are written to a http.Request
*/
type GetPrivateGetOpenOrdersByCurrencyParams struct {

	/*Currency
	  The currency symbol

	*/
	Currency string
	/*Kind
	  Instrument kind, if not provided instruments of all kinds are considered

	*/
	Kind *string
	/*Type
	  Order type, default - `all`

	*/
	Type *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get private get open orders by currency params
func (o *GetPrivateGetOpenOrdersByCurrencyParams) WithTimeout(timeout time.Duration) *GetPrivateGetOpenOrdersByCurrencyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get private get open orders by currency params
func (o *GetPrivateGetOpenOrdersByCurrencyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get private get open orders by currency params
func (o *GetPrivateGetOpenOrdersByCurrencyParams) WithContext(ctx context.Context) *GetPrivateGetOpenOrdersByCurrencyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get private get open orders by currency params
func (o *GetPrivateGetOpenOrdersByCurrencyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get private get open orders by currency params
func (o *GetPrivateGetOpenOrdersByCurrencyParams) WithHTTPClient(client *http.Client) *GetPrivateGetOpenOrdersByCurrencyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get private get open orders by currency params
func (o *GetPrivateGetOpenOrdersByCurrencyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCurrency adds the currency to the get private get open orders by currency params
func (o *GetPrivateGetOpenOrdersByCurrencyParams) WithCurrency(currency string) *GetPrivateGetOpenOrdersByCurrencyParams {
	o.SetCurrency(currency)
	return o
}

// SetCurrency adds the currency to the get private get open orders by currency params
func (o *GetPrivateGetOpenOrdersByCurrencyParams) SetCurrency(currency string) {
	o.Currency = currency
}

// WithKind adds the kind to the get private get open orders by currency params
func (o *GetPrivateGetOpenOrdersByCurrencyParams) WithKind(kind *string) *GetPrivateGetOpenOrdersByCurrencyParams {
	o.SetKind(kind)
	return o
}

// SetKind adds the kind to the get private get open orders by currency params
func (o *GetPrivateGetOpenOrdersByCurrencyParams) SetKind(kind *string) {
	o.Kind = kind
}

// WithType adds the typeVar to the get private get open orders by currency params
func (o *GetPrivateGetOpenOrdersByCurrencyParams) WithType(typeVar *string) *GetPrivateGetOpenOrdersByCurrencyParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the get private get open orders by currency params
func (o *GetPrivateGetOpenOrdersByCurrencyParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *GetPrivateGetOpenOrdersByCurrencyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param currency
	qrCurrency := o.Currency
	qCurrency := qrCurrency
	if qCurrency != "" {
		if err := r.SetQueryParam("currency", qCurrency); err != nil {
			return err
		}
	}

	if o.Kind != nil {

		// query param kind
		var qrKind string
		if o.Kind != nil {
			qrKind = *o.Kind
		}
		qKind := qrKind
		if qKind != "" {
			if err := r.SetQueryParam("kind", qKind); err != nil {
				return err
			}
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
