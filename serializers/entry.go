package serializers

import (
	"encoding/json"
	. "github.com/carlosmaniero/budgetgo/errors"
	. "github.com/carlosmaniero/budgetgo/models"
	"io"
	"strconv"
	"time"
)

type EntryData struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Amount  string `json:"amount"`
	Date    string `json:"date"`
	Comment string `json:"comment"`
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
	data := EntryData{
		Id:      serializer.Entry.Id,
		Amount:  strconv.FormatFloat(serializer.Entry.Amount, 'f', 6, 64),
		Name:    serializer.Entry.Name,
		Comment: serializer.Entry.Comment,
	}

	return json.Marshal(data)
}

func StringToEntryData(body io.Reader) (EntryData, error) {
	data := EntryData{}
	decoder := json.NewDecoder(body)
	err := decoder.Decode(&data)

	return data, err
}

func EntryDataToEntry(entryData *EntryData) *Entry {
	amount, _ := strconv.ParseFloat(entryData.Amount, 64)
	date, _ := time.Parse(time.RFC3339, entryData.Date)

	return &Entry{
		Id:      entryData.Id,
		Name:    entryData.Name,
		Amount:  amount,
		Date:    date,
		Comment: entryData.Comment,
	}
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
