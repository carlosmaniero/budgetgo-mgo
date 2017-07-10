package handlers

import (
	"fmt"
	"net/http"

	"github.com/carlosmaniero/budgetgo/interfaces/serializers"
	"github.com/carlosmaniero/budgetgo/usecases"
	"github.com/julienschmidt/httprouter"
)

func (handlers *Handlers) FundingCreate(response http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	defer handlers.catchPanics(response)

	iteractor := usecases.FundingInteractor{Repository: handlers.Application.FundingRepository}
	data, err := serializers.UnserializeFundingData(request.Body)

	if err != nil {
		handlers.UnserializerErrorHandler(err, response)
		return
	}

	err, funding := iteractor.Register(data.Name, data.Amount, data.ClosingDay, data.Limit)

	if err != nil {
		handlers.FundingCreateErrorHandler(err, response)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	fmt.Fprint(response, string(serializers.SerializeFunding(funding)))
}

func (handlers *Handlers) FundingCreateErrorHandler(err error, response http.ResponseWriter) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusBadRequest)

	switch err := err.(type) {
	case *usecases.FundingValidationErrors:

		errorResponse := serializers.FundingValidationErrorData{
			Type:    "validation_error",
			Message: err.Error(),
			Errors:  handlers.convertFieldValidationErrors(err.Errors),
		}

		data := serializers.SerializeFundingValidationError(&errorResponse)
		fmt.Fprint(response, string(data))
	default:
		errorResponse := serializers.ErrorResponseData{
			Type:    "domain_error",
			Message: err.Error(),
		}

		data := serializers.SerializeErrorResponse(&errorResponse)
		fmt.Fprint(response, string(data))
	}
}
