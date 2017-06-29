package serializers

import (
	"encoding/json"
	. "github.com/carlosmaniero/budgetgo/errors"
	. "github.com/carlosmaniero/budgetgo/models"
	"io"
	"time"
)

type entryData struct {
	Id      string    `json:"id"`
	Name    string    `json:"name"`
	Amount  float32   `json:"amount,string"`
	Date    time.Time `json:"date"`
	Comment string    `json:"comment"`
}

type entryErrorData struct {
	Code    int               `json:"code"`
	Message string            `json:"message"`
	Fields  map[string]string `json:"fields,omitempty"`
}

type EntrySerializer struct {
	Entry *Entry
}

func (serializer *EntrySerializer) Serialize() ([]byte, error) {
	data := entryData{
		Id:      serializer.Entry.Id,
		Name:    serializer.Entry.Name,
		Amount:  serializer.Entry.Amount,
		Date:    serializer.Entry.Date,
		Comment: serializer.Entry.Comment,
	}

	return json.Marshal(data)
}

func (serializer *EntrySerializer) Unserialize(body io.Reader) error {
	data := entryData{}
	decoder := json.NewDecoder(body)
	err := decoder.Decode(&data)

	if err != nil {
		return err
	}

	serializer.Entry.Id = data.Id
	serializer.Entry.Name = data.Name
	serializer.Entry.Amount = data.Amount
	serializer.Entry.Date = data.Date
	serializer.Entry.Comment = data.Comment

	return err
}

type EntryErrorSerializer struct {
	Error *EntryError
}

func (serializer *EntryErrorSerializer) Serialize() ([]byte, error) {
	data := entryErrorData{
		Code:    serializer.Error.Code,
		Message: serializer.Error.Error(),
		Fields:  serializer.Error.FieldErrors,
	}

	return json.Marshal(data)
}
