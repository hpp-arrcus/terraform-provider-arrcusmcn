// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SamlConfig saml config
//
// swagger:model saml_config
type SamlConfig struct {

	// cert file
	CertFile string `json:"cert_file,omitempty"`

	// enable
	Enable bool `json:"enable,omitempty"`

	// entity id
	EntityID string `json:"entity_id,omitempty"`

	// key file
	KeyFile string `json:"key_file,omitempty"`

	// metadata url
	MetadataURL string `json:"metadata_url,omitempty"`

	// saml metadata
	SamlMetadata string `json:"saml_metadata,omitempty"`

	// update time
	UpdateTime int64 `json:"update_time,omitempty"`
}

// Validate validates this saml config
func (m *SamlConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this saml config based on context it is used
func (m *SamlConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SamlConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SamlConfig) UnmarshalBinary(b []byte) error {
	var res SamlConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
