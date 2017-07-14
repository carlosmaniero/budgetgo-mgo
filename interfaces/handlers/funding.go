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
	serializer := serializers.FundingSerializer{}

	if err := serializer.Unserialize(request.Body); err != nil {
		handlers.unserializerErrorHandler(err, response)
		return
	}

	funding, err := iteractor.Register(
		serializer.Name,
		serializer.Amount,
		serializer.ClosingDay,
		serializer.Limit,
	)

	if err != nil {
		handlers.fundingCreateErrorHandler(err, response)
		return
	}

	serializer.Loads(funding)
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(201)
	response.Write(serializer.Serialize())
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

	serializer := serializers.FundingSerializer{}
	serializer.Loads(funding)
	response.Write(serializer.Serialize())
}

// Error handler of the funding creation
func (handlers *Handlers) fundingCreateErrorHandler(err error, response http.ResponseWriter) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusBadRequest)

	switch err := err.(type) {
	case *usecases.FundingValidationErrors:

		serializer := serializers.FundingValidationErrorData{
			Type:    "validation_error",
			Message: err.Error(),
			Errors:  handlers.convertFieldValidationErrors(err.Errors),
		}
		response.Write(serializer.Serialize())

	default:
		errorResponse := serializers.ErrorResponseData{
			Type:    "domain_error",
			Message: err.Error(),
		}

		data := serializers.SerializeErrorResponse(&errorResponse)
		fmt.Fprint(response, string(data))
	}
}
