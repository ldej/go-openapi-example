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

// UpdateThingRequest is an object.
type UpdateThingRequest struct {
	// Score:
	Score float64 `json:"score" mapstructure:"score"`
}

// Validate implements basic validation for this model
func (m UpdateThingRequest) Validate() error {
	return validation.Errors{}.Filter()
}

// GetScore returns the Score property
func (m UpdateThingRequest) GetScore() float64 {
	return m.Score
}

// SetScore sets the Score property
func (m *UpdateThingRequest) SetScore(val float64) {
	m.Score = val
}
