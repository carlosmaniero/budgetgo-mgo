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
	if len(repository.transactions) == 5 {
		panic(&MemoryMaxTransactionsError{})
	}
	repository.transactions = append(repository.transactions, transaction)
	return strconv.Itoa(len(repository.transactions))
}

// NewMemoryTransactionRepository Create a new transaction memory repository
func NewMemoryTransactionRepository() usecases.TransactionRepository {
	return &MemoryTransactionRepository{make([]*domain.Transaction, 0)}
}

// MemoryMaxTransactionsError is the error raised when the limit of stored
// transaction is exceeded
type MemoryMaxTransactionsError struct{}

func (err *MemoryMaxTransactionsError) Error() string {
	return "5 is the limit of in memory transaction"
}
