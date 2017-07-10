package handlers

import (
	"fmt"
	"github.com/carlosmaniero/budgetgo/domain"
	"github.com/carlosmaniero/budgetgo/interfaces/serializers"
	"github.com/carlosmaniero/budgetgo/usecases"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (handlers *Handlers) TransactionCreate(response http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	defer handlers.catchPanics(response)

	iterator := usecases.TransactionInteractor{Repository: handlers.Application.TransactionRepository}
	funding := domain.Funding{Name: "Default funding", Limit: 1000, Amount: 0, ClosingDay: 1}

	data, err := serializers.UnserializeTransactionData(request.Body)

	if err != nil {
		handlers.UnserializerErrorHandler(err, response)
		return
	}

	err, transaction := iterator.Register(data.Description, data.Amount, data.Date, funding)

	if err != nil {
		handlers.TransactionCreateErrorHandler(err, response)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	fmt.Fprint(response, string(serializers.SerializeTransaction(transaction)))
}

func (handlers *Handlers) TransactionCreateErrorHandler(err error, response http.ResponseWriter) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusBadRequest)

	switch err := err.(type) {
	case *usecases.TransactionValidationErrors:

		errorResponse := serializers.TransactionValidationErrorData{
			Type:    "validation_error",
			Message: err.Error(),
			Errors:  handlers.convertTransactionFieldErrors(err.Errors),
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
