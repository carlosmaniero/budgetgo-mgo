package memory_repository

import (
	"strconv"

	"github.com/carlosmaniero/budgetgo/domain"
	"github.com/carlosmaniero/budgetgo/usecases"
)

type MemoryTransactionRepository struct {
	transactions []*domain.Transaction
}

type MemoryMaxTransactionsError struct{}

func (err *MemoryMaxTransactionsError) Error() string {
	return "5 is the limit of in memory transaction"
}

func (repository *MemoryTransactionRepository) Store(transaction *domain.Transaction) string {
	if len(repository.transactions) == 5 {
		panic(&MemoryMaxTransactionsError{})
	}
	repository.transactions = append(repository.transactions, transaction)
	return strconv.Itoa(len(repository.transactions))
}

func NewMemoryTransactionRepository() usecases.TransactionRepository {
	return &MemoryTransactionRepository{make([]*domain.Transaction, 0)}
}
