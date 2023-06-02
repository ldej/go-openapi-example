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

// Error is an object.
type Error struct {
	// Code:
	Code string `json:"code" mapstructure:"code"`
	// Message:
	Message string `json:"message" mapstructure:"message"`
}

// Validate implements basic validation for this model
func (m Error) Validate() error {
	return validation.Errors{}.Filter()
}

// GetCode returns the Code property
func (m Error) GetCode() string {
	return m.Code
}

// SetCode sets the Code property
func (m *Error) SetCode(val string) {
	m.Code = val
}

// GetMessage returns the Message property
func (m Error) GetMessage() string {
	return m.Message
}

// SetMessage sets the Message property
func (m *Error) SetMessage(val string) {
	m.Message = val
}
