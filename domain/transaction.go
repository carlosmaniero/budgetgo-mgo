package domain

type Transaction struct {
	Description string
	Amount      float64
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

	if len(errors) == 0 {
		return nil
	}
	return errors
}
