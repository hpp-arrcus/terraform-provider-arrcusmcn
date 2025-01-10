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

// AzureDeployment azure deployment
//
// swagger:model azure_deployment
type AzureDeployment struct {

	// accelerated networking enabled
	AcceleratedNetworkingEnabled bool `json:"accelerated_networking_enabled,omitempty"`

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

	// enable accelerated networking
	EnableAcceleratedNetworking bool `json:"enable_accelerated_networking,omitempty"`

	// instance key
	InstanceKey *InstanceKey `json:"instance_key,omitempty"`

	// instance type
	// Min Length: 1
	InstanceType string `json:"instance_type,omitempty"`

	// networks
	// Min Items: 1
	// Unique: true
	Networks []*AzureNetwork `json:"networks"`

	// region
	// Min Length: 1
	Region string `json:"region,omitempty"`

	// resource group
	// Min Length: 1
	ResourceGroup string `json:"resource_group,omitempty"`

	// vnet
	Vnet string `json:"vnet,omitempty"`

	// zone
	Zone string `json:"zone,omitempty"`
}

// Validate validates this azure deployment
func (m *AzureDeployment) Validate(formats strfmt.Registry) error {
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

	if err := m.validateResourceGroup(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AzureDeployment) validateByoip(formats strfmt.Registry) error {
	if swag.IsZero(m.Byoip) { // not required
		return nil
	}

	if err := validate.MinLength("byoip", "body", m.Byoip, 1); err != nil {
		return err
	}

	return nil
}

func (m *AzureDeployment) validateByoipBackup(formats strfmt.Registry) error {
	if swag.IsZero(m.ByoipBackup) { // not required
		return nil
	}

	if err := validate.MinLength("byoip_backup", "body", m.ByoipBackup, 1); err != nil {
		return err
	}

	return nil
}

func (m *AzureDeployment) validateInstanceKey(formats strfmt.Registry) error {
	if swag.IsZero(m.InstanceKey) { // not required
		return nil
	}

	if m.InstanceKey != nil {
		if err := m.InstanceKey.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("instance_key")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("instance_key")
			}
			return err
		}
	}

	return nil
}

func (m *AzureDeployment) validateInstanceType(formats strfmt.Registry) error {
	if swag.IsZero(m.InstanceType) { // not required
		return nil
	}

	if err := validate.MinLength("instance_type", "body", m.InstanceType, 1); err != nil {
		return err
	}

	return nil
}

func (m *AzureDeployment) validateNetworks(formats strfmt.Registry) error {
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

func (m *AzureDeployment) validateRegion(formats strfmt.Registry) error {
	if swag.IsZero(m.Region) { // not required
		return nil
	}

	if err := validate.MinLength("region", "body", m.Region, 1); err != nil {
		return err
	}

	return nil
}

func (m *AzureDeployment) validateResourceGroup(formats strfmt.Registry) error {
	if swag.IsZero(m.ResourceGroup) { // not required
		return nil
	}

	if err := validate.MinLength("resource_group", "body", m.ResourceGroup, 1); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this azure deployment based on the context it is used
func (m *AzureDeployment) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateInstanceKey(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNetworks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AzureDeployment) contextValidateInstanceKey(ctx context.Context, formats strfmt.Registry) error {

	if m.InstanceKey != nil {
		if err := m.InstanceKey.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("instance_key")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("instance_key")
			}
			return err
		}
	}

	return nil
}

func (m *AzureDeployment) contextValidateNetworks(ctx context.Context, formats strfmt.Registry) error {

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
func (m *AzureDeployment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AzureDeployment) UnmarshalBinary(b []byte) error {
	var res AzureDeployment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
