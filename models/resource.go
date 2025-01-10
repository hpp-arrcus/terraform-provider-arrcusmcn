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

// Resource resource
//
// swagger:model resource
type Resource string

func NewResource(value Resource) *Resource {
	return &value
}

// Pointer returns a pointer to a freshly-allocated Resource.
func (m Resource) Pointer() *Resource {
	return &m
}

const (

	// ResourceTenants captures enum value "tenants"
	ResourceTenants Resource = "tenants"

	// ResourceUsers captures enum value "users"
	ResourceUsers Resource = "users"

	// ResourcePasswords captures enum value "passwords"
	ResourcePasswords Resource = "passwords"

	// ResourceCloudCredentials captures enum value "cloud_credentials"
	ResourceCloudCredentials Resource = "cloud_credentials"

	// ResourceDeployments captures enum value "deployments"
	ResourceDeployments Resource = "deployments"

	// ResourceConnections captures enum value "connections"
	ResourceConnections Resource = "connections"

	// ResourceArcedgeCredentials captures enum value "arcedge_credentials"
	ResourceArcedgeCredentials Resource = "arcedge_credentials"

	// ResourceCertificates captures enum value "certificates"
	ResourceCertificates Resource = "certificates"

	// ResourceDNS captures enum value "dns"
	ResourceDNS Resource = "dns"

	// ResourceUpgrade captures enum value "upgrade"
	ResourceUpgrade Resource = "upgrade"

	// ResourceCloudResources captures enum value "cloud_resources"
	ResourceCloudResources Resource = "cloud_resources"

	// ResourceFqdnlist captures enum value "fqdnlist"
	ResourceFqdnlist Resource = "fqdnlist"

	// ResourceActivefqdnlists captures enum value "activefqdnlists"
	ResourceActivefqdnlists Resource = "activefqdnlists"

	// ResourceTrusteddns captures enum value "trusteddns"
	ResourceTrusteddns Resource = "trusteddns"

	// ResourceDiscoveredfqdns captures enum value "discoveredfqdns"
	ResourceDiscoveredfqdns Resource = "discoveredfqdns"

	// ResourceFqdndiscovery captures enum value "fqdndiscovery"
	ResourceFqdndiscovery Resource = "fqdndiscovery"

	// ResourceZerotrust captures enum value "zerotrust"
	ResourceZerotrust Resource = "zerotrust"

	// ResourceGlobaldashboard captures enum value "globaldashboard"
	ResourceGlobaldashboard Resource = "globaldashboard"

	// ResourceTenantdashboard captures enum value "tenantdashboard"
	ResourceTenantdashboard Resource = "tenantdashboard"

	// ResourceAuditlogs captures enum value "auditlogs"
	ResourceAuditlogs Resource = "auditlogs"

	// ResourceArccost captures enum value "arccost"
	ResourceArccost Resource = "arccost"

	// ResourceArcorchapikey captures enum value "arcorchapikey"
	ResourceArcorchapikey Resource = "arcorchapikey"

	// ResourceSamlconfg captures enum value "samlconfg"
	ResourceSamlconfg Resource = "samlconfg"
)

// for schema
var resourceEnum []interface{}

func init() {
	var res []Resource
	if err := json.Unmarshal([]byte(`["tenants","users","passwords","cloud_credentials","deployments","connections","arcedge_credentials","certificates","dns","upgrade","cloud_resources","fqdnlist","activefqdnlists","trusteddns","discoveredfqdns","fqdndiscovery","zerotrust","globaldashboard","tenantdashboard","auditlogs","arccost","arcorchapikey","samlconfg"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		resourceEnum = append(resourceEnum, v)
	}
}

func (m Resource) validateResourceEnum(path, location string, value Resource) error {
	if err := validate.EnumCase(path, location, value, resourceEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this resource
func (m Resource) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateResourceEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this resource based on context it is used
func (m Resource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}