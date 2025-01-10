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

// TerraformOperation terraform operation
//
// swagger:model TerraformOperation
type TerraformOperation string

func NewTerraformOperation(value TerraformOperation) *TerraformOperation {
	return &value
}

// Pointer returns a pointer to a freshly-allocated TerraformOperation.
func (m TerraformOperation) Pointer() *TerraformOperation {
	return &m
}

const (

	// TerraformOperationCreate captures enum value "Create"
	TerraformOperationCreate TerraformOperation = "Create"

	// TerraformOperationUpdate captures enum value "Update"
	TerraformOperationUpdate TerraformOperation = "Update"

	// TerraformOperationDelete captures enum value "Delete"
	TerraformOperationDelete TerraformOperation = "Delete"

	// TerraformOperationShow captures enum value "Show"
	TerraformOperationShow TerraformOperation = "Show"

	// TerraformOperationApply captures enum value "Apply"
	TerraformOperationApply TerraformOperation = "Apply"

	// TerraformOperationDestroy captures enum value "Destroy"
	TerraformOperationDestroy TerraformOperation = "Destroy"
)

// for schema
var terraformOperationEnum []interface{}

func init() {
	var res []TerraformOperation
	if err := json.Unmarshal([]byte(`["Create","Update","Delete","Show","Apply","Destroy"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		terraformOperationEnum = append(terraformOperationEnum, v)
	}
}

func (m TerraformOperation) validateTerraformOperationEnum(path, location string, value TerraformOperation) error {
	if err := validate.EnumCase(path, location, value, terraformOperationEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this terraform operation
func (m TerraformOperation) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateTerraformOperationEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this terraform operation based on context it is used
func (m TerraformOperation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
