// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RunJobResponse run job response
//
// swagger:model RunJobResponse
type RunJobResponse struct {

	// run Id
	RunID string `json:"runId,omitempty"`
}

// Validate validates this run job response
func (m *RunJobResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this run job response based on context it is used
func (m *RunJobResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RunJobResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RunJobResponse) UnmarshalBinary(b []byte) error {
	var res RunJobResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}