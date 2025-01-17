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

// OnpremCredentials onprem credentials
//
// swagger:model onprem_credentials
type OnpremCredentials struct {

	// coordinates
	Coordinates *Coordinates `json:"coordinates,omitempty"`

	// data if name
	DataIfName []string `json:"data_if_name"`

	// domain
	// Min Length: 1
	Domain string `json:"domain,omitempty"`

	// id
	// Min Length: 1
	ID string `json:"id,omitempty"`

	// server ip
	// Min Length: 1
	ServerIP string `json:"server_ip,omitempty"`

	// site address
	SiteAddress string `json:"site_address,omitempty"`

	// ssh key
	// Min Length: 1
	SSHKey string `json:"ssh_key,omitempty"`

	// user name
	// Min Length: 1
	UserName string `json:"user_name,omitempty"`
}

// Validate validates this onprem credentials
func (m *OnpremCredentials) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCoordinates(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDomain(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServerIP(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSSHKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OnpremCredentials) validateCoordinates(formats strfmt.Registry) error {
	if swag.IsZero(m.Coordinates) { // not required
		return nil
	}

	if m.Coordinates != nil {
		if err := m.Coordinates.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("coordinates")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("coordinates")
			}
			return err
		}
	}

	return nil
}

func (m *OnpremCredentials) validateDomain(formats strfmt.Registry) error {
	if swag.IsZero(m.Domain) { // not required
		return nil
	}

	if err := validate.MinLength("domain", "body", m.Domain, 1); err != nil {
		return err
	}

	return nil
}

func (m *OnpremCredentials) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.MinLength("id", "body", m.ID, 1); err != nil {
		return err
	}

	return nil
}

func (m *OnpremCredentials) validateServerIP(formats strfmt.Registry) error {
	if swag.IsZero(m.ServerIP) { // not required
		return nil
	}

	if err := validate.MinLength("server_ip", "body", m.ServerIP, 1); err != nil {
		return err
	}

	return nil
}

func (m *OnpremCredentials) validateSSHKey(formats strfmt.Registry) error {
	if swag.IsZero(m.SSHKey) { // not required
		return nil
	}

	if err := validate.MinLength("ssh_key", "body", m.SSHKey, 1); err != nil {
		return err
	}

	return nil
}

func (m *OnpremCredentials) validateUserName(formats strfmt.Registry) error {
	if swag.IsZero(m.UserName) { // not required
		return nil
	}

	if err := validate.MinLength("user_name", "body", m.UserName, 1); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this onprem credentials based on the context it is used
func (m *OnpremCredentials) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCoordinates(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OnpremCredentials) contextValidateCoordinates(ctx context.Context, formats strfmt.Registry) error {

	if m.Coordinates != nil {
		if err := m.Coordinates.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("coordinates")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("coordinates")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OnpremCredentials) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OnpremCredentials) UnmarshalBinary(b []byte) error {
	var res OnpremCredentials
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
