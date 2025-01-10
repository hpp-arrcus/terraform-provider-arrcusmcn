// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ArcCostVisibilityInput arc cost visibility input
//
// swagger:model arc_cost_visibility_input
type ArcCostVisibilityInput struct {

	// deployment name
	DeploymentName string `json:"deploymentName,omitempty"`

	// interface ids
	InterfaceIds []string `json:"interfaceIds"`

	// interface names
	InterfaceNames []string `json:"interfaceNames"`
}

// Validate validates this arc cost visibility input
func (m *ArcCostVisibilityInput) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this arc cost visibility input based on context it is used
func (m *ArcCostVisibilityInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ArcCostVisibilityInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ArcCostVisibilityInput) UnmarshalBinary(b []byte) error {
	var res ArcCostVisibilityInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
