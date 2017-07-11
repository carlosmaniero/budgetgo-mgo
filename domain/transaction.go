package domain

import "time"

type Transaction struct {
	Id          string
	Description string
	Amount      float64
	Date        time.Time
	Funding     Funding
}

func (transaction *Transaction) ValidateAmout() error {
	if transaction.Amount == 0 {
		return &FieldValidationError{"Amount", "can't be equal zero"}
	}
	return nil
}

func (transaction *Transaction) ValidateDescription() error {
	if len(transaction.Description) == 0 {
		return &FieldValidationError{"Description", "can't be empty"}
	}
	return nil
}

func (transaction *Transaction) ValidateFunding() error {
	if transaction.Funding.Validate() != nil {
		return &FieldValidationError{"Funding", "isn't valid"}
	}
	if len(transaction.Funding.Id) == 0 {
		return &FieldValidationError{"Funding", "need to have an ID"}
	}
	return nil
}

func (transaction *Transaction) ValidateDate() error {
	dateLimit := time.Now().AddDate(0, -1, 0)

	if transaction.Date.Before(dateLimit) {
		return &FieldValidationError{"Date", "can't be greater than one month"}
	}
	return nil
}

func (transaction *Transaction) Validate() []error {
	errors := make([]error, 0)

	if err := transaction.ValidateAmout(); err != nil {
		errors = append(errors, err)
	}

	if err := transaction.ValidateDescription(); err != nil {
		errors = append(errors, err)
	}

	if err := transaction.ValidateFunding(); err != nil {
		errors = append(errors, err)
	}

	if err := transaction.ValidateDate(); err != nil {
		errors = append(errors, err)
	}

	if len(errors) == 0 {
		return nil
	}
	return errors
}
