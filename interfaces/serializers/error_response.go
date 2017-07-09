package serializers

import (
	"encoding/json"
)

type ErrorResponseData struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

func SerializeErrorResponse(errorResponse *ErrorResponseData) []byte {
	b, _ := json.Marshal(errorResponse)
	return b
}
