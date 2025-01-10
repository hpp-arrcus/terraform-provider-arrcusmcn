// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AwsDeployment aws deployment
//
// swagger:model aws_deployment
type AwsDeployment struct {

	// assign public ip
	AssignPublicIP bool `json:"assign_public_ip,omitempty"`

	// backup zone
	BackupZone string `json:"backup_zone,omitempty"`

	// byoip
	// Min Length: 1
	Byoip string `json:"byoip,omitempty"`

	// byoip backup
	// Min Length: 1
	ByoipBackup string `json:"byoip_backup,omitempty"`

	// instance key
	// Min Length: 1
	InstanceKey string `json:"instance_key,omitempty"`

	// instance type
	// Min Length: 1
	InstanceType string `json:"instance_type,omitempty"`

	// networks
	// Min Items: 1
	// Unique: true
	Networks []*AwsNetwork `json:"networks"`

	// region
	// Min Length: 1
	Region string `json:"region,omitempty"`

	// vpc cidr block
	// Min Length: 1
	VpcCidrBlock string `json:"vpc_cidr_block,omitempty"`

	// vpc id
	// Min Length: 1
	VpcID string `json:"vpc_id,omitempty"`

	// zone
	// Min Length: 1
	Zone string `json:"zone,omitempty"`
}

// Validate validates this aws deployment
func (m *AwsDeployment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateByoip(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateByoipBackup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstanceKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstanceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVpcCidrBlock(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVpcID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateZone(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AwsDeployment) validateByoip(formats strfmt.Registry) error {
	if swag.IsZero(m.Byoip) { // not required
		return nil
	}

	if err := validate.MinLength("byoip", "body", m.Byoip, 1); err != nil {
		return err
	}

	return nil
}

func (m *AwsDeployment) validateByoipBackup(formats strfmt.Registry) error {
	if swag.IsZero(m.ByoipBackup) { // not required
		return nil
	}

	if err := validate.MinLength("byoip_backup", "body", m.ByoipBackup, 1); err != nil {
		return err
	}

	return nil
}

func (m *AwsDeployment) validateInstanceKey(formats strfmt.Registry) error {
	if swag.IsZero(m.InstanceKey) { // not required
		return nil
	}

	if err := validate.MinLength("instance_key", "body", m.InstanceKey, 1); err != nil {
		return err
	}

	return nil
}

func (m *AwsDeployment) validateInstanceType(formats strfmt.Registry) error {
	if swag.IsZero(m.InstanceType) { // not required
		return nil
	}

	if err := validate.MinLength("instance_type", "body", m.InstanceType, 1); err != nil {
		return err
	}

	return nil
}

func (m *AwsDeployment) validateNetworks(formats strfmt.Registry) error {
	if swag.IsZero(m.Networks) { // not required
		return nil
	}

	iNetworksSize := int64(len(m.Networks))

	if err := validate.MinItems("networks", "body", iNetworksSize, 1); err != nil {
		return err
	}

	if err := validate.UniqueItems("networks", "body", m.Networks); err != nil {
		return err
	}

	for i := 0; i < len(m.Networks); i++ {
		if swag.IsZero(m.Networks[i]) { // not required
			continue
		}

		if m.Networks[i] != nil {
			if err := m.Networks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("networks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("networks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AwsDeployment) validateRegion(formats strfmt.Registry) error {
	if swag.IsZero(m.Region) { // not required
		return nil
	}

	if err := validate.MinLength("region", "body", m.Region, 1); err != nil {
		return err
	}

	return nil
}

func (m *AwsDeployment) validateVpcCidrBlock(formats strfmt.Registry) error {
	if swag.IsZero(m.VpcCidrBlock) { // not required
		return nil
	}

	if err := validate.MinLength("vpc_cidr_block", "body", m.VpcCidrBlock, 1); err != nil {
		return err
	}

	return nil
}

func (m *AwsDeployment) validateVpcID(formats strfmt.Registry) error {
	if swag.IsZero(m.VpcID) { // not required
		return nil
	}

	if err := validate.MinLength("vpc_id", "body", m.VpcID, 1); err != nil {
		return err
	}

	return nil
}

func (m *AwsDeployment) validateZone(formats strfmt.Registry) error {
	if swag.IsZero(m.Zone) { // not required
		return nil
	}

	if err := validate.MinLength("zone", "body", m.Zone, 1); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this aws deployment based on the context it is used
func (m *AwsDeployment) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNetworks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AwsDeployment) contextValidateNetworks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Networks); i++ {

		if m.Networks[i] != nil {
			if err := m.Networks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("networks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("networks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AwsDeployment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AwsDeployment) UnmarshalBinary(b []byte) error {
	var res AwsDeployment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}