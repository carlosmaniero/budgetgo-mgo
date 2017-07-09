package serializers

import (
	"encoding/json"
	"github.com/carlosmaniero/budgetgo/domain"
	"io"
	"time"
)

type TransactionData struct {
	Description string    `json:"description"`
	Amount      float64   `json:"amount"`
	Date        time.Time `json:"date"`
}

type TransactionValidationErrorData struct {
	Type    string            `json:"type"`
	Message string            `json:"message"`
	Errors  []*FieldErrorData `json:"errors"`
}

type FieldErrorData struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func UnserializeTransactionData(reader io.Reader) (*TransactionData, error) {
	data := TransactionData{}
	err := json.NewDecoder(reader).Decode(&data)
	return &data, err
}

func SerializeTransaction(transaction *domain.Transaction) []byte {
	data := TransactionData{
		Description: transaction.Description,
		Amount:      transaction.Amount,
		Date:        transaction.Date,
	}
	b, _ := json.Marshal(data)
	return b
}

func SerializeTransactionValidationError(data *TransactionValidationErrorData) []byte {
	b, _ := json.Marshal(data)
	return b
}
