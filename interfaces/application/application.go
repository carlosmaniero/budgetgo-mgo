package application

import (
	"github.com/carlosmaniero/budgetgo/usecases"
	"github.com/carlosmaniero/budgetgo/interfaces/repositories"
)

type Application struct {
	TransactionRepository usecases.TransactionRepository
}

func Init() *Application {
	registerRepositories()

	repository, err := repositories.NewTransactionRepository("memory")

	if err != nil {
		panic(err)
	}

	return &Application{
		TransactionRepository: repository,
	}
}