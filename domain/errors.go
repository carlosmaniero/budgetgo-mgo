package domain

type FieldValidationError struct {
	Field   string
	Message string
}

func (err *FieldValidationError) Error() string {
	return "The \"" + err.Field + "\" field " + err.Message
}
