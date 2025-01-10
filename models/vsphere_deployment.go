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

// VsphereDeployment vsphere deployment
//
// swagger:model vsphere_deployment
type VsphereDeployment struct {

	// admin ssh public key
	AdminSSHPublicKey *InstanceKey `json:"admin_ssh_public_key,omitempty"`

	// assign public ip
	AssignPublicIP bool `json:"assign_public_ip,omitempty"`

	// cluster
	// Min Length: 1
	Cluster string `json:"cluster,omitempty"`

	// content library
	// Min Length: 1
	ContentLibrary string `json:"content_library,omitempty"`

	// datacenter
	// Min Length: 1
	Datacenter string `json:"datacenter,omitempty"`

	// datastore
	// Min Length: 1
	Datastore string `json:"datastore,omitempty"`

	// default gw
	// Min Length: 1
	DefaultGw string `json:"default_gw,omitempty"`

	// disk size
	DiskSize int64 `json:"disk_size,omitempty"`

	// networks
	Networks []*VsphereNetwork `json:"networks"`

	// private ip
	// Min Length: 1
	PrivateIP string `json:"private_ip,omitempty"`

	// public ip
	// Min Length: 1
	PublicIP string `json:"public_ip,omitempty"`

	// ssh psw
	SSHPsw string `json:"ssh_psw,omitempty"`

	// subnet mask
	SubnetMask int64 `json:"subnet_mask,omitempty"`

	// vcpus
	Vcpus int64 `json:"vcpus,omitempty"`

	// vm memory
	VMMemory int64 `json:"vm_memory,omitempty"`

	// vsphere uuid
	VsphereUUID string `json:"vsphere_uuid,omitempty"`

	// vspherehost
	// Min Length: 1
	Vspherehost string `json:"vspherehost,omitempty"`
}

// Validate validates this vsphere deployment
func (m *VsphereDeployment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdminSSHPublicKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCluster(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContentLibrary(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDatacenter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDatastore(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDefaultGw(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrivateIP(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePublicIP(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVspherehost(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VsphereDeployment) validateAdminSSHPublicKey(formats strfmt.Registry) error {
	if swag.IsZero(m.AdminSSHPublicKey) { // not required
		return nil
	}

	if m.AdminSSHPublicKey != nil {
		if err := m.AdminSSHPublicKey.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("admin_ssh_public_key")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("admin_ssh_public_key")
			}
			return err
		}
	}

	return nil
}

func (m *VsphereDeployment) validateCluster(formats strfmt.Registry) error {
	if swag.IsZero(m.Cluster) { // not required
		return nil
	}

	if err := validate.MinLength("cluster", "body", m.Cluster, 1); err != nil {
		return err
	}

	return nil
}

func (m *VsphereDeployment) validateContentLibrary(formats strfmt.Registry) error {
	if swag.IsZero(m.ContentLibrary) { // not required
		return nil
	}

	if err := validate.MinLength("content_library", "body", m.ContentLibrary, 1); err != nil {
		return err
	}

	return nil
}

func (m *VsphereDeployment) validateDatacenter(formats strfmt.Registry) error {
	if swag.IsZero(m.Datacenter) { // not required
		return nil
	}

	if err := validate.MinLength("datacenter", "body", m.Datacenter, 1); err != nil {
		return err
	}

	return nil
}

func (m *VsphereDeployment) validateDatastore(formats strfmt.Registry) error {
	if swag.IsZero(m.Datastore) { // not required
		return nil
	}

	if err := validate.MinLength("datastore", "body", m.Datastore, 1); err != nil {
		return err
	}

	return nil
}

func (m *VsphereDeployment) validateDefaultGw(formats strfmt.Registry) error {
	if swag.IsZero(m.DefaultGw) { // not required
		return nil
	}

	if err := validate.MinLength("default_gw", "body", m.DefaultGw, 1); err != nil {
		return err
	}

	return nil
}

func (m *VsphereDeployment) validateNetworks(formats strfmt.Registry) error {
	if swag.IsZero(m.Networks) { // not required
		return nil
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

func (m *VsphereDeployment) validatePrivateIP(formats strfmt.Registry) error {
	if swag.IsZero(m.PrivateIP) { // not required
		return nil
	}

	if err := validate.MinLength("private_ip", "body", m.PrivateIP, 1); err != nil {
		return err
	}

	return nil
}

func (m *VsphereDeployment) validatePublicIP(formats strfmt.Registry) error {
	if swag.IsZero(m.PublicIP) { // not required
		return nil
	}

	if err := validate.MinLength("public_ip", "body", m.PublicIP, 1); err != nil {
		return err
	}

	return nil
}

func (m *VsphereDeployment) validateVspherehost(formats strfmt.Registry) error {
	if swag.IsZero(m.Vspherehost) { // not required
		return nil
	}

	if err := validate.MinLength("vspherehost", "body", m.Vspherehost, 1); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this vsphere deployment based on the context it is used
func (m *VsphereDeployment) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAdminSSHPublicKey(ctx, formats); err != nil {
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

func (m *VsphereDeployment) contextValidateAdminSSHPublicKey(ctx context.Context, formats strfmt.Registry) error {

	if m.AdminSSHPublicKey != nil {
		if err := m.AdminSSHPublicKey.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("admin_ssh_public_key")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("admin_ssh_public_key")
			}
			return err
		}
	}

	return nil
}

func (m *VsphereDeployment) contextValidateNetworks(ctx context.Context, formats strfmt.Registry) error {

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
func (m *VsphereDeployment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VsphereDeployment) UnmarshalBinary(b []byte) error {
	var res VsphereDeployment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}