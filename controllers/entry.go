package controllers

import (
	. "github.com/carlosmaniero/budgetgo/models"
	. "github.com/carlosmaniero/budgetgo/services"
	. "github.com/carlosmaniero/budgetgo/validators"
)

const (
	AlreadyCreatedError = 0
	ValidationError     = 1
)

var errorMessage = map[int]string{
	AlreadyCreatedError: "This entry was already created",
	ValidationError:     "There is validation errors",
}

type EntryError struct {
	Code int
}

func (err *EntryError) Error() string {
	return errorMessage[err.Code]
}

func EntryCreateController(entry *Entry) *EntryError {
	if entry.Id != "" {
		return &EntryError{Code: AlreadyCreatedError}
	}
	validator := NewEntryValidator(entry)
	ok := validator.IsValid()

	if !ok {
		return &EntryError{Code: ValidationError}
	}

	service := NewEntryService()
	service.Insert(entry)

	return nil
}
