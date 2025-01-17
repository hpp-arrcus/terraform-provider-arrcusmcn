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

// Booleantype booleantype
//
// swagger:model booleantype
type Booleantype string

func NewBooleantype(value Booleantype) *Booleantype {
	return &value
}

// Pointer returns a pointer to a freshly-allocated Booleantype.
func (m Booleantype) Pointer() *Booleantype {
	return &m
}

const (

	// BooleantypeTrue captures enum value "true"
	BooleantypeTrue Booleantype = "true"

	// BooleantypeFalse captures enum value "false"
	BooleantypeFalse Booleantype = "false"
)

// for schema
var booleantypeEnum []interface{}

func init() {
	var res []Booleantype
	if err := json.Unmarshal([]byte(`["true","false"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		booleantypeEnum = append(booleantypeEnum, v)
	}
}

func (m Booleantype) validateBooleantypeEnum(path, location string, value Booleantype) error {
	if err := validate.EnumCase(path, location, value, booleantypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this booleantype
func (m Booleantype) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateBooleantypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this booleantype based on context it is used
func (m Booleantype) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
