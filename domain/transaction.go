package domain

import "time"

// Transaction is a transaction representation
//
// An exemple of a transaction is a purshase on credit card
type Transaction struct {
	ID          string
	Description string
	Amount      float64
	Date        time.Time
	Funding     Funding
}

// ValidateAmount is an amount validation
func (transaction *Transaction) ValidateAmount() error {
	if transaction.Amount == 0 {
		return &FieldValidationError{"Amount", "can't be equal zero"}
	}
	return nil
}

// ValidateDescription is a description validation
func (transaction *Transaction) ValidateDescription() error {
	if len(transaction.Description) == 0 {
		return &FieldValidationError{"Description", "can't be empty"}
	}
	return nil
}

// ValidateFunding valildate if the funding is valid and if it has an ID
func (transaction *Transaction) ValidateFunding() error {
	if transaction.Funding.Validate() != nil {
		return &FieldValidationError{"Funding", "isn't valid"}
	}
	if len(transaction.Funding.ID) == 0 {
		return &FieldValidationError{"Funding", "need to have an ID"}
	}
	return nil
}

// ValidateDate is a date validation
func (transaction *Transaction) ValidateDate() error {
	dateLimit := time.Now().AddDate(0, -1, 0)

	if transaction.Date.Before(dateLimit) {
		return &FieldValidationError{"Date", "can't be greater than one month"}
	}
	return nil
}

// Validate method Run all validation methods and return a list of errors
// or nil if no error given
func (transaction *Transaction) Validate() []error {
	errors := make([]error, 0)

	if err := transaction.ValidateAmount(); err != nil {
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
