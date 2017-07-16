package memoryrepository

import (
	"strconv"

	"github.com/carlosmaniero/budgetgo/domain"
	"github.com/carlosmaniero/budgetgo/usecases"
)

// MemoryTransactionRepository implements the usecases.TransactionRepository
//
// This struct will use the memory to store and retrieve all Transactions
type MemoryTransactionRepository struct {
	transactions []*domain.Transaction
}

// Store a Transaction
//
// This mathod raises an panic if the number of 5 transactions is exceeded
func (repository *MemoryTransactionRepository) Store(transaction *domain.Transaction) string {
	repository.transactions = append(repository.transactions, transaction)
	return strconv.Itoa(len(repository.transactions))
}

// NewMemoryTransactionRepository Create a new transaction memory repository
func NewMemoryTransactionRepository() usecases.TransactionRepository {
	return &MemoryTransactionRepository{make([]*domain.Transaction, 0)}
}
