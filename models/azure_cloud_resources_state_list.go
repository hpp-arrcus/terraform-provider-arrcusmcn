// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AzureCloudResourcesStateList azure cloud resources state list
//
// swagger:model azure_cloud_resources_state_list
type AzureCloudResourcesStateList struct {

	// express route circuit
	ExpressRouteCircuit *AzureCloudResourcesStateListExpressRouteCircuit `json:"express_route_circuit,omitempty"`

	// gateway
	Gateway *AzureCloudResourcesStateListGateway `json:"gateway,omitempty"`

	// gateway subnet
	GatewaySubnet string `json:"gateway_subnet,omitempty"`

	// public ip
	PublicIP *AzureCloudResourcesStateListPublicIP `json:"public_ip,omitempty"`

	// resource group
	ResourceGroup string `json:"resource_group,omitempty"`

	// route association
	RouteAssociation string `json:"route_association,omitempty"`

	// route table
	RouteTable string `json:"route_table,omitempty"`

	// virtual network
	VirtualNetwork string `json:"virtual_network,omitempty"`
}

// Validate validates this azure cloud resources state list
func (m *AzureCloudResourcesStateList) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExpressRouteCircuit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGateway(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePublicIP(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AzureCloudResourcesStateList) validateExpressRouteCircuit(formats strfmt.Registry) error {
	if swag.IsZero(m.ExpressRouteCircuit) { // not required
		return nil
	}

	if m.ExpressRouteCircuit != nil {
		if err := m.ExpressRouteCircuit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("express_route_circuit")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("express_route_circuit")
			}
			return err
		}
	}

	return nil
}

func (m *AzureCloudResourcesStateList) validateGateway(formats strfmt.Registry) error {
	if swag.IsZero(m.Gateway) { // not required
		return nil
	}

	if m.Gateway != nil {
		if err := m.Gateway.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gateway")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("gateway")
			}
			return err
		}
	}

	return nil
}

func (m *AzureCloudResourcesStateList) validatePublicIP(formats strfmt.Registry) error {
	if swag.IsZero(m.PublicIP) { // not required
		return nil
	}

	if m.PublicIP != nil {
		if err := m.PublicIP.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("public_ip")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("public_ip")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this azure cloud resources state list based on the context it is used
func (m *AzureCloudResourcesStateList) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateExpressRouteCircuit(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGateway(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePublicIP(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AzureCloudResourcesStateList) contextValidateExpressRouteCircuit(ctx context.Context, formats strfmt.Registry) error {

	if m.ExpressRouteCircuit != nil {
		if err := m.ExpressRouteCircuit.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("express_route_circuit")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("express_route_circuit")
			}
			return err
		}
	}

	return nil
}

func (m *AzureCloudResourcesStateList) contextValidateGateway(ctx context.Context, formats strfmt.Registry) error {

	if m.Gateway != nil {
		if err := m.Gateway.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gateway")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("gateway")
			}
			return err
		}
	}

	return nil
}

func (m *AzureCloudResourcesStateList) contextValidatePublicIP(ctx context.Context, formats strfmt.Registry) error {

	if m.PublicIP != nil {
		if err := m.PublicIP.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("public_ip")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("public_ip")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AzureCloudResourcesStateList) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AzureCloudResourcesStateList) UnmarshalBinary(b []byte) error {
	var res AzureCloudResourcesStateList
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AzureCloudResourcesStateListExpressRouteCircuit azure cloud resources state list express route circuit
//
// swagger:model AzureCloudResourcesStateListExpressRouteCircuit
type AzureCloudResourcesStateListExpressRouteCircuit struct {

	// auth id
	AuthID string `json:"auth_id,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// peering id
	PeeringID string `json:"peering_id,omitempty"`

	// service key
	ServiceKey string `json:"service_key,omitempty"`
}

// Validate validates this azure cloud resources state list express route circuit
func (m *AzureCloudResourcesStateListExpressRouteCircuit) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this azure cloud resources state list express route circuit based on context it is used
func (m *AzureCloudResourcesStateListExpressRouteCircuit) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AzureCloudResourcesStateListExpressRouteCircuit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AzureCloudResourcesStateListExpressRouteCircuit) UnmarshalBinary(b []byte) error {
	var res AzureCloudResourcesStateListExpressRouteCircuit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AzureCloudResourcesStateListGateway azure cloud resources state list gateway
//
// swagger:model AzureCloudResourcesStateListGateway
type AzureCloudResourcesStateListGateway struct {

	// connection id
	ConnectionID string `json:"connection_id,omitempty"`

	// id
	ID string `json:"id,omitempty"`
}

// Validate validates this azure cloud resources state list gateway
func (m *AzureCloudResourcesStateListGateway) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this azure cloud resources state list gateway based on context it is used
func (m *AzureCloudResourcesStateListGateway) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AzureCloudResourcesStateListGateway) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AzureCloudResourcesStateListGateway) UnmarshalBinary(b []byte) error {
	var res AzureCloudResourcesStateListGateway
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AzureCloudResourcesStateListPublicIP azure cloud resources state list public IP
//
// swagger:model AzureCloudResourcesStateListPublicIP
type AzureCloudResourcesStateListPublicIP struct {

	// address
	Address string `json:"address,omitempty"`

	// id
	ID string `json:"id,omitempty"`
}

// Validate validates this azure cloud resources state list public IP
func (m *AzureCloudResourcesStateListPublicIP) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this azure cloud resources state list public IP based on context it is used
func (m *AzureCloudResourcesStateListPublicIP) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AzureCloudResourcesStateListPublicIP) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AzureCloudResourcesStateListPublicIP) UnmarshalBinary(b []byte) error {
	var res AzureCloudResourcesStateListPublicIP
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
