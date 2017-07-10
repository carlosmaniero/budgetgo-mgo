package application

import (
	"github.com/carlosmaniero/budgetgo/interfaces/repositories"
	"github.com/carlosmaniero/budgetgo/interfaces/repositories/memory_repository"
)

func registerRepositories() {
	repositories.RegisterTransactionRepository("memory", memory_repository.NewMemoryTransactionRepository)
}
