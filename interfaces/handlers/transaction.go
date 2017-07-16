package handlers

import (
	"github.com/carlosmaniero/budgetgo/domain"
	"github.com/carlosmaniero/budgetgo/interfaces/serializers"
	"github.com/carlosmaniero/budgetgo/usecases"
	"github.com/julienschmidt/httprouter"
	"net/http"
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

	response.WriteHeader(http.StatusCreated)
	response.Header().Set("Content-Type", "application/json")
	response.Write(serializer.Serialize())
}
