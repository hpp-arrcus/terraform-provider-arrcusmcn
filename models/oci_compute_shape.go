// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// OciComputeShape oci compute shape
//
// swagger:model oci_compute_shape
type OciComputeShape struct {

	// default if per ocpu
	DefaultIfPerOcpu *int64 `json:"default_if_per_ocpu,omitempty"`

	// is flexible
	IsFlexible bool `json:"is_flexible,omitempty"`

	// max memory in gbs
	MaxMemoryInGbs *float64 `json:"max_memory_in_gbs,omitempty"`

	// max ocpus
	MaxOcpus *float64 `json:"max_ocpus,omitempty"`

	// max vnic attachments
	MaxVnicAttachments *int64 `json:"max_vnic_attachments,omitempty"`

	// memory in gbs
	MemoryInGbs float64 `json:"memory_in_gbs,omitempty"`

	// min memory in gbs
	MinMemoryInGbs *float64 `json:"min_memory_in_gbs,omitempty"`

	// min ocpus
	MinOcpus *float64 `json:"min_ocpus,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// ocpus
	Ocpus float64 `json:"ocpus,omitempty"`
}

// Validate validates this oci compute shape
func (m *OciComputeShape) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this oci compute shape based on context it is used
func (m *OciComputeShape) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OciComputeShape) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OciComputeShape) UnmarshalBinary(b []byte) error {
	var res OciComputeShape
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
