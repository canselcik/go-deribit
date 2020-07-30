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

// NewGetPublicAuthParams creates a new GetPublicAuthParams object
// with the default values initialized.
func NewGetPublicAuthParams() *GetPublicAuthParams {
	var ()
	return &GetPublicAuthParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPublicAuthParamsWithTimeout creates a new GetPublicAuthParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPublicAuthParamsWithTimeout(timeout time.Duration) *GetPublicAuthParams {
	var ()
	return &GetPublicAuthParams{

		timeout: timeout,
	}
}

// NewGetPublicAuthParamsWithContext creates a new GetPublicAuthParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPublicAuthParamsWithContext(ctx context.Context) *GetPublicAuthParams {
	var ()
	return &GetPublicAuthParams{

		Context: ctx,
	}
}

// NewGetPublicAuthParamsWithHTTPClient creates a new GetPublicAuthParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPublicAuthParamsWithHTTPClient(client *http.Client) *GetPublicAuthParams {
	var ()
	return &GetPublicAuthParams{
		HTTPClient: client,
	}
}

/*GetPublicAuthParams contains all the parameters to send to the API endpoint
for the get public auth operation typically these are written to a http.Request
*/
type GetPublicAuthParams struct {

	/*ClientID
	  Required for grant type `client_credentials` and `client_signature`

	*/
	ClientID string
	/*ClientSecret
	  Required for grant type `client_credentials`

	*/
	ClientSecret string
	/*GrantType
	  Method of authentication

	*/
	GrantType string
	/*Nonce
	  Optional for grant type `client_signature`; delivers user generated initialization vector for the server token

	*/
	Nonce *string
	/*Password
	  Required for grant type `password`

	*/
	Password string
	/*RefreshToken
	  Required for grant type `refresh_token`

	*/
	RefreshToken string
	/*Scope
	  Describes type of the access for assigned token, possible values: `connection`, `session:name`, `trade:[read, read_write, none]`, `wallet:[read, read_write, none]`, `account:[read, read_write, none]`, `expires:NUMBER`, `ip:ADDR`.</BR></BR> **NOTICE:** Depending on choosing an authentication method (```grant type```) some scopes could be narrowed by the server. e.g. when ```grant_type = client_credentials``` and ```scope = wallet:read_write``` it's modified by the server as ```scope = wallet:read```</BR></BR> Please check details described in [Access scope](#access-scope)

	*/
	Scope *string
	/*Signature
	  Required for grant type `client_signature`; it's a cryptographic signature calculated over provided fields using user **secret key**. The signature should be calculated as an HMAC (Hash-based Message Authentication Code) with `SHA256` hash algorithm

	*/
	Signature string
	/*State
	  Will be passed back in the response

	*/
	State *string
	/*Timestamp
	  Required for grant type `client_signature`, provides time when request has been generated

	*/
	Timestamp string
	/*Username
	  Required for grant type `password`

	*/
	Username string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get public auth params
func (o *GetPublicAuthParams) WithTimeout(timeout time.Duration) *GetPublicAuthParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get public auth params
func (o *GetPublicAuthParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get public auth params
func (o *GetPublicAuthParams) WithContext(ctx context.Context) *GetPublicAuthParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get public auth params
func (o *GetPublicAuthParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get public auth params
func (o *GetPublicAuthParams) WithHTTPClient(client *http.Client) *GetPublicAuthParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get public auth params
func (o *GetPublicAuthParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClientID adds the clientID to the get public auth params
func (o *GetPublicAuthParams) WithClientID(clientID string) *GetPublicAuthParams {
	o.SetClientID(clientID)
	return o
}

// SetClientID adds the clientId to the get public auth params
func (o *GetPublicAuthParams) SetClientID(clientID string) {
	o.ClientID = clientID
}

// WithClientSecret adds the clientSecret to the get public auth params
func (o *GetPublicAuthParams) WithClientSecret(clientSecret string) *GetPublicAuthParams {
	o.SetClientSecret(clientSecret)
	return o
}

// SetClientSecret adds the clientSecret to the get public auth params
func (o *GetPublicAuthParams) SetClientSecret(clientSecret string) {
	o.ClientSecret = clientSecret
}

// WithGrantType adds the grantType to the get public auth params
func (o *GetPublicAuthParams) WithGrantType(grantType string) *GetPublicAuthParams {
	o.SetGrantType(grantType)
	return o
}

// SetGrantType adds the grantType to the get public auth params
func (o *GetPublicAuthParams) SetGrantType(grantType string) {
	o.GrantType = grantType
}

// WithNonce adds the nonce to the get public auth params
func (o *GetPublicAuthParams) WithNonce(nonce *string) *GetPublicAuthParams {
	o.SetNonce(nonce)
	return o
}

// SetNonce adds the nonce to the get public auth params
func (o *GetPublicAuthParams) SetNonce(nonce *string) {
	o.Nonce = nonce
}

// WithPassword adds the password to the get public auth params
func (o *GetPublicAuthParams) WithPassword(password string) *GetPublicAuthParams {
	o.SetPassword(password)
	return o
}

// SetPassword adds the password to the get public auth params
func (o *GetPublicAuthParams) SetPassword(password string) {
	o.Password = password
}

// WithRefreshToken adds the refreshToken to the get public auth params
func (o *GetPublicAuthParams) WithRefreshToken(refreshToken string) *GetPublicAuthParams {
	o.SetRefreshToken(refreshToken)
	return o
}

// SetRefreshToken adds the refreshToken to the get public auth params
func (o *GetPublicAuthParams) SetRefreshToken(refreshToken string) {
	o.RefreshToken = refreshToken
}

// WithScope adds the scope to the get public auth params
func (o *GetPublicAuthParams) WithScope(scope *string) *GetPublicAuthParams {
	o.SetScope(scope)
	return o
}

// SetScope adds the scope to the get public auth params
func (o *GetPublicAuthParams) SetScope(scope *string) {
	o.Scope = scope
}

// WithSignature adds the signature to the get public auth params
func (o *GetPublicAuthParams) WithSignature(signature string) *GetPublicAuthParams {
	o.SetSignature(signature)
	return o
}

// SetSignature adds the signature to the get public auth params
func (o *GetPublicAuthParams) SetSignature(signature string) {
	o.Signature = signature
}

// WithState adds the state to the get public auth params
func (o *GetPublicAuthParams) WithState(state *string) *GetPublicAuthParams {
	o.SetState(state)
	return o
}

// SetState adds the state to the get public auth params
func (o *GetPublicAuthParams) SetState(state *string) {
	o.State = state
}

// WithTimestamp adds the timestamp to the get public auth params
func (o *GetPublicAuthParams) WithTimestamp(timestamp string) *GetPublicAuthParams {
	o.SetTimestamp(timestamp)
	return o
}

// SetTimestamp adds the timestamp to the get public auth params
func (o *GetPublicAuthParams) SetTimestamp(timestamp string) {
	o.Timestamp = timestamp
}

// WithUsername adds the username to the get public auth params
func (o *GetPublicAuthParams) WithUsername(username string) *GetPublicAuthParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the get public auth params
func (o *GetPublicAuthParams) SetUsername(username string) {
	o.Username = username
}

// WriteToRequest writes these params to a swagger request
func (o *GetPublicAuthParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param client_id
	qrClientID := o.ClientID
	qClientID := qrClientID
	if qClientID != "" {
		if err := r.SetQueryParam("client_id", qClientID); err != nil {
			return err
		}
	}

	// query param client_secret
	qrClientSecret := o.ClientSecret
	qClientSecret := qrClientSecret
	if qClientSecret != "" {
		if err := r.SetQueryParam("client_secret", qClientSecret); err != nil {
			return err
		}
	}

	// query param grant_type
	qrGrantType := o.GrantType
	qGrantType := qrGrantType
	if qGrantType != "" {
		if err := r.SetQueryParam("grant_type", qGrantType); err != nil {
			return err
		}
	}

	if o.Nonce != nil {

		// query param nonce
		var qrNonce string
		if o.Nonce != nil {
			qrNonce = *o.Nonce
		}
		qNonce := qrNonce
		if qNonce != "" {
			if err := r.SetQueryParam("nonce", qNonce); err != nil {
				return err
			}
		}

	}

	// query param password
	qrPassword := o.Password
	qPassword := qrPassword
	if qPassword != "" {
		if err := r.SetQueryParam("password", qPassword); err != nil {
			return err
		}
	}

	// query param refresh_token
	qrRefreshToken := o.RefreshToken
	qRefreshToken := qrRefreshToken
	if qRefreshToken != "" {
		if err := r.SetQueryParam("refresh_token", qRefreshToken); err != nil {
			return err
		}
	}

	if o.Scope != nil {

		// query param scope
		var qrScope string
		if o.Scope != nil {
			qrScope = *o.Scope
		}
		qScope := qrScope
		if qScope != "" {
			if err := r.SetQueryParam("scope", qScope); err != nil {
				return err
			}
		}

	}

	// query param signature
	qrSignature := o.Signature
	qSignature := qrSignature
	if qSignature != "" {
		if err := r.SetQueryParam("signature", qSignature); err != nil {
			return err
		}
	}

	if o.State != nil {

		// query param state
		var qrState string
		if o.State != nil {
			qrState = *o.State
		}
		qState := qrState
		if qState != "" {
			if err := r.SetQueryParam("state", qState); err != nil {
				return err
			}
		}

	}

	// query param timestamp
	qrTimestamp := o.Timestamp
	qTimestamp := qrTimestamp
	if qTimestamp != "" {
		if err := r.SetQueryParam("timestamp", qTimestamp); err != nil {
			return err
		}
	}

	// query param username
	qrUsername := o.Username
	qUsername := qrUsername
	if qUsername != "" {
		if err := r.SetQueryParam("username", qUsername); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
