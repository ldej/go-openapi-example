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



type ListThingsResponse struct {

	Things []ThingResponse `json:"things"`

	Total int32 `json:"total"`
}

// UnmarshalJSON sets *m to a copy of data while respecting defaults if specified.
func (m *ListThingsResponse) UnmarshalJSON(data []byte) error {

	type Alias ListThingsResponse // To avoid infinite recursion
    return json.Unmarshal(data, (*Alias)(m))
}

// AssertListThingsResponseRequired checks if the required fields are not zero-ed
func AssertListThingsResponseRequired(obj ListThingsResponse) error {
	elements := map[string]interface{}{
		"things": obj.Things,
		"total": obj.Total,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	for _, el := range obj.Things {
		if err := AssertThingResponseRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertListThingsResponseConstraints checks if the values respects the defined constraints
func AssertListThingsResponseConstraints(obj ListThingsResponse) error {
	return nil
}