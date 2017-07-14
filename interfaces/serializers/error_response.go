package serializers

import (
	"encoding/json"
)

// ErrorResponseData is the representation of an system error response
//
// The type can be by exemple; "not-found"
// And the message: "The entity was not found"
type ErrorResponseData struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

// FieldErrorData is the representations of an field error validation
//
// This contains the name of the field and the error message
type FieldErrorData struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// SerializeErrorResponse serializes the ErrorResponseData and returns a
// string json representation
func SerializeErrorResponse(errorResponse *ErrorResponseData) []byte {
	b, _ := json.Marshal(errorResponse)
	return b
}
