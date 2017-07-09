package repositories

import (
	"errors"
	"github.com/carlosmaniero/budgetgo/usecases"
)

type transactionRepositoryConstructor func() usecases.TransactionRepository

var transactionRepositories = make(map[string]transactionRepositoryConstructor)

func NewTransactionRepository(engine string) (usecases.TransactionRepository, error) {
	repository, ok := transactionRepositories[engine]
	if !ok {
		return nil, errors.New("The engine " + engine + "was not found")
	}
	return repository(), nil
}

func RegisterTransactionRepository(engine string, repository transactionRepositoryConstructor) {
	transactionRepositories[engine] = repository
}
