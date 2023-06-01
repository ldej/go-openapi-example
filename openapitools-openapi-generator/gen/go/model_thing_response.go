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
	"time"
	"encoding/json"
)



type ThingResponse struct {

	Uuid string `json:"uuid"`

	Type ThingType `json:"type"`

	Name string `json:"name"`

	Rank int64 `json:"rank"`

	Score float64 `json:"score"`

	Rating float32 `json:"rating"`

	Created time.Time `json:"created"`
}

// UnmarshalJSON sets *m to a copy of data while respecting defaults if specified.
func (m *ThingResponse) UnmarshalJSON(data []byte) error {

	type Alias ThingResponse // To avoid infinite recursion
    return json.Unmarshal(data, (*Alias)(m))
}

// AssertThingResponseRequired checks if the required fields are not zero-ed
func AssertThingResponseRequired(obj ThingResponse) error {
	elements := map[string]interface{}{
		"uuid": obj.Uuid,
		"type": obj.Type,
		"name": obj.Name,
		"rank": obj.Rank,
		"score": obj.Score,
		"rating": obj.Rating,
		"created": obj.Created,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertThingResponseConstraints checks if the values respects the defined constraints
func AssertThingResponseConstraints(obj ThingResponse) error {
	return nil
}
