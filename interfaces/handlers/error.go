package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/carlosmaniero/budgetgo/domain"
	"github.com/carlosmaniero/budgetgo/interfaces/serializers"
	"github.com/carlosmaniero/budgetgo/usecases"
)

func (handlers *Handlers) unserializerErrorHandler(err error, response http.ResponseWriter) {
	serializer := handlers.getErrorResponse(err)
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusBadRequest)
	response.Write(serializer.Serialize())
}

func (handlers *Handlers) catchPanics(response http.ResponseWriter) {
	if err := recover(); err != nil {
		serializer := handlers.getErrorResponse(err)
		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusInternalServerError)
		response.Write(serializer.Serialize())
	}
}

func (handlers *Handlers) getErrorResponse(err interface{}) serializers.ErrorResponseSerializer {
	switch err := err.(type) {
	case *json.UnmarshalTypeError:
		return serializers.ErrorResponseSerializer{
			Type:    "parser",
			Message: "cannot add " + err.Value + " value into field " + err.Field + " of type " + err.Type.String(),
		}
	case *time.ParseError:
		return serializers.ErrorResponseSerializer{
			Type:    "parser",
			Message: "cannot parse the sent date. Check the date format. Date Formate: " + time.RFC3339 + " (RFC3339)",
		}
	case error:
		return serializers.ErrorResponseSerializer{
			Type:    "server_error",
			Message: err.Error(),
		}
	default:
		return serializers.ErrorResponseSerializer{
			Type:    "server_error",
			Message: "An error was occurred check your request body",
		}
	}
}

// This is the error handler of the transaction creation
func (handlers *Handlers) usecaseErrorHandler(err error, response http.ResponseWriter) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusBadRequest)

	switch err := err.(type) {
	case *usecases.ValidationErrors:

		serializer := serializers.ValidationErrorData{
			Type:    "validation_error",
			Message: err.Error(),
			Errors:  handlers.convertFieldValidationErrors(err.Errors),
		}

		response.Write(serializer.Serialize())
	default:
		serializer := serializers.ErrorResponseSerializer{
			Type:    "domain_error",
			Message: err.Error(),
		}

		response.Write(serializer.Serialize())
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
