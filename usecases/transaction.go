package usecases

import (
	"errors"

	"github.com/carlosmaniero/budgetgo/domain"
)

// ErrTransactionNotFound is the error returned when the funding was not found in
// the repository
var ErrTransactionNotFound = errors.New("the transaction was not found")

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

// Retrieve return a stored transaction of the database
func (iterator *TransactionInteractor) Retrieve(id string) (*domain.Transaction, error) {
	if transaction := iterator.Repository.FindByID(id); transaction != nil {
		return transaction, nil
	}

	return nil, ErrTransactionNotFound
}

// TransactionRepository is the transaction repository specification
type TransactionRepository interface {
	Store(*domain.Transaction) string
	FindByID(string) *domain.Transaction
}
