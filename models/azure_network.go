// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AzureNetwork azure network
//
// swagger:model azure_network
type AzureNetwork struct {

	// name
	Name string `json:"name,omitempty"`

	// subnetwork
	Subnetwork string `json:"subnetwork,omitempty"`
}

// Validate validates this azure network
func (m *AzureNetwork) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this azure network based on context it is used
func (m *AzureNetwork) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AzureNetwork) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AzureNetwork) UnmarshalBinary(b []byte) error {
	var res AzureNetwork
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}