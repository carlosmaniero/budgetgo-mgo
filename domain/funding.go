package domain

// Funding is the entity representation of an Account Funding
//
// An example of a Funding is a "Credit Card". In a credit card we have
// a limit of credit, the amount (the total of used credit) and a closing day
// (the invoice closure).
type Funding struct {
	ID         string
	Name       string
	Limit      float64
	Amount     float64
	ClosingDay int
}

// validateName checks if the name is valid
func (funding *Funding) validateName() error {
	if len(funding.Name) == 0 {
		return &FieldValidationError{"Name", "can't be empty"}
	}
	return nil
}

// validateLimit checks if the limit is valid
func (funding *Funding) validateLimit() error {
	if funding.Limit < 0 {
		return &FieldValidationError{"Limit", "can't be negative"}
	}
	return nil
}

// validateClosingDay checks if the closing day is valid
func (funding *Funding) validateClosingDay() error {
	if funding.ClosingDay <= 0 {
		return &FieldValidationError{"ClosingDay", "should be greater than zero"}
	}
	return nil
}

// Validate the Funding structure
//
// This method call all Funding field validations and returns a list of
// errors.
func (funding *Funding) Validate() []error {
	errors := make([]error, 0)

	if err := funding.validateName(); err != nil {
		errors = append(errors, err)
	}

	if err := funding.validateLimit(); err != nil {
		errors = append(errors, err)
	}

	if err := funding.validateClosingDay(); err != nil {
		errors = append(errors, err)
	}

	if len(errors) == 0 {
		return nil
	}
	return errors
}
