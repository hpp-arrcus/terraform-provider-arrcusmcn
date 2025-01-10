// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CloudResourcesOutput cloud resources output
//
// swagger:model cloud_resources_output
type CloudResourcesOutput struct {

	// deployed config
	DeployedConfig *CloudResources `json:"deployed_config,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// last error
	LastError string `json:"last_error,omitempty"`

	// last updated
	LastUpdated string `json:"last_updated,omitempty"`

	// operation
	Operation TerraformOperation `json:"operation,omitempty"`

	// provider
	Provider Providers `json:"provider,omitempty"`

	// resource state
	ResourceState struct {
		AwsCloudResourcesStateList

		GcpCloudResourcesStateList

		AzureCloudResourcesStateList
	} `json:"resource_state,omitempty"`

	// status
	Status Status `json:"status,omitempty"`
}

// Validate validates this cloud resources output
func (m *CloudResourcesOutput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeployedConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProvider(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourceState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudResourcesOutput) validateDeployedConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.DeployedConfig) { // not required
		return nil
	}

	if m.DeployedConfig != nil {
		if err := m.DeployedConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deployed_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("deployed_config")
			}
			return err
		}
	}

	return nil
}

func (m *CloudResourcesOutput) validateOperation(formats strfmt.Registry) error {
	if swag.IsZero(m.Operation) { // not required
		return nil
	}

	if err := m.Operation.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("operation")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("operation")
		}
		return err
	}

	return nil
}

func (m *CloudResourcesOutput) validateProvider(formats strfmt.Registry) error {
	if swag.IsZero(m.Provider) { // not required
		return nil
	}

	if err := m.Provider.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("provider")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("provider")
		}
		return err
	}

	return nil
}

func (m *CloudResourcesOutput) validateResourceState(formats strfmt.Registry) error {
	if swag.IsZero(m.ResourceState) { // not required
		return nil
	}

	return nil
}

func (m *CloudResourcesOutput) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := m.Status.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("status")
		}
		return err
	}

	return nil
}

// ContextValidate validate this cloud resources output based on the context it is used
func (m *CloudResourcesOutput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDeployedConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOperation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProvider(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResourceState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudResourcesOutput) contextValidateDeployedConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.DeployedConfig != nil {
		if err := m.DeployedConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deployed_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("deployed_config")
			}
			return err
		}
	}

	return nil
}

func (m *CloudResourcesOutput) contextValidateOperation(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Operation.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("operation")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("operation")
		}
		return err
	}

	return nil
}

func (m *CloudResourcesOutput) contextValidateProvider(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Provider.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("provider")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("provider")
		}
		return err
	}

	return nil
}

func (m *CloudResourcesOutput) contextValidateResourceState(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *CloudResourcesOutput) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Status.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("status")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CloudResourcesOutput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloudResourcesOutput) UnmarshalBinary(b []byte) error {
	var res CloudResourcesOutput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
