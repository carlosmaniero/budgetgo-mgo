package controllers

import (
	. "github.com/carlosmaniero/budgetgo/models"
	. "github.com/carlosmaniero/budgetgo/services"
)

const (
	AlreadyCreatedError = 0
)

var errorMessage = map[int]string{
	AlreadyCreatedError: "This promotion was already created",
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
	service := NewEntryService()
	service.Insert(entry)

	return nil
}
