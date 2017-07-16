package usecases

import (
	"github.com/carlosmaniero/budgetgo/domain"
)

// TransactionInteractor contains the transaction usecases
type TransactionInteractor struct {
	Repository TransactionRepository
}

// Register a transaction into the repository if transaction is valid
func (iterator *TransactionInteractor) Register(transaction *domain.Transaction) error {
	if errs := transaction.Validate(); errs != nil {
		err := ValidationErrors{errs}
		return &err
	}

	id := iterator.Repository.Store(transaction)
	transaction.ID = id
	return nil
}

// TransactionRepository is the transaction repository specification
type TransactionRepository interface {
	Store(*domain.Transaction) string
}
