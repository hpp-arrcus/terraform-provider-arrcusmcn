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

// UpgradeCheck upgrade check
//
// swagger:model upgrade_check
type UpgradeCheck struct {

	// available
	// Required: true
	Available *bool `json:"available"`

	// current image
	CurrentImage string `json:"currentImage,omitempty"`

	// current version
	CurrentVersion string `json:"currentVersion,omitempty"`

	// latest image
	LatestImage string `json:"latestImage,omitempty"`

	// latest version
	LatestVersion string `json:"latestVersion,omitempty"`
}

// Validate validates this upgrade check
func (m *UpgradeCheck) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAvailable(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpgradeCheck) validateAvailable(formats strfmt.Registry) error {

	if err := validate.Required("available", "body", m.Available); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this upgrade check based on context it is used
func (m *UpgradeCheck) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UpgradeCheck) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpgradeCheck) UnmarshalBinary(b []byte) error {
	var res UpgradeCheck
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}