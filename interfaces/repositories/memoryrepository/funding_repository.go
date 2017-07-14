package memoryrepository

import (
	"strconv"

	"github.com/carlosmaniero/budgetgo/domain"
	"github.com/carlosmaniero/budgetgo/usecases"
)

// MemoryFundingRepository implements the usecases.FundingRepository
//
// This struct will use the memory to store and retrieve all fundings
type MemoryFundingRepository struct {
	fundings []*domain.Funding
}

// Store a Funding
func (repository *MemoryFundingRepository) Store(funding *domain.Funding) string {
	repository.fundings = append(repository.fundings, funding)
	return strconv.Itoa(len(repository.fundings))
}

// FindByID a Funding. It will return nil if the Funding was not founded
func (repository *MemoryFundingRepository) FindByID(id string) *domain.Funding {
	index, err := strconv.Atoi(id)

	if err != nil {
		return nil
	}

	return repository.fundings[index-1]
}

// NewMemoryFundingRepository Create a new funding memory repository
func NewMemoryFundingRepository() usecases.FundingRepository {
	return &MemoryFundingRepository{make([]*domain.Funding, 0)}
}
