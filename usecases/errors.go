package usecases

// ValidationErrors contains the validation errors of an transaction
type ValidationErrors struct {
	Errors []error
}

func (err *ValidationErrors) Error() string {
	return "This entity is not valid"
}
