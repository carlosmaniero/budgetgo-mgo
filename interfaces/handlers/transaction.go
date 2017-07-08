package handlers

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"github.com/carlosmaniero/budgetgo/usecases"
	"github.com/carlosmaniero/budgetgo/domain"
)

func (handlers *Handlers) TransactionCreate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	iterator := usecases.TransactionInteractor{Repository: handlers.Application.TransactionRepository}
	funding := domain.Funding{Name: "Default funding", Limit: 1000, Amount: 0, ClosingDay: 1}
	iterator.Register("4 beers", 10.5, funding)

	fmt.Fprint(w, "Transaction created")
}
