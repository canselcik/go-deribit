// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BookNotificationRaw book notification raw
// swagger:model book_notification_raw
type BookNotificationRaw struct {

	// asks
	// Required: true
	Asks [][]interface{} `json:"asks"`

	// bids
	// Required: true
	Bids [][]interface{} `json:"bids"`

	// Identifier of the notification
	// Required: true
	ChangeID *int64 `json:"change_id"`

	// instrument name
	// Required: true
	InstrumentName InstrumentName `json:"instrument_name"`

	// Identifier of the previous notification (it's **not** included for the first notification)
	PrevChangeID int64 `json:"prev_change_id,omitempty"`

	// timestamp
	Timestamp TimestampForBookNotifications `json:"timestamp,omitempty"`

	// Type of notification: `snapshot` for initial, `change` for others. The field is only included in aggregated response (when input parameter `interval` != `raw`)
	// Enum: [snapshot change]
	Type string `json:"type,omitempty"`
}

// Validate validates this book notification raw
func (m *BookNotificationRaw) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAsks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBids(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChangeID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstrumentName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BookNotificationRaw) validateAsks(formats strfmt.Registry) error {

	if err := validate.Required("asks", "body", m.Asks); err != nil {
		return err
	}

	return nil
}

func (m *BookNotificationRaw) validateBids(formats strfmt.Registry) error {

	if err := validate.Required("bids", "body", m.Bids); err != nil {
		return err
	}

	return nil
}

func (m *BookNotificationRaw) validateChangeID(formats strfmt.Registry) error {

	if err := validate.Required("change_id", "body", m.ChangeID); err != nil {
		return err
	}

	return nil
}

func (m *BookNotificationRaw) validateInstrumentName(formats strfmt.Registry) error {

	if err := m.InstrumentName.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("instrument_name")
		}
		return err
	}

	return nil
}

func (m *BookNotificationRaw) validateTimestamp(formats strfmt.Registry) error {

	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := m.Timestamp.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("timestamp")
		}
		return err
	}

	return nil
}

var bookNotificationRawTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["snapshot","change"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		bookNotificationRawTypeTypePropEnum = append(bookNotificationRawTypeTypePropEnum, v)
	}
}

const (

	// BookNotificationRawTypeSnapshot captures enum value "snapshot"
	BookNotificationRawTypeSnapshot string = "snapshot"

	// BookNotificationRawTypeChange captures enum value "change"
	BookNotificationRawTypeChange string = "change"
)

// prop value enum
func (m *BookNotificationRaw) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, bookNotificationRawTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *BookNotificationRaw) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BookNotificationRaw) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BookNotificationRaw) UnmarshalBinary(b []byte) error {
	var res BookNotificationRaw
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
