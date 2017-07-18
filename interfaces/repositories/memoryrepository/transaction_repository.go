package memoryrepository

import (
	"strconv"
	"sync"

	"github.com/carlosmaniero/budgetgo/domain"
	"github.com/carlosmaniero/budgetgo/usecases"
)

// MemoryTransactionRepository implements the usecases.TransactionRepository
//
// This struct will use the memory to store and retrieve all Transactions
type MemoryTransactionRepository struct {
	transactions []*domain.Transaction
	mux          sync.Mutex
}

// Store a Transaction
//
// This mathod raises an panic if the number of 5 transactions is exceeded
func (repository *MemoryTransactionRepository) Store(transaction *domain.Transaction) string {
	repository.mux.Lock()
	repository.transactions = append(repository.transactions, transaction)
	id := strconv.Itoa(len(repository.transactions))
	repository.mux.Unlock()
	return id
}

// FindByID returns an transaction by ID
func (repository *MemoryTransactionRepository) FindByID(ID string) *domain.Transaction {
	return nil
}

// NewMemoryTransactionRepository Create a new transaction memory repository
func NewMemoryTransactionRepository() usecases.TransactionRepository {
	return &MemoryTransactionRepository{transactions: make([]*domain.Transaction, 0)}
}
