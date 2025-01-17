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

// ArcCostRegionCoordinates arc cost region coordinates
//
// swagger:model arc_cost_region_coordinates
type ArcCostRegionCoordinates struct {

	// coordinate
	Coordinate Coordinateslonglat `json:"Coordinate,omitempty"`

	// coordinate Json
	CoordinateJSON *Coordinates `json:"CoordinateJson,omitempty"`
}

// Validate validates this arc cost region coordinates
func (m *ArcCostRegionCoordinates) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCoordinate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCoordinateJSON(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ArcCostRegionCoordinates) validateCoordinate(formats strfmt.Registry) error {
	if swag.IsZero(m.Coordinate) { // not required
		return nil
	}

	if err := m.Coordinate.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Coordinate")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("Coordinate")
		}
		return err
	}

	return nil
}

func (m *ArcCostRegionCoordinates) validateCoordinateJSON(formats strfmt.Registry) error {
	if swag.IsZero(m.CoordinateJSON) { // not required
		return nil
	}

	if m.CoordinateJSON != nil {
		if err := m.CoordinateJSON.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("CoordinateJson")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("CoordinateJson")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this arc cost region coordinates based on the context it is used
func (m *ArcCostRegionCoordinates) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCoordinate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCoordinateJSON(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ArcCostRegionCoordinates) contextValidateCoordinate(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Coordinate.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Coordinate")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("Coordinate")
		}
		return err
	}

	return nil
}

func (m *ArcCostRegionCoordinates) contextValidateCoordinateJSON(ctx context.Context, formats strfmt.Registry) error {

	if m.CoordinateJSON != nil {
		if err := m.CoordinateJSON.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("CoordinateJson")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("CoordinateJson")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ArcCostRegionCoordinates) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ArcCostRegionCoordinates) UnmarshalBinary(b []byte) error {
	var res ArcCostRegionCoordinates
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
