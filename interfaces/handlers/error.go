package handlers

import (
	"net/http"
	"encoding/json"
	"fmt"
	"github.com/carlosmaniero/budgetgo/interfaces/serializers"
	"time"
)


func (handlers *Handlers) UnserializerErrorHandler (err error, response http.ResponseWriter) {
	errorResponse := handlers.createErrorResponse(err)
	data := serializers.SerializeErrorResponse(&errorResponse)

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusBadRequest)
	fmt.Fprint(response, string(data))
}

func (handlers *Handlers) createErrorResponse(err error) serializers.ErrorResponseData{
	switch err := err.(type) {
	case *json.UnmarshalTypeError:
		return serializers.ErrorResponseData{
			Type: "parser",
			Message: "cannot add " + err.Value + " value into field " +  err.Field + " of type " + err.Type.String(),
		}
	case *time.ParseError:
		return serializers.ErrorResponseData{
			Type: "parser",
			Message: "cannot parse the sent date. Check the date format. Date Formate: " + time.RFC3339 + " (RFC3339)",
		}
	default:
		return serializers.ErrorResponseData{
			Type: "parser",
			Message: "An error was occurred when parse your request body",
		}
	}
}