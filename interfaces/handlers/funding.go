package handlers

import (
	"fmt"
	"net/http"

	"github.com/carlosmaniero/budgetgo/interfaces/serializers"
	"github.com/carlosmaniero/budgetgo/usecases"
	"github.com/julienschmidt/httprouter"
)

// FundingCreate is the handler of the funding creation entrypoint
func (handlers *Handlers) FundingCreate(response http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	defer handlers.catchPanics(response)

	iteractor := usecases.FundingInteractor{Repository: handlers.Application.FundingRepository}
	data, err := serializers.UnserializeFundingData(request.Body)

	if err != nil {
		handlers.unserializerErrorHandler(err, response)
		return
	}

	funding, err := iteractor.Register(data.Name, data.Amount, data.ClosingDay, data.Limit)

	if err != nil {
		handlers.fundingCreateErrorHandler(err, response)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(201)
	fmt.Fprint(response, string(serializers.SerializeFunding(funding)))
}

// FundingRetrieve is the handler of the funding retrieve entrypoint
func (handlers *Handlers) FundingRetrieve(response http.ResponseWriter, request *http.Request, params httprouter.Params) {
	defer handlers.catchPanics(response)

	iteractor := usecases.FundingInteractor{Repository: handlers.Application.FundingRepository}
	funding, err := iteractor.Retrieve(params.ByName("id"))
	response.Header().Set("Content-Type", "application/json")

	if err != nil {
		errorData := serializers.ErrorResponseData{
			Type:    "not-found",
			Message: err.Error(),
		}

		response.WriteHeader(http.StatusNotFound)
		fmt.Fprint(response, string(serializers.SerializeErrorResponse(&errorData)))
		return
	}

	fmt.Fprint(response, string(serializers.SerializeFunding(funding)))
}

// Error handler of the funding creation
func (handlers *Handlers) fundingCreateErrorHandler(err error, response http.ResponseWriter) {
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
