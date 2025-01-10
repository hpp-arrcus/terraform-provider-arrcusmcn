// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// OciNetwork oci network
//
// swagger:model oci_network
type OciNetwork struct {

	// subnet access
	// Enum: [public private]
	SubnetAccess string `json:"subnet_access,omitempty"`

	// subnet name
	SubnetName string `json:"subnet_name,omitempty"`

	// subnet ocid
	SubnetOcid string `json:"subnet_ocid,omitempty"`

	// vcn name
	VcnName string `json:"vcn_name,omitempty"`

	// vcn ocid
	VcnOcid string `json:"vcn_ocid,omitempty"`
}

// Validate validates this oci network
func (m *OciNetwork) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSubnetAccess(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var ociNetworkTypeSubnetAccessPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["public","private"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ociNetworkTypeSubnetAccessPropEnum = append(ociNetworkTypeSubnetAccessPropEnum, v)
	}
}

const (

	// OciNetworkSubnetAccessPublic captures enum value "public"
	OciNetworkSubnetAccessPublic string = "public"

	// OciNetworkSubnetAccessPrivate captures enum value "private"
	OciNetworkSubnetAccessPrivate string = "private"
)

// prop value enum
func (m *OciNetwork) validateSubnetAccessEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, ociNetworkTypeSubnetAccessPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *OciNetwork) validateSubnetAccess(formats strfmt.Registry) error {
	if swag.IsZero(m.SubnetAccess) { // not required
		return nil
	}

	// value enum
	if err := m.validateSubnetAccessEnum("subnet_access", "body", m.SubnetAccess); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this oci network based on context it is used
func (m *OciNetwork) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OciNetwork) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OciNetwork) UnmarshalBinary(b []byte) error {
	var res OciNetwork
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
