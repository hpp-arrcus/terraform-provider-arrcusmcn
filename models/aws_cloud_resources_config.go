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

// AwsCloudResourcesConfig aws cloud resources config
//
// swagger:model aws_cloud_resources_config
type AwsCloudResourcesConfig struct {

	// direct connect
	DirectConnect *AwsCloudResourcesConfigDirectConnect `json:"direct_connect,omitempty"`

	// vif
	Vif *AwsCloudResourcesConfigVif `json:"vif,omitempty"`

	// vpc
	Vpc *AwsCloudResourcesConfigVpc `json:"vpc,omitempty"`
}

// Validate validates this aws cloud resources config
func (m *AwsCloudResourcesConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDirectConnect(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVif(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVpc(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AwsCloudResourcesConfig) validateDirectConnect(formats strfmt.Registry) error {
	if swag.IsZero(m.DirectConnect) { // not required
		return nil
	}

	if m.DirectConnect != nil {
		if err := m.DirectConnect.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("direct_connect")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("direct_connect")
			}
			return err
		}
	}

	return nil
}

func (m *AwsCloudResourcesConfig) validateVif(formats strfmt.Registry) error {
	if swag.IsZero(m.Vif) { // not required
		return nil
	}

	if m.Vif != nil {
		if err := m.Vif.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vif")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vif")
			}
			return err
		}
	}

	return nil
}

func (m *AwsCloudResourcesConfig) validateVpc(formats strfmt.Registry) error {
	if swag.IsZero(m.Vpc) { // not required
		return nil
	}

	if m.Vpc != nil {
		if err := m.Vpc.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vpc")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vpc")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this aws cloud resources config based on the context it is used
func (m *AwsCloudResourcesConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDirectConnect(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVif(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVpc(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AwsCloudResourcesConfig) contextValidateDirectConnect(ctx context.Context, formats strfmt.Registry) error {

	if m.DirectConnect != nil {
		if err := m.DirectConnect.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("direct_connect")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("direct_connect")
			}
			return err
		}
	}

	return nil
}

func (m *AwsCloudResourcesConfig) contextValidateVif(ctx context.Context, formats strfmt.Registry) error {

	if m.Vif != nil {
		if err := m.Vif.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vif")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vif")
			}
			return err
		}
	}

	return nil
}

func (m *AwsCloudResourcesConfig) contextValidateVpc(ctx context.Context, formats strfmt.Registry) error {

	if m.Vpc != nil {
		if err := m.Vpc.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vpc")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vpc")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AwsCloudResourcesConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AwsCloudResourcesConfig) UnmarshalBinary(b []byte) error {
	var res AwsCloudResourcesConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AwsCloudResourcesConfigDirectConnect aws cloud resources config direct connect
//
// swagger:model AwsCloudResourcesConfigDirectConnect
type AwsCloudResourcesConfigDirectConnect struct {

	// accept hosted connection
	AcceptHostedConnection bool `json:"accept_hosted_connection"`

	// amazon side asn
	// Required: true
	AmazonSideAsn *int64 `json:"amazon_side_asn"`

	// direct connect gateway
	DirectConnectGateway bool `json:"direct_connect_gateway"`

	// hosted connection id
	// Min Length: 8
	HostedConnectionID string `json:"hosted_connection_id,omitempty"`
}

// Validate validates this aws cloud resources config direct connect
func (m *AwsCloudResourcesConfigDirectConnect) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmazonSideAsn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHostedConnectionID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AwsCloudResourcesConfigDirectConnect) validateAmazonSideAsn(formats strfmt.Registry) error {

	if err := validate.Required("direct_connect"+"."+"amazon_side_asn", "body", m.AmazonSideAsn); err != nil {
		return err
	}

	return nil
}

