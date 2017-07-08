package handlers

import (
	"net/http"
	"encoding/json"
	"fmt"
	"github.com/carlosmaniero/budgetgo/interfaces/serializers"
	"reflect"
	"time"
)


func (handlers *Handlers) UnserializerErrorHandler (err error, response http.ResponseWriter) {
	response.Header().Set("Content-Type", "application/json")

	switch err := err.(type) {
	case *json.UnmarshalTypeError:
		errorResponse := serializers.ErrorResponseData{
			Type: "parser",
			Message: "cannot add " + err.Value + " value into field " +  err.Field + " of type " + err.Type.String(),
		}

		data := serializers.SerializeErrorResponse(&errorResponse)
		fmt.Fprint(response, string(data))
	case *time.ParseError:
		errorResponse := serializers.ErrorResponseData{
			Type: "parser",
			Message: "cannot parse the sent date. Check the date format. Date Formate: " + time.RFC3339 + " (RFC3339)",
		}

		data := serializers.SerializeErrorResponse(&errorResponse)
		fmt.Fprint(response, string(data))
	default:
		fmt.Println(reflect.TypeOf(err))
		errorResponse := serializers.ErrorResponseData{
			Type: "parser",
			Message: "An error was occurred when parse your request body",
		}

		data := serializers.SerializeErrorResponse(&errorResponse)
		fmt.Fprint(response, string(data))
	}
}