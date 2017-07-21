package usecases

import (
	"errors"
	"time"

	"github.com/carlosmaniero/budgetgo/domain"
)

var (
	// ErrTransactionNotFound is the error returned when the funding was not found in
	// the repository
	ErrTransactionNotFound = errors.New("the transaction was not found")
	// ErrInvalidMonth is the error returned when the month passed is invalid
	ErrInvalidMonth = errors.New("this is a invalid month")
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

// Retrieve return a stored transaction of the database
func (iterator *TransactionInteractor) Retrieve(id string) (*domain.Transaction, error) {
	if transaction := iterator.Repository.FindByID(id); transaction != nil {
		return transaction, nil
	}

	return nil, ErrTransactionNotFound
}

// RetriveByFundingMonth return a list of transactions in a period
func (iterator *TransactionInteractor) RetriveByFundingMonth(funding *domain.Funding, year int, month int) (TransactionList, error) {
	if month > 12 || month <= 0 {
		return nil, ErrInvalidMonth
	}
	start := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Local)
	end := start.AddDate(0, 1, -1)
	return iterator.Repository.FindByFundingAndInterval(funding, start, end), nil
}

// TransactionRepository is the transaction repository specification
type TransactionRepository interface {
	Store(*domain.Transaction) string
	FindByID(string) *domain.Transaction
	FindByFundingAndInterval(*domain.Funding, time.Time, time.Time) TransactionList
}

// TransactionList is the iterator specification of a transaction list
type TransactionList interface {
	Next(*domain.Transaction) bool
}
