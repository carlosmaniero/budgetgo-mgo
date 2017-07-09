package usecases

import (
	"github.com/carlosmaniero/budgetgo/domain"
)

type FundingInteractor struct {
	Repository FundingRepository
}

func (iterator *FundingInteractor) Register(name string, amount float64, closingDay int, limit float64) (error, *domain.Funding) {
	funding := domain.Funding{
		Name:       name,
		Amount:     amount,
		ClosingDay: closingDay,
		Limit:      limit,
	}

	if errs := funding.Validate(); errs != nil {
		err := FundingValidationErrors{errs}
		return &err, nil
	}

	iterator.Repository.Store(&funding)
	return nil, nil
}

type FundingValidationErrors struct {
	Errors []error
}

type FundingRepository interface {
	Store(*domain.Funding)
}

func (err *FundingValidationErrors) Error() string {
	return "The funding has validation errors"
}
