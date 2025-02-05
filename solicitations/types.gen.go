// Package solicitations provides primitives to interact the openapi HTTP API.
//
// Code generated by go-sdk-codegen DO NOT EDIT.
package solicitations

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

// CreateProductReviewAndSellerFeedbackSolicitationResponse defines model for CreateProductReviewAndSellerFeedbackSolicitationResponse.
type CreateProductReviewAndSellerFeedbackSolicitationResponse struct {

	// A list of error responses returned when a request is unsuccessful.
	Errors *ErrorList `json:"errors,omitempty"`
}

// Error defines model for Error.
type Error struct {

	// An error code that identifies the type of error that occurred.
	Code string `json:"code"`

	// Additional details that can help the caller understand or fix the issue.
	Details *string `json:"details,omitempty"`

	// A message that describes the error condition in a human-readable form.
	Message string `json:"message"`
}

// ErrorList defines model for ErrorList.
type ErrorList []Error

// GetSchemaResponse defines model for GetSchemaResponse.
type GetSchemaResponse struct {
	Links *struct {

		// A Link object.
		Self LinkObject `json:"self"`
	} `json:"_links,omitempty"`

	// A list of error responses returned when a request is unsuccessful.
	Errors *ErrorList `json:"errors,omitempty"`

	// A JSON schema document describing the expected payload of the action. This object can be validated against <a href=http://json-schema.org/draft-04/schema>http://json-schema.org/draft-04/schema</a>.
	Payload *Schema `json:"payload,omitempty"`
}

// GetSolicitationActionResponse defines model for GetSolicitationActionResponse.
type GetSolicitationActionResponse struct {
	Embedded *struct {
		Schema *GetSchemaResponse `json:"schema,omitempty"`
	} `json:"_embedded,omitempty"`
	Links *struct {

		// A Link object.
		Schema LinkObject `json:"schema"`

		// A Link object.
		Self LinkObject `json:"self"`
	} `json:"_links,omitempty"`

	// A list of error responses returned when a request is unsuccessful.
	Errors *ErrorList `json:"errors,omitempty"`

	// A simple object containing the name of the template.
	Payload *SolicitationsAction `json:"payload,omitempty"`
}

// GetSolicitationActionsForOrderResponse defines model for GetSolicitationActionsForOrderResponse.
type GetSolicitationActionsForOrderResponse struct {
	Embedded *struct {
		Actions []GetSolicitationActionResponse `json:"actions"`
	} `json:"_embedded,omitempty"`
	Links *struct {

		// Eligible actions for the specified amazonOrderId.
		Actions []LinkObject `json:"actions"`

		// A Link object.
		Self LinkObject `json:"self"`
	} `json:"_links,omitempty"`

	// A list of error responses returned when a request is unsuccessful.
	Errors *ErrorList `json:"errors,omitempty"`
}

// LinkObject defines model for LinkObject.
type LinkObject struct {

	// A URI for this object.
	Href string `json:"href"`

	// An identifier for this object.
	Name *string `json:"name,omitempty"`
}

// Schema defines model for Schema.
type Schema struct {
	AdditionalProperties map[string]interface{} `json:"-"`
}

// SolicitationsAction defines model for SolicitationsAction.
type SolicitationsAction struct {
	Name string `json:"name"`
}

// GetSolicitationActionsForOrderParams defines parameters for GetSolicitationActionsForOrder.
type GetSolicitationActionsForOrderParams struct {

	// A marketplace identifier. This specifies the marketplace in which the order was placed. Only one marketplace can be specified.
	MarketplaceIds []string `json:"marketplaceIds"`
}

// CreateProductReviewAndSellerFeedbackSolicitationParams defines parameters for CreateProductReviewAndSellerFeedbackSolicitation.
type CreateProductReviewAndSellerFeedbackSolicitationParams struct {

	// A marketplace identifier. This specifies the marketplace in which the order was placed. Only one marketplace can be specified.
	MarketplaceIds []string `json:"marketplaceIds"`
}

// Getter for additional properties for Schema. Returns the specified
// element and whether it was found
func (a Schema) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for Schema
func (a *Schema) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]interface{})
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for Schema to handle AdditionalProperties
func (a *Schema) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]interface{})
		for fieldName, fieldBuf := range object {
			var fieldVal interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return errors.Wrap(err, fmt.Sprintf("error unmarshaling field %s", fieldName))
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for Schema to handle AdditionalProperties
func (a Schema) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("error marshaling '%s'", fieldName))
		}
	}
	return json.Marshal(object)
}
