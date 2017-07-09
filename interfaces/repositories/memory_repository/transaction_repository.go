package memory_repository

import (
	"github.com/carlosmaniero/budgetgo/domain"
	"github.com/carlosmaniero/budgetgo/usecases"
)

type MemoryRepository struct {
	transactions []*domain.Transaction
}

type MemoryMaxTransactionsError struct{}

func (err *MemoryMaxTransactionsError) Error() string {
	return "5 is the limit of in memory transaction"
}

func (repository *MemoryRepository) Store(transaction *domain.Transaction) {
	if len(repository.transactions) == 5 {
		panic(&MemoryMaxTransactionsError{})
	}
	repository.transactions = append(repository.transactions, transaction)
}

func NewMemoryRepository() usecases.TransactionRepository {
	return &MemoryRepository{make([]*domain.Transaction, 0)}
}
