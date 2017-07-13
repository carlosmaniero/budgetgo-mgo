package memory_repository

import (
	"strconv"

	"github.com/carlosmaniero/budgetgo/domain"
	"github.com/carlosmaniero/budgetgo/usecases"
)

type MemoryFundingRepository struct {
	fundings []*domain.Funding
}

func (repository *MemoryFundingRepository) Store(funding *domain.Funding) string {
	repository.fundings = append(repository.fundings, funding)
	return strconv.Itoa(len(repository.fundings))
}

func (m *MemoryFundingRepository) FindById(string) *domain.Funding {
	panic("not implemented")
}

func NewMemoryFundingRepository() usecases.FundingRepository {
	return &MemoryFundingRepository{make([]*domain.Funding, 0)}
}
