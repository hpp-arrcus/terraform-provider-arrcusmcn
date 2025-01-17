// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AwsInterface aws interface
//
// swagger:model aws_interface
type AwsInterface struct {

	// interface id
	InterfaceID string `json:"interface_id,omitempty"`

	// route table id
	RouteTableID string `json:"route_table_id,omitempty"`

	// security groups
	SecurityGroups []string `json:"security_groups"`

	// subnet id
	SubnetID string `json:"subnet_id,omitempty"`
}

// Validate validates this aws interface
func (m *AwsInterface) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this aws interface based on context it is used
func (m *AwsInterface) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AwsInterface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AwsInterface) UnmarshalBinary(b []byte) error {
	var res AwsInterface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
