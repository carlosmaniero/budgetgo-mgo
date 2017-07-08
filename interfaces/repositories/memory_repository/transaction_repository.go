package memory_repository

import (
	"github.com/carlosmaniero/budgetgo/domain"
	"github.com/carlosmaniero/budgetgo/usecases"
)

type MemoryRepository struct {
	transactions []*domain.Transaction
}

func (repository *MemoryRepository) Store(transaction *domain.Transaction) {
	repository.transactions = append(repository.transactions, transaction)
}

func NewMemoryRepository() usecases.TransactionRepository {
	return &MemoryRepository{make([]*domain.Transaction, 0)}
}