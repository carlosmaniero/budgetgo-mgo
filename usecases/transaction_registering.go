package usecases

import (
	"github.com/carlosmaniero/budgetgo/domain"
	"time"
)

type TransactionRepository interface {
	Store(*domain.Transaction)
}

type TransactionInteractor struct {
	Repository TransactionRepository
}

type TransactionValidationErrors struct {
	Errors []error
}

func (err *TransactionValidationErrors) Error() string {
	return "The transaction has validation errors"
}

func (iterator *TransactionInteractor) Register(description string, amount float64, date time.Time, funding domain.Funding) (error, *domain.Transaction) {
	transaction := domain.Transaction{
		Description: description,
		Amount:      amount,
		Date: 		 date,
		Funding:     funding,
	}

	if errs := transaction.Validate(); errs != nil {
		err := TransactionValidationErrors{errs}
		return &err, nil
	}

	iterator.Repository.Store(&transaction)
	return nil, &transaction
}
