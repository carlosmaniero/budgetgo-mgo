package handlers

import (
	"net/http"

	"github.com/carlosmaniero/budgetgo/domain"
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

	funding := domain.Funding{
		Name:       serializer.Name,
		Amount:     serializer.Amount,
		ClosingDay: serializer.ClosingDay,
		Limit:      serializer.Limit,
	}
	err := iteractor.Register(&funding)

	if err != nil {
		handlers.usecaseErrorHandler(err, response)
		return
	}

	serializer.Loads(&funding)
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
		serializer := serializers.ErrorResponseSerializer{
			Type:    "not-found",
			Message: err.Error(),
		}

		response.WriteHeader(http.StatusNotFound)
		response.Write(serializer.Serialize())
		return
	}

	serializer := serializers.FundingSerializer{}
	serializer.Loads(funding)
	response.Write(serializer.Serialize())
}
