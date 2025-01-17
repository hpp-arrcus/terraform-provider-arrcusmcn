// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GcpInterface gcp interface
//
// swagger:model gcp_interface
type GcpInterface struct {

	// interface id
	InterfaceID string `json:"interface_id,omitempty"`

	// network
	Network string `json:"network,omitempty"`

	// route name
	RouteName string `json:"route_name,omitempty"`

	// subnetwork
	Subnetwork string `json:"subnetwork,omitempty"`
}

// Validate validates this gcp interface
func (m *GcpInterface) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this gcp interface based on context it is used
func (m *GcpInterface) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GcpInterface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GcpInterface) UnmarshalBinary(b []byte) error {
	var res GcpInterface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
