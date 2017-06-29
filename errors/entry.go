package errors

const (
	AlreadyCreatedError = 0
	ValidationError     = 1
)

var errorMessage = map[int]string{
	AlreadyCreatedError: "This entry was already created",
	ValidationError:     "There is validation errors",
}

type EntryError struct {
	Code        int
	FieldErrors map[string]string
}

func (err *EntryError) Error() string {
	return errorMessage[err.Code]
}
