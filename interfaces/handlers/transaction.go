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

	iterator := usecases.TransactionInteractor{Repository: handlers.Application.TransactionRepository}
	funding := domain.Funding{ID: "fake-funding", Name: "Default funding", Limit: 1000, Amount: 0, ClosingDay: 1}
	serializer := serializers.TransactionResponseSerializer{}

	if err := serializer.Unserialize(request.Body); err != nil {
		handlers.unserializerErrorHandler(err, response)
		return
	}

	transaction := domain.Transaction{
		Description: serializer.Description,
		Amount:      serializer.Amount,
		Date:        serializer.Date,
		Funding:     funding,
	}

	err := iterator.Register(&transaction)

	if err != nil {
		handlers.usecaseErrorHandler(err, response)
		return
	}

	serializer.Loads(&transaction)

	response.WriteHeader(http.StatusCreated)
	response.Header().Set("Content-Type", "application/json")
	response.Write(serializer.Serialize())
}
