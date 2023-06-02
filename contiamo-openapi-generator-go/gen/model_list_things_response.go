// This file is auto-generated, DO NOT EDIT.
//
// Source:
//
//	Title: Things API
//	Version: 1.0
package api

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// ListThingsResponse is an object.
type ListThingsResponse struct {
	// Things:
	Things []ThingResponse `json:"things" mapstructure:"things"`
	// Total:
	Total int32 `json:"total" mapstructure:"total"`
}

// Validate implements basic validation for this model
func (m ListThingsResponse) Validate() error {
	return validation.Errors{
		"things": validation.Validate(
			m.Things, validation.NotNil,
		),
	}.Filter()
}

// GetThings returns the Things property
func (m ListThingsResponse) GetThings() []ThingResponse {
	return m.Things
}

// SetThings sets the Things property
func (m *ListThingsResponse) SetThings(val []ThingResponse) {
	m.Things = val
}

// GetTotal returns the Total property
func (m ListThingsResponse) GetTotal() int32 {
	return m.Total
}

// SetTotal sets the Total property
func (m *ListThingsResponse) SetTotal(val int32) {
	m.Total = val
}
