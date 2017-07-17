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
	Funding     *Funding
}

// validateAmount is an amount validation
func (transaction *Transaction) validateAmount() error {
	if transaction.Amount == 0 {
		return &FieldValidationError{"Amount", "can't be equal zero"}
	}
	return nil
}

// validateDescription is a description validation
func (transaction *Transaction) validateDescription() error {
	if len(transaction.Description) == 0 {
		return &FieldValidationError{"Description", "can't be empty"}
	}
	return nil
}

// validateFunding valildate if the funding is valid and if it has an ID
func (transaction *Transaction) validateFunding() error {
	if transaction.Funding.Validate() != nil {
		return &FieldValidationError{"Funding", "isn't valid"}
	}
	if len(transaction.Funding.ID) == 0 {
		return &FieldValidationError{"Funding", "need to have an ID"}
	}
	return nil
}

// validateDate is a date validation
func (transaction *Transaction) validateDate() error {
	dateLimit := time.Now().AddDate(-1, 0, 0)

	if transaction.Date.Before(dateLimit) {
		return &FieldValidationError{"Date", "can't be greater than one month"}
	}
	return nil
}

// Validate method Run all validation methods and return a list of errors
// or nil if no error given
func (transaction *Transaction) Validate() []error {
	errors := make([]error, 0)

	if err := transaction.validateAmount(); err != nil {
		errors = append(errors, err)
	}

	if err := transaction.validateDescription(); err != nil {
		errors = append(errors, err)
	}

	if err := transaction.validateFunding(); err != nil {
		errors = append(errors, err)
	}

	if err := transaction.validateDate(); err != nil {
		errors = append(errors, err)
	}

	if len(errors) == 0 {
		return nil
	}
	return errors
}
