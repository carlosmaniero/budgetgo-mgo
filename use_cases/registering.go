package use_cases

import (
	. "github.com/carlosmaniero/budgetgo/domain"
)

type TransactionRepository interface {
	Store(*Transaction)
}

type TransactionInteractor struct {
	Repository TransactionRepository
}

func (iterator *TransactionInteractor) Register(description string, amount float64) (error, *Transaction) {
	transaction := Transaction{Description: description, Amount: amount}
	iterator.Repository.Store(&transaction)
	return nil, &transaction
}
