package handlers

import (
	"net/http"

	"github.com/carlosmaniero/budgetgo/domain"
	"github.com/carlosmaniero/budgetgo/interfaces/serializers"
	"github.com/carlosmaniero/budgetgo/usecases"
	"github.com/julienschmidt/httprouter"
)

// TransactionCreate is the handler of the transaction creation entrypoint
func (handlers *Handlers) TransactionCreate(response http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	defer handlers.catchPanics(response)

	serializer := serializers.TransactionResponseSerializer{}

	if err := serializer.Unserialize(request.Body); err != nil {
		handlers.unserializerErrorHandler(err, response)
		return
	}

	iterator := usecases.TransactionInteractor{Repository: handlers.Application.TransactionRepository}
	fundingIteractor := usecases.FundingInteractor{Repository: handlers.Application.FundingRepository}

	funding, err := fundingIteractor.Retrieve(serializer.FundingID)

	if err != nil {
		handlers.usecaseErrorHandler(err, response)
		return
	}

	transaction := domain.Transaction{
		Description: serializer.Description,
		Amount:      serializer.Amount,
		Date:        serializer.Date,
		Funding:     funding,
	}

	err = iterator.Register(&transaction)

	if err != nil {
		handlers.usecaseErrorHandler(err, response)
		return
	}

	serializer.Loads(&transaction)

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusCreated)
	response.Write(serializer.Serialize())
}

// TransactionRetrieve is the handler that gets an transaction
func (handlers *Handlers) TransactionRetrieve(response http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id := params.ByName("id")

	iterator := usecases.TransactionInteractor{Repository: handlers.Application.TransactionRepository}
	transaction, err := iterator.Retrieve(id)

	if err != nil {
		serializer := serializers.ErrorResponseSerializer{
			Type:    "not-found",
			Message: err.Error(),
		}

		response.WriteHeader(http.StatusNotFound)
		response.Write(serializer.Serialize())
		return
	}

	serializer := serializers.TransactionResponseSerializer{}
	serializer.Loads(transaction)
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	response.Write(serializer.Serialize())
}
