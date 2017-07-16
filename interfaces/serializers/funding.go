package serializers

import (
	"encoding/json"
	"io"

	"github.com/carlosmaniero/budgetgo/domain"
)

// FundingSerializer is the serializable representation of a Funding
type FundingSerializer struct {
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	Limit      float64 `json:"limit"`
	Amount     float64 `json:"amount"`
	ClosingDay int     `json:"closing_day"`
}

// Loads load data from an funding
func (data *FundingSerializer) Loads(funding *domain.Funding) {
	data.ID = funding.ID
	data.Name = funding.Name
	data.Amount = funding.Amount
	data.ClosingDay = funding.ClosingDay
	data.Limit = funding.Limit
}

// Unserialize gets an io.Reader and convert it into the serializer
//
// This return an error if the input is not an valid json representation of
// the FundingResponseData
func (data *FundingSerializer) Unserialize(reader io.Reader) error {
	return json.NewDecoder(reader).Decode(&data)
}

// Serialize returns the FundingResposeData string json representation
func (data *FundingSerializer) Serialize() []byte {
	b, _ := json.Marshal(data)
	return b
}
