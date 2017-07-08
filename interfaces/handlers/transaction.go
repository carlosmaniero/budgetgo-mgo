package handlers

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"github.com/carlosmaniero/budgetgo/usecases"
	"github.com/carlosmaniero/budgetgo/domain"
	"github.com/carlosmaniero/budgetgo/interfaces/serializers"
)

func (handlers *Handlers) TransactionCreate(response http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	iterator := usecases.TransactionInteractor{Repository: handlers.Application.TransactionRepository}
	funding := domain.Funding{Name: "Default funding", Limit: 1000, Amount: 0, ClosingDay: 1}

	data, err := serializers.UnserializeTransactionData(request.Body)

	if err != nil {
		handlers.UnserializerErrorHandler(err, response)
		return
	}
	_, transaction := iterator.Register(data.Description, data.Amount, funding)

	response.Header().Set("Content-Type", "application/json")
	fmt.Fprint(response, string(serializers.SerializeTransaction(transaction)))
}
