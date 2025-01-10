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

// ArcCostCsps arc cost csps
//
// swagger:model ArcCostCsps
type ArcCostCsps string

func NewArcCostCsps(value ArcCostCsps) *ArcCostCsps {
	return &value
}

// Pointer returns a pointer to a freshly-allocated ArcCostCsps.
func (m ArcCostCsps) Pointer() *ArcCostCsps {
	return &m
}

const (

	// ArcCostCspsAws captures enum value "aws"
	ArcCostCspsAws ArcCostCsps = "aws"

	// ArcCostCspsGcp captures enum value "gcp"
	ArcCostCspsGcp ArcCostCsps = "gcp"

	// ArcCostCspsAzure captures enum value "azure"
	ArcCostCspsAzure ArcCostCsps = "azure"
)

// for schema
var arcCostCspsEnum []interface{}

func init() {
	var res []ArcCostCsps
	if err := json.Unmarshal([]byte(`["aws","gcp","azure"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		arcCostCspsEnum = append(arcCostCspsEnum, v)
	}
}

func (m ArcCostCsps) validateArcCostCspsEnum(path, location string, value ArcCostCsps) error {
	if err := validate.EnumCase(path, location, value, arcCostCspsEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this arc cost csps
func (m ArcCostCsps) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateArcCostCspsEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this arc cost csps based on context it is used
func (m ArcCostCsps) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}