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

// Connections connections
//
// swagger:model connections
type Connections struct {

	// deployment1
	// Required: true
	Deployment1 *Components `json:"deployment1"`

	// deployment2
	// Required: true
	Deployment2 *Components `json:"deployment2"`

	// id
	// Read Only: true
	// Format: ObjectId
	ID *strfmt.ObjectId `json:"id,omitempty" bson:"_id, omitempty"`

	// name
	// Required: true
	// Min Length: 1
	Name *string `json:"name"`
}

// Validate validates this connections
func (m *Connections) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeployment1(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeployment2(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Connections) validateDeployment1(formats strfmt.Registry) error {

	if err := validate.Required("deployment1", "body", m.Deployment1); err != nil {
		return err
	}

	if m.Deployment1 != nil {
		if err := m.Deployment1.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deployment1")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("deployment1")
			}
			return err
		}
	}

	return nil
}

func (m *Connections) validateDeployment2(formats strfmt.Registry) error {

	if err := validate.Required("deployment2", "body", m.Deployment2); err != nil {
		return err
	}

	if m.Deployment2 != nil {
		if err := m.Deployment2.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deployment2")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("deployment2")
			}
			return err
		}
	}

	return nil
}

func (m *Connections) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "ObjectId", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Connections) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", *m.Name, 1); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this connections based on the context it is used
func (m *Connections) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDeployment1(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDeployment2(ctx, formats); err != nil {
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

func (m *Connections) contextValidateDeployment1(ctx context.Context, formats strfmt.Registry) error {

	if m.Deployment1 != nil {
		if err := m.Deployment1.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deployment1")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("deployment1")
			}
			return err
		}
	}

	return nil
}

func (m *Connections) contextValidateDeployment2(ctx context.Context, formats strfmt.Registry) error {

	if m.Deployment2 != nil {
		if err := m.Deployment2.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deployment2")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("deployment2")
			}
			return err
		}
	}

	return nil
}

func (m *Connections) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Connections) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Connections) UnmarshalBinary(b []byte) error {
	var res Connections
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
