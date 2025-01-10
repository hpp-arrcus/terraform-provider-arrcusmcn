// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ArcCostLinkdata arc cost linkdata
//
// swagger:model arc_cost_linkdata
type ArcCostLinkdata struct {

	// atobcost
	Atobcost float64 `json:"atobcost,omitempty"`

	// btoacost
	Btoacost float64 `json:"btoacost,omitempty"`

	// vertex a
	VertexA *ArcCostVertex `json:"vertexA,omitempty"`

	// vertex b
	VertexB *ArcCostVertex `json:"vertexB,omitempty"`
}

// Validate validates this arc cost linkdata
func (m *ArcCostLinkdata) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVertexA(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVertexB(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ArcCostLinkdata) validateVertexA(formats strfmt.Registry) error {
	if swag.IsZero(m.VertexA) { // not required
		return nil
	}

	if m.VertexA != nil {
		if err := m.VertexA.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vertexA")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vertexA")
			}
			return err
		}
	}

	return nil
}

func (m *ArcCostLinkdata) validateVertexB(formats strfmt.Registry) error {
	if swag.IsZero(m.VertexB) { // not required
		return nil
	}

	if m.VertexB != nil {
		if err := m.VertexB.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vertexB")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vertexB")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this arc cost linkdata based on the context it is used
func (m *ArcCostLinkdata) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateVertexA(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVertexB(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ArcCostLinkdata) contextValidateVertexA(ctx context.Context, formats strfmt.Registry) error {

	if m.VertexA != nil {
		if err := m.VertexA.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vertexA")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vertexA")
			}
			return err
		}
	}

	return nil
}

func (m *ArcCostLinkdata) contextValidateVertexB(ctx context.Context, formats strfmt.Registry) error {

	if m.VertexB != nil {
		if err := m.VertexB.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vertexB")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vertexB")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ArcCostLinkdata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ArcCostLinkdata) UnmarshalBinary(b []byte) error {
	var res ArcCostLinkdata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
