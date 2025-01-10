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

// AzureSku azure sku
//
// swagger:model azure_sku
type AzureSku string

func NewAzureSku(value AzureSku) *AzureSku {
	return &value
}

// Pointer returns a pointer to a freshly-allocated AzureSku.
func (m AzureSku) Pointer() *AzureSku {
	return &m
}

const (

	// AzureSkuBasic captures enum value "Basic"
	AzureSkuBasic AzureSku = "Basic"

	// AzureSkuStandard captures enum value "Standard"
	AzureSkuStandard AzureSku = "Standard"

	// AzureSkuHighPerformance captures enum value "HighPerformance"
	AzureSkuHighPerformance AzureSku = "HighPerformance"

	// AzureSkuUltraPerformance captures enum value "UltraPerformance"
	AzureSkuUltraPerformance AzureSku = "UltraPerformance"

	// AzureSkuErGw1AZ captures enum value "ErGw1AZ"
	AzureSkuErGw1AZ AzureSku = "ErGw1AZ"

	// AzureSkuErGw2AZ captures enum value "ErGw2AZ"
	AzureSkuErGw2AZ AzureSku = "ErGw2AZ"

	// AzureSkuErGw3AZ captures enum value "ErGw3AZ"
	AzureSkuErGw3AZ AzureSku = "ErGw3AZ"

	// AzureSkuVpnGw1 captures enum value "VpnGw1"
	AzureSkuVpnGw1 AzureSku = "VpnGw1"

	// AzureSkuVpnGw2 captures enum value "VpnGw2"
	AzureSkuVpnGw2 AzureSku = "VpnGw2"

	// AzureSkuVpnGw3 captures enum value "VpnGw3"
	AzureSkuVpnGw3 AzureSku = "VpnGw3"

	// AzureSkuVpnGw4 captures enum value "VpnGw4"
	AzureSkuVpnGw4 AzureSku = "VpnGw4"

	// AzureSkuVpnGw1AZ captures enum value "VpnGw1AZ"
	AzureSkuVpnGw1AZ AzureSku = "VpnGw1AZ"

	// AzureSkuVpnGw2AZ captures enum value "VpnGw2AZ"
	AzureSkuVpnGw2AZ AzureSku = "VpnGw2AZ"

	// AzureSkuVpnGw3AZ captures enum value "VpnGw3AZ"
	AzureSkuVpnGw3AZ AzureSku = "VpnGw3AZ"

	// AzureSkuVpnGw4AZ captures enum value "VpnGw4AZ"
	AzureSkuVpnGw4AZ AzureSku = "VpnGw4AZ"

	// AzureSkuVpnGw5AZVpnGw5 captures enum value "VpnGw5AZVpnGw5"
	AzureSkuVpnGw5AZVpnGw5 AzureSku = "VpnGw5AZVpnGw5"
)

// for schema
var azureSkuEnum []interface{}

func init() {
	var res []AzureSku
	if err := json.Unmarshal([]byte(`["Basic","Standard","HighPerformance","UltraPerformance","ErGw1AZ","ErGw2AZ","ErGw3AZ","VpnGw1","VpnGw2","VpnGw3","VpnGw4","VpnGw1AZ","VpnGw2AZ","VpnGw3AZ","VpnGw4AZ","VpnGw5AZVpnGw5"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		azureSkuEnum = append(azureSkuEnum, v)
	}
}

func (m AzureSku) validateAzureSkuEnum(path, location string, value AzureSku) error {
	if err := validate.EnumCase(path, location, value, azureSkuEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this azure sku
func (m AzureSku) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateAzureSkuEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this azure sku based on context it is used
func (m AzureSku) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}