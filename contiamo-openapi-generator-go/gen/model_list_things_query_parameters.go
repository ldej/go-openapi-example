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

// ListThingsQueryParameters is an object.
type ListThingsQueryParameters struct {
	// Page: Page number
	Page int32 `json:"page,omitempty" mapstructure:"page,omitempty"`
	// Keyword: Filter things by keyword
	Keyword string `json:"keyword,omitempty" mapstructure:"keyword,omitempty"`
}

// Validate implements basic validation for this model
func (m ListThingsQueryParameters) Validate() error {
	return validation.Errors{}.Filter()
}

// GetPage returns the Page property
func (m ListThingsQueryParameters) GetPage() int32 {
	return m.Page
}

// SetPage sets the Page property
func (m *ListThingsQueryParameters) SetPage(val int32) {
	m.Page = val
}

// GetKeyword returns the Keyword property
func (m ListThingsQueryParameters) GetKeyword() string {
	return m.Keyword
}

// SetKeyword sets the Keyword property
func (m *ListThingsQueryParameters) SetKeyword(val string) {
	m.Keyword = val
}
