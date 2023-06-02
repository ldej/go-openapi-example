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

// CreateThingRequest is an object.
type CreateThingRequest struct {
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// Rank:
	Rank int64 `json:"rank" mapstructure:"rank"`
	// Rating:
	Rating float32 `json:"rating,omitempty" mapstructure:"rating,omitempty"`
	// Score:
	Score float64 `json:"score,omitempty" mapstructure:"score,omitempty"`
	// Type:
	Type ThingType `json:"type" mapstructure:"type"`
}

// Validate implements basic validation for this model
func (m CreateThingRequest) Validate() error {
	return validation.Errors{
		"name": validation.Validate(
			m.Name, validation.NotNil, validation.Length(0, 10),
		),
		"rating": validation.Validate(
			m.Rating, validation.Min(float32(0)), validation.Max(float32(5)),
		),
		"type": validation.Validate(
			m.Type, validation.Required,
		),
	}.Filter()
}

// GetName returns the Name property
func (m CreateThingRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *CreateThingRequest) SetName(val string) {
	m.Name = val
}

// GetRank returns the Rank property
func (m CreateThingRequest) GetRank() int64 {
	return m.Rank
}

// SetRank sets the Rank property
func (m *CreateThingRequest) SetRank(val int64) {
	m.Rank = val
}

// GetRating returns the Rating property
func (m CreateThingRequest) GetRating() float32 {
	return m.Rating
}

// SetRating sets the Rating property
func (m *CreateThingRequest) SetRating(val float32) {
	m.Rating = val
}

// GetScore returns the Score property
func (m CreateThingRequest) GetScore() float64 {
	return m.Score
}

// SetScore sets the Score property
func (m *CreateThingRequest) SetScore(val float64) {
	m.Score = val
}

// GetType returns the Type property
func (m CreateThingRequest) GetType() ThingType {
	return m.Type
}

// SetType sets the Type property
func (m *CreateThingRequest) SetType(val ThingType) {
	m.Type = val
}
