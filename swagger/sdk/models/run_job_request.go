// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RunJobRequest run job request
//
// swagger:model RunJobRequest
type RunJobRequest struct {

	// group ID
	GroupID uint64 `json:"groupId,omitempty"`

	// payload
	Payload map[string]string `json:"payload,omitempty"`

	// target ID
	TargetID uint64 `json:"targetId,omitempty"`
}

// Validate validates this run job request
func (m *RunJobRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this run job request based on context it is used
func (m *RunJobRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RunJobRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RunJobRequest) UnmarshalBinary(b []byte) error {
	var res RunJobRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
