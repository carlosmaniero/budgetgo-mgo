package usecases

import (
	"github.com/carlosmaniero/budgetgo/domain"
	"time"
)

type TransactionInteractor struct {
	Repository TransactionRepository
}

func (iterator *TransactionInteractor) Register(description string, amount float64, date time.Time, funding domain.Funding) (*domain.Transaction, error) {
	transaction := domain.Transaction{
		Description: description,
		Amount:      amount,
		Date:        date,
		Funding:     funding,
	}

	if errs := transaction.Validate(); errs != nil {
		err := TransactionValidationErrors{errs}
		return nil, &err
	}

	id := iterator.Repository.Store(&transaction)
	transaction.Id = id
	return &transaction, nil
}

type TransactionRepository interface {
	Store(*domain.Transaction) string
}

type TransactionValidationErrors struct {
	Errors []error
}

func (err *TransactionValidationErrors) Error() string {
	return "The transaction has validation errors"
}
