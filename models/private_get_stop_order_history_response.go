// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PrivateGetStopOrderHistoryResponse private get stop order history response
//
// swagger:model private_get_stop_order_history_response
type PrivateGetStopOrderHistoryResponse struct {
	BaseMessage

	// result
	// Required: true
	Result *PrivateGetStopOrderHistoryResponseAO1Result `json:"result"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *PrivateGetStopOrderHistoryResponse) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 BaseMessage
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.BaseMessage = aO0

	// AO1
	var dataAO1 struct {
		Result *PrivateGetStopOrderHistoryResponseAO1Result `json:"result"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Result = dataAO1.Result

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m PrivateGetStopOrderHistoryResponse) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.BaseMessage)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		Result *PrivateGetStopOrderHistoryResponseAO1Result `json:"result"`
	}

	dataAO1.Result = m.Result

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this private get stop order history response
func (m *PrivateGetStopOrderHistoryResponse) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with BaseMessage
	if err := m.BaseMessage.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PrivateGetStopOrderHistoryResponse) validateResult(formats strfmt.Registry) error {

	if err := validate.Required("result", "body", m.Result); err != nil {
		return err
	}

	if m.Result != nil {
		if err := m.Result.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("result")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PrivateGetStopOrderHistoryResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PrivateGetStopOrderHistoryResponse) UnmarshalBinary(b []byte) error {
	var res PrivateGetStopOrderHistoryResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PrivateGetStopOrderHistoryResponseAO1Result private get stop order history response a o1 result
//
// swagger:model PrivateGetStopOrderHistoryResponseAO1Result
type PrivateGetStopOrderHistoryResponseAO1Result struct {

	// continuation
	Continuation Continuation `json:"continuation,omitempty"`

	// entities
	Entities []*StopOrderHistoryRecord `json:"entities"`
}

// Validate validates this private get stop order history response a o1 result
func (m *PrivateGetStopOrderHistoryResponseAO1Result) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContinuation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntities(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PrivateGetStopOrderHistoryResponseAO1Result) validateContinuation(formats strfmt.Registry) error {

	if swag.IsZero(m.Continuation) { // not required
		return nil
	}

	if err := m.Continuation.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("result" + "." + "continuation")
		}
		return err
	}

	return nil
}

func (m *PrivateGetStopOrderHistoryResponseAO1Result) validateEntities(formats strfmt.Registry) error {

	if swag.IsZero(m.Entities) { // not required
		return nil
	}

	for i := 0; i < len(m.Entities); i++ {
		if swag.IsZero(m.Entities[i]) { // not required
			continue
		}

		if m.Entities[i] != nil {
			if err := m.Entities[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("result" + "." + "entities" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PrivateGetStopOrderHistoryResponseAO1Result) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PrivateGetStopOrderHistoryResponseAO1Result) UnmarshalBinary(b []byte) error {
	var res PrivateGetStopOrderHistoryResponseAO1Result
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
