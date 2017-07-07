package domain

type Transaction struct {
	Description string
	Amount      float64
}

func (entry *Transaction) ValidateAmout() error {
	if entry.Amount == 0 {
		return &FieldValidationError{"Amount", "can't be equal zero"}
	}
	return nil
}

func (entry *Transaction) ValidateDescription() error {
	if len(entry.Description) == 0 {
		return &FieldValidationError{"Description", "can't be empty"}
	}
	return nil
}

func (entry *Transaction) Validate() []error {
	errors := make([]error, 0)

	if err := entry.ValidateAmout(); err != nil {
		errors = append(errors, err)
	}

	if err := entry.ValidateDescription(); err != nil {
		errors = append(errors, err)
	}

	if len(errors) == 0 {
		return nil
	}
	return errors
}