func (m *AwsCloudResourcesConfigDirectConnect) validateHostedConnectionID(formats strfmt.Registry) error {
	if swag.IsZero(m.HostedConnectionID) { // not required
		return nil
	}

	if err := validate.MinLength("direct_connect"+"."+"hosted_connection_id", "body", m.HostedConnectionID, 8); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this aws cloud resources config direct connect based on context it is used
func (m *AwsCloudResourcesConfigDirectConnect) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AwsCloudResourcesConfigDirectConnect) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AwsCloudResourcesConfigDirectConnect) UnmarshalBinary(b []byte) error {
	var res AwsCloudResourcesConfigDirectConnect
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AwsCloudResourcesConfigVif aws cloud resources config vif
//
// swagger:model AwsCloudResourcesConfigVif
type AwsCloudResourcesConfigVif struct {

	// address family
	// Required: true
	// Enum: [ipv4 ipv6]
	AddressFamily *string `json:"address_family"`

	// amazon address
	// Min Length: 1
	AmazonAddress string `json:"amazon_address,omitempty"`

	// bgp asn
	BgpAsn int64 `json:"bgp_asn,omitempty"`

	// customer address
	// Min Length: 1
	CustomerAddress string `json:"customer_address,omitempty"`

	// vlan id
	// Maximum: 4096
	// Minimum: 1
	VlanID int64 `json:"vlan_id,omitempty"`
}

// Validate validates this aws cloud resources config vif
func (m *AwsCloudResourcesConfigVif) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddressFamily(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAmazonAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCustomerAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVlanID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var awsCloudResourcesConfigVifTypeAddressFamilyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ipv4","ipv6"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		awsCloudResourcesConfigVifTypeAddressFamilyPropEnum = append(awsCloudResourcesConfigVifTypeAddressFamilyPropEnum, v)
	}
}

const (

	// AwsCloudResourcesConfigVifAddressFamilyIPV4 captures enum value "ipv4"
	AwsCloudResourcesConfigVifAddressFamilyIPV4 string = "ipv4"

	// AwsCloudResourcesConfigVifAddressFamilyIPV6 captures enum value "ipv6"
	AwsCloudResourcesConfigVifAddressFamilyIPV6 string = "ipv6"
)

// prop value enum
func (m *AwsCloudResourcesConfigVif) validateAddressFamilyEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, awsCloudResourcesConfigVifTypeAddressFamilyPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AwsCloudResourcesConfigVif) validateAddressFamily(formats strfmt.Registry) error {

	if err := validate.Required("vif"+"."+"address_family", "body", m.AddressFamily); err != nil {
		return err
	}

	// value enum
	if err := m.validateAddressFamilyEnum("vif"+"."+"address_family", "body", *m.AddressFamily); err != nil {
		return err
	}

	return nil
}

func (m *AwsCloudResourcesConfigVif) validateAmazonAddress(formats strfmt.Registry) error {
	if swag.IsZero(m.AmazonAddress) { // not required
		return nil
	}

	if err := validate.MinLength("vif"+"."+"amazon_address", "body", m.AmazonAddress, 1); err != nil {
		return err
	}

	return nil
}

func (m *AwsCloudResourcesConfigVif) validateCustomerAddress(formats strfmt.Registry) error {
	if swag.IsZero(m.CustomerAddress) { // not required
		return nil
	}

	if err := validate.MinLength("vif"+"."+"customer_address", "body", m.CustomerAddress, 1); err != nil {
		return err
	}

	return nil
}

func (m *AwsCloudResourcesConfigVif) validateVlanID(formats strfmt.Registry) error {
	if swag.IsZero(m.VlanID) { // not required
		return nil
	}

	if err := validate.MinimumInt("vif"+"."+"vlan_id", "body", m.VlanID, 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("vif"+"."+"vlan_id", "body", m.VlanID, 4096, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this aws cloud resources config vif based on context it is used
func (m *AwsCloudResourcesConfigVif) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AwsCloudResourcesConfigVif) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AwsCloudResourcesConfigVif) UnmarshalBinary(b []byte) error {
	var res AwsCloudResourcesConfigVif
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AwsCloudResourcesConfigVpc aws cloud resources config vpc
//
// swagger:model AwsCloudResourcesConfigVpc
type AwsCloudResourcesConfigVpc struct {

	// cidr
	// Min Length: 1
	Cidr string `json:"cidr,omitempty"`

	// subnet
	// Min Length: 1
	Subnet string `json:"subnet,omitempty"`
}

// Validate validates this aws cloud resources config vpc
func (m *AwsCloudResourcesConfigVpc) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCidr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubnet(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AwsCloudResourcesConfigVpc) validateCidr(formats strfmt.Registry) error {
	if swag.IsZero(m.Cidr) { // not required
		return nil
	}

	if err := validate.MinLength("vpc"+"."+"cidr", "body", m.Cidr, 1); err != nil {
		return err
	}

	return nil
}

func (m *AwsCloudResourcesConfigVpc) validateSubnet(formats strfmt.Registry) error {
	if swag.IsZero(m.Subnet) { // not required
		return nil
	}

	if err := validate.MinLength("vpc"+"."+"subnet", "body", m.Subnet, 1); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this aws cloud resources config vpc based on context it is used
func (m *AwsCloudResourcesConfigVpc) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AwsCloudResourcesConfigVpc) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AwsCloudResourcesConfigVpc) UnmarshalBinary(b []byte) error {
	var res AwsCloudResourcesConfigVpc
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
