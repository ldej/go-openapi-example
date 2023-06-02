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

// ThingType is an enum.
type ThingType string

// Validate implements basic validation for this model
func (m ThingType) Validate() error {
	return InKnownThingType.Validate(m)
}

var (
	ThingTypeAny   ThingType = "any"
	ThingTypeFalse ThingType = "false"
	ThingTypeSome  ThingType = "some"

	// KnownThingType is the list of valid ThingType
	KnownThingType = []ThingType{
		ThingTypeAny,
		ThingTypeFalse,
		ThingTypeSome,
	}
	// KnownThingTypeString is the list of valid ThingType as string
	KnownThingTypeString = []string{
		string(ThingTypeAny),
		string(ThingTypeFalse),
		string(ThingTypeSome),
	}

	// InKnownThingType is an ozzo-validator for ThingType
	InKnownThingType = validation.In(
		ThingTypeAny,
		ThingTypeFalse,
		ThingTypeSome,
	)
)
