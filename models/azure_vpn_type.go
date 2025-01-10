// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// AzureVpnType azure vpn type
//
// swagger:model azure_vpn_type
type AzureVpnType string

func NewAzureVpnType(value AzureVpnType) *AzureVpnType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated AzureVpnType.
func (m AzureVpnType) Pointer() *AzureVpnType {
	return &m
}

const (

	// AzureVpnTypeRouteBased captures enum value "RouteBased"
	AzureVpnTypeRouteBased AzureVpnType = "RouteBased"

	// AzureVpnTypePolicyBased captures enum value "PolicyBased"
	AzureVpnTypePolicyBased AzureVpnType = "PolicyBased"
)

// for schema
var azureVpnTypeEnum []interface{}

func init() {
	var res []AzureVpnType
	if err := json.Unmarshal([]byte(`["RouteBased","PolicyBased"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		azureVpnTypeEnum = append(azureVpnTypeEnum, v)
	}
}

func (m AzureVpnType) validateAzureVpnTypeEnum(path, location string, value AzureVpnType) error {
	if err := validate.EnumCase(path, location, value, azureVpnTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this azure vpn type
func (m AzureVpnType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateAzureVpnTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this azure vpn type based on context it is used
func (m AzureVpnType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
