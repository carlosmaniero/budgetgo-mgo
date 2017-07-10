package application

import (
	"github.com/carlosmaniero/budgetgo/interfaces/repositories/memory_repository"
	"github.com/carlosmaniero/budgetgo/usecases"
)

type Application struct {
	TransactionRepository usecases.TransactionRepository
	FundingRepository     usecases.FundingRepository
}

func Init() *Application {
	return &Application{
		TransactionRepository: memory_repository.NewMemoryTransactionRepository(),
		FundingRepository:     memory_repository.NewMemoryFundingRepository(),
	}
}
