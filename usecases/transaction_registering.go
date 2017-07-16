package usecases

import (
	"github.com/carlosmaniero/budgetgo/domain"
	"time"
)

// TransactionInteractor contains the transaction usecases
type TransactionInteractor struct {
	Repository TransactionRepository
}

// Register a transaction into the repository if transaction is valid
func (iterator *TransactionInteractor) Register(description string, amount float64, date time.Time, funding domain.Funding) (*domain.Transaction, error) {
	transaction := domain.Transaction{
		Description: description,
		Amount:      amount,
		Date:        date,
		Funding:     funding,
	}

	if errs := transaction.Validate(); errs != nil {
		err := ValidationErrors{errs}
		return nil, &err
	}

	id := iterator.Repository.Store(&transaction)
	transaction.ID = id
	return &transaction, nil
}

// TransactionRepository is the transaction repository specification
type TransactionRepository interface {
	Store(*domain.Transaction) string
}
