// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ArcCostProbeEntry arc cost probe entry
//
// swagger:model arc_cost_probe_entry
type ArcCostProbeEntry struct {

	// deployments
	Deployments []*ArcCostInput `json:"deployments"`

	// id
	// Read Only: true
	// Min Length: 1
	// Format: ObjectId
	ID *strfmt.ObjectId `json:"id,omitempty" bson:"_id, omitempty"`

	// probename
	Probename string `json:"probename,omitempty"`
}

// Validate validates this arc cost probe entry
func (m *ArcCostProbeEntry) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeployments(formats); err != nil {
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

func (m *ArcCostProbeEntry) validateDeployments(formats strfmt.Registry) error {
	if swag.IsZero(m.Deployments) { // not required
		return nil
	}

	for i := 0; i < len(m.Deployments); i++ {
		if swag.IsZero(m.Deployments[i]) { // not required
			continue
		}

		if m.Deployments[i] != nil {
			if err := m.Deployments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("deployments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("deployments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ArcCostProbeEntry) validateID(formats strfmt.Registry) error {
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

// ContextValidate validate this arc cost probe entry based on the context it is used
func (m *ArcCostProbeEntry) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDeployments(ctx, formats); err != nil {
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

func (m *ArcCostProbeEntry) contextValidateDeployments(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Deployments); i++ {

		if m.Deployments[i] != nil {
			if err := m.Deployments[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("deployments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("deployments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ArcCostProbeEntry) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ArcCostProbeEntry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ArcCostProbeEntry) UnmarshalBinary(b []byte) error {
	var res ArcCostProbeEntry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
