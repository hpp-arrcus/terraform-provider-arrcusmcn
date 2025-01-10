// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	hourly_octet_time "time"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ArcCostHourlyOctet arc cost hourly octet
//
// swagger:model arc_cost_hourly_octet
type ArcCostHourlyOctet struct {

	// currency code
	CurrencyCode string `json:"currencyCode,omitempty"`

	// end time
	EndTime hourly_octet_time.Time `json:"endTime,omitempty"`

	// hourly in octets
	HourlyInOctets float64 `json:"hourlyInOctets,omitempty"`

	// hourly out octets
	HourlyOutOctets float64 `json:"hourlyOutOctets,omitempty"`

	// in octet cost
	InOctetCost float64 `json:"inOctetCost,omitempty"`

	// out octet cost
	OutOctetCost float64 `json:"outOctetCost,omitempty"`

	// start time
	StartTime hourly_octet_time.Time `json:"startTime,omitempty"`
}

// Validate validates this arc cost hourly octet
func (m *ArcCostHourlyOctet) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this arc cost hourly octet based on context it is used
func (m *ArcCostHourlyOctet) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ArcCostHourlyOctet) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ArcCostHourlyOctet) UnmarshalBinary(b []byte) error {
	var res ArcCostHourlyOctet
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}