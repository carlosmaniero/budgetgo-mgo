package domain

// FieldValidationError is an error Representation of a validation error
//
// When a field isn't valid, domain validation functions returns this struct
type FieldValidationError struct {
	Field   string
	Message string
}

func (err *FieldValidationError) Error() string {
	return "The \"" + err.Field + "\" field " + err.Message
}
