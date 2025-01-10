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

// TerraformLog terraform log
//
// swagger:model terraform_log
type TerraformLog struct {

	// deployment id
	// Min Length: 1
	DeploymentID string `json:"deployment_id,omitempty"`

	// deployment name
	// Min Length: 1
	DeploymentName string `json:"deployment_name,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// initiator
	Initiator string `json:"initiator,omitempty"`

	// log
	// Read Only: true
	Log string `json:"log,omitempty"`

	// operation
	Operation string `json:"operation,omitempty"`

	// time
	// Min Length: 1
	Time string `json:"time,omitempty"`
}

// Validate validates this terraform log
func (m *TerraformLog) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeploymentID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeploymentName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TerraformLog) validateDeploymentID(formats strfmt.Registry) error {
	if swag.IsZero(m.DeploymentID) { // not required
		return nil
	}

	if err := validate.MinLength("deployment_id", "body", m.DeploymentID, 1); err != nil {
		return err
	}

	return nil
}

func (m *TerraformLog) validateDeploymentName(formats strfmt.Registry) error {
	if swag.IsZero(m.DeploymentName) { // not required
		return nil
	}

	if err := validate.MinLength("deployment_name", "body", m.DeploymentName, 1); err != nil {
		return err
	}

	return nil
}

func (m *TerraformLog) validateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.Time) { // not required
		return nil
	}

	if err := validate.MinLength("time", "body", m.Time, 1); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this terraform log based on the context it is used
func (m *TerraformLog) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLog(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TerraformLog) contextValidateLog(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "log", "body", string(m.Log)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TerraformLog) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TerraformLog) UnmarshalBinary(b []byte) error {
	var res TerraformLog
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
