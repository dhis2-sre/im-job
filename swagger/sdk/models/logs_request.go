// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// LogsRequest logs request
//
// swagger:model LogsRequest
type LogsRequest struct {

	// group ID
	GroupID uint64 `json:"groupId,omitempty"`
}

// Validate validates this logs request
func (m *LogsRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this logs request based on context it is used
func (m *LogsRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LogsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LogsRequest) UnmarshalBinary(b []byte) error {
	var res LogsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}