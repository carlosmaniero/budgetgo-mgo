package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/carlosmaniero/budgetgo/domain"
	"github.com/carlosmaniero/budgetgo/interfaces/repositories/memoryrepository"
	"github.com/carlosmaniero/budgetgo/interfaces/serializers"
)

func (handlers *Handlers) unserializerErrorHandler(err error, response http.ResponseWriter) {
	errorResponse := handlers.getErrorResponse(err)
	data := serializers.SerializeErrorResponse(&errorResponse)

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusBadRequest)
	fmt.Fprint(response, string(data))
}

func (handlers *Handlers) catchPanics(response http.ResponseWriter) {
	if err := recover(); err != nil {
		errorResponse := handlers.getErrorResponse(err)
		data := serializers.SerializeErrorResponse(&errorResponse)

		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(response, string(data))
	}
}

func (handlers *Handlers) getErrorResponse(err interface{}) serializers.ErrorResponseData {
	switch err := err.(type) {
	case *json.UnmarshalTypeError:
		return serializers.ErrorResponseData{
			Type:    "parser",
			Message: "cannot add " + err.Value + " value into field " + err.Field + " of type " + err.Type.String(),
		}
	case *time.ParseError:
		return serializers.ErrorResponseData{
			Type:    "parser",
			Message: "cannot parse the sent date. Check the date format. Date Formate: " + time.RFC3339 + " (RFC3339)",
		}
	case *memoryrepository.MemoryMaxTransactionsError:
		return serializers.ErrorResponseData{
			Type:    "server_error",
			Message: err.Error(),
		}
	case error:
		return serializers.ErrorResponseData{
			Type:    "server_error",
			Message: err.Error(),
		}
	default:
		return serializers.ErrorResponseData{
			Type:    "server_error",
			Message: "An error was occurred check your request body",
		}
	}
}

func (handlers *Handlers) convertFieldValidationErrors(errors []error) []*serializers.FieldErrorData {
	fieldErrors := make([]*serializers.FieldErrorData, len(errors))
	for index, value := range errors {
		err := value.(*domain.FieldValidationError)
		fieldErrors[index] = &serializers.FieldErrorData{
			Field:   strings.ToLower(err.Field),
			Message: err.Message,
		}
	}
	return fieldErrors
}
