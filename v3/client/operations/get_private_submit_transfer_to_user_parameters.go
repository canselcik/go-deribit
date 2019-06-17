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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetPrivateSubmitTransferToUserParams creates a new GetPrivateSubmitTransferToUserParams object
// with the default values initialized.
func NewGetPrivateSubmitTransferToUserParams() *GetPrivateSubmitTransferToUserParams {
	var ()
	return &GetPrivateSubmitTransferToUserParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPrivateSubmitTransferToUserParamsWithTimeout creates a new GetPrivateSubmitTransferToUserParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPrivateSubmitTransferToUserParamsWithTimeout(timeout time.Duration) *GetPrivateSubmitTransferToUserParams {
	var ()
	return &GetPrivateSubmitTransferToUserParams{

		timeout: timeout,
	}
}

// NewGetPrivateSubmitTransferToUserParamsWithContext creates a new GetPrivateSubmitTransferToUserParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPrivateSubmitTransferToUserParamsWithContext(ctx context.Context) *GetPrivateSubmitTransferToUserParams {
	var ()
	return &GetPrivateSubmitTransferToUserParams{

		Context: ctx,
	}
}

// NewGetPrivateSubmitTransferToUserParamsWithHTTPClient creates a new GetPrivateSubmitTransferToUserParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPrivateSubmitTransferToUserParamsWithHTTPClient(client *http.Client) *GetPrivateSubmitTransferToUserParams {
	var ()
	return &GetPrivateSubmitTransferToUserParams{
		HTTPClient: client,
	}
}

/*GetPrivateSubmitTransferToUserParams contains all the parameters to send to the API endpoint
for the get private submit transfer to user operation typically these are written to a http.Request
*/
type GetPrivateSubmitTransferToUserParams struct {

	/*Amount
	  Amount of funds to be transferred

	*/
	Amount float64
	/*Currency
	  The currency symbol

	*/
	Currency string
	/*Destination
	  Destination address from address book

	*/
	Destination string
	/*Tfa
	  TFA code, required when TFA is enabled for current account

	*/
	Tfa *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get private submit transfer to user params
func (o *GetPrivateSubmitTransferToUserParams) WithTimeout(timeout time.Duration) *GetPrivateSubmitTransferToUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get private submit transfer to user params
func (o *GetPrivateSubmitTransferToUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get private submit transfer to user params
func (o *GetPrivateSubmitTransferToUserParams) WithContext(ctx context.Context) *GetPrivateSubmitTransferToUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get private submit transfer to user params
func (o *GetPrivateSubmitTransferToUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get private submit transfer to user params
func (o *GetPrivateSubmitTransferToUserParams) WithHTTPClient(client *http.Client) *GetPrivateSubmitTransferToUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get private submit transfer to user params
func (o *GetPrivateSubmitTransferToUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAmount adds the amount to the get private submit transfer to user params
func (o *GetPrivateSubmitTransferToUserParams) WithAmount(amount float64) *GetPrivateSubmitTransferToUserParams {
	o.SetAmount(amount)
	return o
}

// SetAmount adds the amount to the get private submit transfer to user params
func (o *GetPrivateSubmitTransferToUserParams) SetAmount(amount float64) {
	o.Amount = amount
}

// WithCurrency adds the currency to the get private submit transfer to user params
func (o *GetPrivateSubmitTransferToUserParams) WithCurrency(currency string) *GetPrivateSubmitTransferToUserParams {
	o.SetCurrency(currency)
	return o
}

// SetCurrency adds the currency to the get private submit transfer to user params
func (o *GetPrivateSubmitTransferToUserParams) SetCurrency(currency string) {
	o.Currency = currency
}

// WithDestination adds the destination to the get private submit transfer to user params
func (o *GetPrivateSubmitTransferToUserParams) WithDestination(destination string) *GetPrivateSubmitTransferToUserParams {
	o.SetDestination(destination)
	return o
}

// SetDestination adds the destination to the get private submit transfer to user params
func (o *GetPrivateSubmitTransferToUserParams) SetDestination(destination string) {
	o.Destination = destination
}

// WithTfa adds the tfa to the get private submit transfer to user params
func (o *GetPrivateSubmitTransferToUserParams) WithTfa(tfa *string) *GetPrivateSubmitTransferToUserParams {
	o.SetTfa(tfa)
	return o
}

// SetTfa adds the tfa to the get private submit transfer to user params
func (o *GetPrivateSubmitTransferToUserParams) SetTfa(tfa *string) {
	o.Tfa = tfa
}

// WriteToRequest writes these params to a swagger request
func (o *GetPrivateSubmitTransferToUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param amount
	qrAmount := o.Amount
	qAmount := swag.FormatFloat64(qrAmount)
	if qAmount != "" {
		if err := r.SetQueryParam("amount", qAmount); err != nil {
			return err
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

	// query param destination
	qrDestination := o.Destination
	qDestination := qrDestination
	if qDestination != "" {
		if err := r.SetQueryParam("destination", qDestination); err != nil {
			return err
		}
	}

	if o.Tfa != nil {

		// query param tfa
		var qrTfa string
		if o.Tfa != nil {
			qrTfa = *o.Tfa
		}
		qTfa := qrTfa
		if qTfa != "" {
			if err := r.SetQueryParam("tfa", qTfa); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}