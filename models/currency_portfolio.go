// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CurrencyPortfolio currency portfolio
//
// swagger:model currency_portfolio
type CurrencyPortfolio struct {

	// available funds
	// Required: true
	AvailableFunds *float64 `json:"available_funds"`

	// available withdrawal funds
	// Required: true
	AvailableWithdrawalFunds *float64 `json:"available_withdrawal_funds"`

	// balance
	// Required: true
	Balance *float64 `json:"balance"`

	// currency
	// Required: true
	// Enum: [btc eth]
	Currency *string `json:"currency"`

	// equity
	// Required: true
	Equity *float64 `json:"equity"`

	// initial margin
	// Required: true
	InitialMargin *float64 `json:"initial_margin"`

	// maintenance margin
	// Required: true
	MaintenanceMargin *float64 `json:"maintenance_margin"`

	// margin balance
	// Required: true
	MarginBalance *float64 `json:"margin_balance"`
}

// Validate validates this currency portfolio
func (m *CurrencyPortfolio) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAvailableFunds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAvailableWithdrawalFunds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBalance(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCurrency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEquity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInitialMargin(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaintenanceMargin(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMarginBalance(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CurrencyPortfolio) validateAvailableFunds(formats strfmt.Registry) error {

	if err := validate.Required("available_funds", "body", m.AvailableFunds); err != nil {
		return err
	}

	return nil
}

func (m *CurrencyPortfolio) validateAvailableWithdrawalFunds(formats strfmt.Registry) error {

	if err := validate.Required("available_withdrawal_funds", "body", m.AvailableWithdrawalFunds); err != nil {
		return err
	}

	return nil
}

func (m *CurrencyPortfolio) validateBalance(formats strfmt.Registry) error {

	if err := validate.Required("balance", "body", m.Balance); err != nil {
		return err
	}

	return nil
}

var currencyPortfolioTypeCurrencyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["btc","eth"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		currencyPortfolioTypeCurrencyPropEnum = append(currencyPortfolioTypeCurrencyPropEnum, v)
	}
}

const (

	// CurrencyPortfolioCurrencyBtc captures enum value "btc"
	CurrencyPortfolioCurrencyBtc string = "btc"

	// CurrencyPortfolioCurrencyEth captures enum value "eth"
	CurrencyPortfolioCurrencyEth string = "eth"
)

// prop value enum
func (m *CurrencyPortfolio) validateCurrencyEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, currencyPortfolioTypeCurrencyPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CurrencyPortfolio) validateCurrency(formats strfmt.Registry) error {

	if err := validate.Required("currency", "body", m.Currency); err != nil {
		return err
	}

	// value enum
	if err := m.validateCurrencyEnum("currency", "body", *m.Currency); err != nil {
		return err
	}

	return nil
}

func (m *CurrencyPortfolio) validateEquity(formats strfmt.Registry) error {

	if err := validate.Required("equity", "body", m.Equity); err != nil {
		return err
	}

	return nil
}

func (m *CurrencyPortfolio) validateInitialMargin(formats strfmt.Registry) error {

	if err := validate.Required("initial_margin", "body", m.InitialMargin); err != nil {
		return err
	}

	return nil
}

func (m *CurrencyPortfolio) validateMaintenanceMargin(formats strfmt.Registry) error {

	if err := validate.Required("maintenance_margin", "body", m.MaintenanceMargin); err != nil {
		return err
	}

	return nil
}

func (m *CurrencyPortfolio) validateMarginBalance(formats strfmt.Registry) error {

	if err := validate.Required("margin_balance", "body", m.MarginBalance); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CurrencyPortfolio) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CurrencyPortfolio) UnmarshalBinary(b []byte) error {
	var res CurrencyPortfolio
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
