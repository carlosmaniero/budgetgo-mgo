package memory_repository

import (
	"github.com/carlosmaniero/budgetgo/domain"
	"github.com/carlosmaniero/budgetgo/usecases"
)

type MemoryFundingRepository struct {
	funding []*domain.Funding
}

func (repository *MemoryFundingRepository) Store(funding *domain.Funding) {
	repository.funding = append(repository.funding, funding)
}

func NewMemoryFundingRepository() usecases.FundingRepository {
	return &MemoryFundingRepository{make([]*domain.Funding, 0)}
}
