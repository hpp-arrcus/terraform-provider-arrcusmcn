// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VsphereInterface vsphere interface
//
// swagger:model vsphere_interface
type VsphereInterface struct {

	// network
	Network string `json:"network,omitempty"`
}

// Validate validates this vsphere interface
func (m *VsphereInterface) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this vsphere interface based on context it is used
func (m *VsphereInterface) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VsphereInterface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VsphereInterface) UnmarshalBinary(b []byte) error {
	var res VsphereInterface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
