// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GcpCloudResourcesConfig gcp cloud resources config
//
// swagger:model gcp_cloud_resources_config
type GcpCloudResourcesConfig struct {

	// gcp network
	GcpNetwork *GcpCloudResourcesConfigGcpNetwork `json:"gcp_network,omitempty"`

	// gcp router
	GcpRouter *GcpCloudResourcesConfigGcpRouter `json:"gcp_router,omitempty"`
}

// Validate validates this gcp cloud resources config
func (m *GcpCloudResourcesConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGcpNetwork(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGcpRouter(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GcpCloudResourcesConfig) validateGcpNetwork(formats strfmt.Registry) error {
	if swag.IsZero(m.GcpNetwork) { // not required
		return nil
	}

	if m.GcpNetwork != nil {
		if err := m.GcpNetwork.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gcp_network")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("gcp_network")
			}
			return err
		}
	}

	return nil
}

func (m *GcpCloudResourcesConfig) validateGcpRouter(formats strfmt.Registry) error {
	if swag.IsZero(m.GcpRouter) { // not required
		return nil
	}

	if m.GcpRouter != nil {
		if err := m.GcpRouter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gcp_router")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("gcp_router")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this gcp cloud resources config based on the context it is used
func (m *GcpCloudResourcesConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateGcpNetwork(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGcpRouter(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GcpCloudResourcesConfig) contextValidateGcpNetwork(ctx context.Context, formats strfmt.Registry) error {

	if m.GcpNetwork != nil {
		if err := m.GcpNetwork.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gcp_network")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("gcp_network")
			}
			return err
		}
	}

	return nil
}

func (m *GcpCloudResourcesConfig) contextValidateGcpRouter(ctx context.Context, formats strfmt.Registry) error {

	if m.GcpRouter != nil {
		if err := m.GcpRouter.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gcp_router")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("gcp_router")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GcpCloudResourcesConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GcpCloudResourcesConfig) UnmarshalBinary(b []byte) error {
	var res GcpCloudResourcesConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// GcpCloudResourcesConfigGcpNetwork gcp cloud resources config gcp network
//
// swagger:model GcpCloudResourcesConfigGcpNetwork
type GcpCloudResourcesConfigGcpNetwork struct {

	// auto create
	AutoCreate bool `json:"auto_create"`

	// subnet
	// Min Length: 8
	Subnet string `json:"subnet,omitempty"`
}

// Validate validates this gcp cloud resources config gcp network
func (m *GcpCloudResourcesConfigGcpNetwork) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSubnet(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GcpCloudResourcesConfigGcpNetwork) validateSubnet(formats strfmt.Registry) error {
	if swag.IsZero(m.Subnet) { // not required
		return nil
	}

	if err := validate.MinLength("gcp_network"+"."+"subnet", "body", m.Subnet, 8); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this gcp cloud resources config gcp network based on context it is used
func (m *GcpCloudResourcesConfigGcpNetwork) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GcpCloudResourcesConfigGcpNetwork) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GcpCloudResourcesConfigGcpNetwork) UnmarshalBinary(b []byte) error {
	var res GcpCloudResourcesConfigGcpNetwork
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// GcpCloudResourcesConfigGcpRouter gcp cloud resources config gcp router
//
// swagger:model GcpCloudResourcesConfigGcpRouter
type GcpCloudResourcesConfigGcpRouter struct {

	// disable peer asn config
	DisablePeerAsnConfig bool `json:"disable_peer_asn_config"`

	// peer asn
	// Required: true
	PeerAsn *int64 `json:"peer_asn"`
}

// Validate validates this gcp cloud resources config gcp router
func (m *GcpCloudResourcesConfigGcpRouter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePeerAsn(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GcpCloudResourcesConfigGcpRouter) validatePeerAsn(formats strfmt.Registry) error {

	if err := validate.Required("gcp_router"+"."+"peer_asn", "body", m.PeerAsn); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this gcp cloud resources config gcp router based on context it is used
func (m *GcpCloudResourcesConfigGcpRouter) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GcpCloudResourcesConfigGcpRouter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GcpCloudResourcesConfigGcpRouter) UnmarshalBinary(b []byte) error {
	var res GcpCloudResourcesConfigGcpRouter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}