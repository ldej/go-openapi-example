// This file is auto-generated, DO NOT EDIT.
//
// Source:
//
//	Title: Things API
//	Version: 1.0
package api

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// GetThingQueryParameters is an object.
type GetThingQueryParameters struct {
	// Uuid: The uuid of the thing to get
	Uuid string `json:"uuid" mapstructure:"uuid"`
}

// Validate implements basic validation for this model
func (m GetThingQueryParameters) Validate() error {
	return validation.Errors{
		"uuid": validation.Validate(
			m.Uuid, validation.Required, is.UUID,
		),
	}.Filter()
}

// GetUuid returns the Uuid property
func (m GetThingQueryParameters) GetUuid() string {
	return m.Uuid
}

// SetUuid sets the Uuid property
func (m *GetThingQueryParameters) SetUuid(val string) {
	m.Uuid = val
}
