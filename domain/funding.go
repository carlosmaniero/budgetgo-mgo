package domain

type Funding struct {
	Name       string
	Limit      float64
	Amount     float64
	ClosingDay int
}

func (funding *Funding) ValidateName() error {
	if len(funding.Name) == 0 {
		return &FieldValidationError{"Name", "can't be empty"}
	}
	return nil
}

func (funding *Funding) ValidateLimit() error {
	if funding.Limit < 0 {
		return &FieldValidationError{"Limit", "can't be negative"}
	}
	return nil
}

func (funding *Funding) ValidateClosingDay() error {
	if funding.ClosingDay <= 0 {
		return &FieldValidationError{"ClosingDay", "should be greater than zero"}
	}
	return nil
}

func (funding *Funding) Validate() []error {
	errors := make([]error, 0)

	if err := funding.ValidateName(); err != nil {
		errors = append(errors, err)
	}

	if err := funding.ValidateLimit(); err != nil {
		errors = append(errors, err)
	}

	if err := funding.ValidateClosingDay(); err != nil {
		errors = append(errors, err)
	}

	if len(errors) == 0 {
		return nil
	}
	return errors
}
