package application

import (
	"github.com/carlosmaniero/budgetgo/interfaces/repositories/memoryrepository"
	"github.com/carlosmaniero/budgetgo/usecases"
)

// Application is the application context. This contains all repositories of
// the application.
type Application struct {
	TransactionRepository usecases.TransactionRepository
	FundingRepository     usecases.FundingRepository
}

// New create a new application instance
func New() *Application {
	return &Application{
		TransactionRepository: memoryrepository.NewMemoryTransactionRepository(),
		FundingRepository:     memoryrepository.NewMemoryFundingRepository(),
	}
}
