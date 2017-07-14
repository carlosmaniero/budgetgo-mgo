package handlers

import (
	"fmt"
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

	data, err := serializers.UnserializeTransactionData(request.Body)

	if err != nil {
		handlers.unserializerErrorHandler(err, response)
		return
	}

	transaction, err := iterator.Register(data.Description, data.Amount, data.Date, funding)

	if err != nil {
		handlers.transactionCreateErrorHandler(err, response)
		return
	}

	response.WriteHeader(http.StatusCreated)
	response.Header().Set("Content-Type", "application/json")
	fmt.Fprint(response, string(serializers.SerializeTransaction(transaction)))
}

// This is the error handler of the transaction creation
func (handlers *Handlers) transactionCreateErrorHandler(err error, response http.ResponseWriter) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusBadRequest)

	switch err := err.(type) {
	case *usecases.TransactionValidationErrors:

		errorResponse := serializers.TransactionValidationErrorData{
			Type:    "validation_error",
			Message: err.Error(),
			Errors:  handlers.convertFieldValidationErrors(err.Errors),
		}

		data := serializers.SerializeTransactionValidationError(&errorResponse)
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
