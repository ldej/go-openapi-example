/*
 * Things API
 *
 * The Things API creates, reads, updates, lists and deletes things!
 *
 * API version: 1.0
 * Contact: info@ldej.nl
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi


import (
	"encoding/json"
)



type UpdateThingRequest struct {

	Score float64 `json:"score"`
}

// UnmarshalJSON sets *m to a copy of data while respecting defaults if specified.
func (m *UpdateThingRequest) UnmarshalJSON(data []byte) error {

	type Alias UpdateThingRequest // To avoid infinite recursion
    return json.Unmarshal(data, (*Alias)(m))
}

// AssertUpdateThingRequestRequired checks if the required fields are not zero-ed
func AssertUpdateThingRequestRequired(obj UpdateThingRequest) error {
	elements := map[string]interface{}{
		"score": obj.Score,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertUpdateThingRequestConstraints checks if the values respects the defined constraints
func AssertUpdateThingRequestConstraints(obj UpdateThingRequest) error {
	return nil
}
