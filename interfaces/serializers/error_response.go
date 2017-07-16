package serializers

import (
	"encoding/json"
)

// FieldErrorData is the representations of an field error validation
//
// This contains the name of the field and the error message
type FieldErrorData struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// ValidationErrorData is the serializable representaion of a Funding
// validation error
type ValidationErrorData struct {
	Type    string            `json:"type"`
	Message string            `json:"message"`
	Errors  []*FieldErrorData `json:"errors"`
}

// Serialize returns a json string representation of an
// FundingValidationErrorData
func (data *ValidationErrorData) Serialize() []byte {
	b, _ := json.Marshal(data)
	return b
}

// ErrorResponseSerializer is the representation of an system error response
//
// The type can be by exemple; "not-found"
// And the message: "The entity was not found"
type ErrorResponseSerializer struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

// Serialize serializes the ErrorResponseData and returns a
// string json representation
func (errorResponse *ErrorResponseSerializer) Serialize() []byte {
	b, _ := json.Marshal(errorResponse)
	return b
}
