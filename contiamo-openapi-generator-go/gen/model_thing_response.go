// This file is auto-generated, DO NOT EDIT.
//
// Source:
//
//	Title: Things API
//	Version: 1.0
package api

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"time"

	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// ThingResponse is an object.
type ThingResponse struct {
	// Created:
	Created time.Time `json:"created" mapstructure:"created"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// Rank:
	Rank int64 `json:"rank" mapstructure:"rank"`
	// Rating:
	Rating float32 `json:"rating" mapstructure:"rating"`
	// Score:
	Score float64 `json:"score" mapstructure:"score"`
	// Type:
	Type ThingType `json:"type" mapstructure:"type"`
	// Uuid:
	Uuid string `json:"uuid" mapstructure:"uuid"`
}

// Validate implements basic validation for this model
func (m ThingResponse) Validate() error {
	return validation.Errors{
		"type": validation.Validate(
			m.Type, validation.Required,
		),
		"uuid": validation.Validate(
			m.Uuid, validation.Required, is.UUID,
		),
	}.Filter()
}

// GetCreated returns the Created property
func (m ThingResponse) GetCreated() time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *ThingResponse) SetCreated(val time.Time) {
	m.Created = val
}

// GetName returns the Name property
func (m ThingResponse) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *ThingResponse) SetName(val string) {
	m.Name = val
}

// GetRank returns the Rank property
func (m ThingResponse) GetRank() int64 {
	return m.Rank
}

// SetRank sets the Rank property
func (m *ThingResponse) SetRank(val int64) {
	m.Rank = val
}

// GetRating returns the Rating property
func (m ThingResponse) GetRating() float32 {
	return m.Rating
}

// SetRating sets the Rating property
func (m *ThingResponse) SetRating(val float32) {
	m.Rating = val
}

// GetScore returns the Score property
func (m ThingResponse) GetScore() float64 {
	return m.Score
}

// SetScore sets the Score property
func (m *ThingResponse) SetScore(val float64) {
	m.Score = val
}

// GetType returns the Type property
func (m ThingResponse) GetType() ThingType {
	return m.Type
}

// SetType sets the Type property
func (m *ThingResponse) SetType(val ThingType) {
	m.Type = val
}

// GetUuid returns the Uuid property
func (m ThingResponse) GetUuid() string {
	return m.Uuid
}

// SetUuid sets the Uuid property
func (m *ThingResponse) SetUuid(val string) {
	m.Uuid = val
}
