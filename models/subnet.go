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

// Subnet subnet
//
// swagger:model subnet
type Subnet struct {

	// arn
	// Min Length: 1
	Arn string `json:"arn,omitempty"`

	// route table id
	RouteTableID string `json:"route_table_id,omitempty"`

	// subnet id
	// Required: true
	// Min Length: 1
	SubnetID *string `json:"subnet_id"`

	// tags
	// Required: true
	Tags []*Tag `json:"tags"`

	// zone
	Zone string `json:"zone,omitempty"`
}

// Validate validates this subnet
func (m *Subnet) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubnetID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Subnet) validateArn(formats strfmt.Registry) error {
	if swag.IsZero(m.Arn) { // not required
		return nil
	}

	if err := validate.MinLength("arn", "body", m.Arn, 1); err != nil {
		return err
	}

	return nil
}

func (m *Subnet) validateSubnetID(formats strfmt.Registry) error {

	if err := validate.Required("subnet_id", "body", m.SubnetID); err != nil {
		return err
	}

	if err := validate.MinLength("subnet_id", "body", *m.SubnetID, 1); err != nil {
		return err
	}

	return nil
}

func (m *Subnet) validateTags(formats strfmt.Registry) error {

	if err := validate.Required("tags", "body", m.Tags); err != nil {
		return err
	}

	for i := 0; i < len(m.Tags); i++ {
		if swag.IsZero(m.Tags[i]) { // not required
			continue
		}

		if m.Tags[i] != nil {
			if err := m.Tags[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this subnet based on the context it is used
func (m *Subnet) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Subnet) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Tags); i++ {

		if m.Tags[i] != nil {
			if err := m.Tags[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Subnet) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Subnet) UnmarshalBinary(b []byte) error {
	var res Subnet
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
