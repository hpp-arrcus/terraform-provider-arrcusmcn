// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Zerotrust zerotrust
//
// swagger:model zerotrust
type Zerotrust struct {

	// enable
	// Required: true
	// Read Only: true
	Enable bool `json:"enable"`

	// id
	// Read Only: true
	// Min Length: 1
	// Format: ObjectId
	ID *strfmt.ObjectId `json:"id,omitempty" bson:"_id, omitempty"`
}

// Validate validates this zerotrust
func (m *Zerotrust) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnable(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Zerotrust) validateEnable(formats strfmt.Registry) error {

	if err := validate.Required("enable", "body", bool(m.Enable)); err != nil {
		return err
	}

	return nil
}

func (m *Zerotrust) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.MinLength("id", "body", m.ID.String(), 1); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "ObjectId", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this zerotrust based on the context it is used
func (m *Zerotrust) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEnable(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Zerotrust) contextValidateEnable(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "enable", "body", bool(m.Enable)); err != nil {
		return err
	}

	return nil
}

func (m *Zerotrust) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Zerotrust) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Zerotrust) UnmarshalBinary(b []byte) error {
	var res Zerotrust
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}