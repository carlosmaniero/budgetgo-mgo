package usecases

import (
	"errors"

	"github.com/carlosmaniero/budgetgo/domain"
)

// FundingInteractor contains all usecases related to funding
type FundingInteractor struct {
	Repository FundingRepository
}

// Register a funding into the repository if it is valid
func (iterator *FundingInteractor) Register(name string, amount float64, closingDay int, limit float64) (*domain.Funding, error) {
	funding := domain.Funding{
		Name:       name,
		Amount:     amount,
		ClosingDay: closingDay,
		Limit:      limit,
	}

	if errs := funding.Validate(); errs != nil {
		err := ValidationErrors{errs}
		return nil, &err
	}

	funding.ID = iterator.Repository.Store(&funding)
	return &funding, nil
}

// Retrieve a funding in the repository
//
// If the funding was not found, it returns the ErrFundingNotFound error
func (iterator *FundingInteractor) Retrieve(id string) (*domain.Funding, error) {
	if funding := iterator.Repository.FindByID(id); funding != nil {
		return funding, nil
	}

	return nil, ErrFundingNotFound
}

// ErrFundingNotFound is the error returned when the funding was not found in
// the repository
var ErrFundingNotFound = errors.New("the funding was not found")

// FundingRepository interface contains the specification of an funding
// repository
type FundingRepository interface {
	Store(*domain.Funding) string
	FindByID(string) *domain.Funding
}
