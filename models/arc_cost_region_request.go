// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ArcCostRegionRequest arc cost region request
//
// swagger:model arc_cost_region_request
type ArcCostRegionRequest struct {

	// csps
	Csps []string `json:"csps"`
}

// Validate validates this arc cost region request
func (m *ArcCostRegionRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this arc cost region request based on context it is used
func (m *ArcCostRegionRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ArcCostRegionRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ArcCostRegionRequest) UnmarshalBinary(b []byte) error {
	var res ArcCostRegionRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
