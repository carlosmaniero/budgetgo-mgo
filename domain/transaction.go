package domain

type Transaction struct {
	Description string
	Amount      float64
}

type ValidationError struct {
	message string
}

func (err *ValidationError) Error() string {
	return err.message
}

func (entry *Transaction) AmoutValidate() error {
	if entry.Amount == 0 {
		return &ValidationError{"Amount can't be equal zero"}
	}
	return nil
}

func (entry *Transaction) DescriptionValidate() error {
	if len(entry.Description) == 0 {
		return &ValidationError{"Description can't be empty"}
	}
	return nil
}

func (entry *Transaction) Validate() []error {
	errors := make([]error, 0)

	if err := entry.AmoutValidate(); err != nil {
		errors = append(errors, err)
	}

	if err := entry.DescriptionValidate(); err != nil {
		errors = append(errors, err)
	}

	if len(errors) == 0 {
		return nil
	}
	return errors
}
