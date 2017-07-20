package serializers

import (
	"encoding/json"
	"github.com/carlosmaniero/budgetgo/domain"
	"io"
	"time"
)

// TransactionResponseSerializer is the serializable representation of a Transaction
type TransactionResponseSerializer struct {
	ID          string             `json:"id"`
	Description string             `json:"description"`
	Amount      float64            `json:"amount"`
	Date        time.Time          `json:"date"`
	FundingID   string             `json:"funding_id,omitempty"`
	Funding     *FundingSerializer `json:"funding"`
}

// Loads the date of a Transaction
func (data *TransactionResponseSerializer) Loads(transaction *domain.Transaction) {
	data.ID = transaction.ID
	data.Description = transaction.Description
	data.Amount = transaction.Amount
	data.Date = transaction.Date
	data.FundingID = transaction.Funding.ID
	data.Funding = &FundingSerializer{}
	data.Funding.Loads(transaction.Funding)
}

// Serialize to json string
func (data *TransactionResponseSerializer) Serialize() []byte {
	b, _ := json.Marshal(data)
	return b
}

// Unserialize a json string representation
//
// This return an error if the input is not an valid json representation of
// the TransactionData
func (data *TransactionResponseSerializer) Unserialize(reader io.Reader) error {
	err := json.NewDecoder(reader).Decode(&data)
	return err
}
