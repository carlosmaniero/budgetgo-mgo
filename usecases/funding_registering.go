package usecases

import (
	"github.com/carlosmaniero/budgetgo/domain"
)

type FundingInteractor struct {
	Repository FundingRepository
}

func (iterator *FundingInteractor) Register(name string, amount float64, closingDay int, limit float64) (*domain.Funding, error) {
	funding := domain.Funding{
		Name:       name,
		Amount:     amount,
		ClosingDay: closingDay,
		Limit:      limit,
	}

	if errs := funding.Validate(); errs != nil {
		err := FundingValidationErrors{errs}
		return nil, &err
	}

	funding.Id = iterator.Repository.Store(&funding)
	return &funding, nil
}

type FundingRepository interface {
	Store(*domain.Funding) string
}

type FundingValidationErrors struct {
	Errors []error
}

func (err *FundingValidationErrors) Error() string {
	return "The funding has validation errors"
}
