package serializers

import (
	"encoding/json"
	"github.com/carlosmaniero/budgetgo/domain"
	"io"
	"time"
)

// TransactionData is the serializable representation of a Transaction Data
//
// This is used to parse the information sended to the Transaction creation
// entrypoint
type TransactionData struct {
	Description string    `json:"description"`
	Amount      float64   `json:"amount"`
	Date        time.Time `json:"date"`
}

// TransactionResponseData is the serializable representation of a Transaction
type TransactionResponseData struct {
	ID          string    `json:"id"`
	Description string    `json:"description"`
	Amount      float64   `json:"amount"`
	Date        time.Time `json:"date"`
}

// TransactionValidationErrorData is the serializable representaion of a
// Transaction validation error
type TransactionValidationErrorData struct {
	Type    string            `json:"type"`
	Message string            `json:"message"`
	Errors  []*FieldErrorData `json:"errors"`
}

// UnserializeTransactionData get an io.Reader and convert it to a TransactionData
//
// This return an error if the input is not an valid json representation of
// the TransactionData
func UnserializeTransactionData(reader io.Reader) (*TransactionData, error) {
	data := TransactionData{}
	err := json.NewDecoder(reader).Decode(&data)
	return &data, err
}

// SerializeTransaction receive a Transaction and return its string representation
func SerializeTransaction(transaction *domain.Transaction) []byte {
	data := TransactionResponseData{
		ID:          transaction.ID,
		Description: transaction.Description,
		Amount:      transaction.Amount,
		Date:        transaction.Date,
	}
	b, _ := json.Marshal(data)
	return b
}

// SerializeTransactionValidationError returns a json string representation of a
// TransactionValidationErrorData
func SerializeTransactionValidationError(data *TransactionValidationErrorData) []byte {
	b, _ := json.Marshal(data)
	return b
}
