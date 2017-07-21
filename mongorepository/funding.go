package mongorepository

import (
	"github.com/carlosmaniero/budgetgo/domain"
	"github.com/carlosmaniero/budgetgo/usecases"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// MongoFundingRepository is the mongodb implementation of the funding
// repository
type MongoFundingRepository struct {
	Collection *mgo.Collection
}

// Store a funding inside the database
func (repository *MongoFundingRepository) Store(funding *domain.Funding) string {
	bid := bson.NewObjectId()

	data := fundingData{}
	data.puts(funding)
	data.ID = bid

	repository.Collection.Insert(data)
	return bid.Hex()
}

// FindByID returns one funding to the database
func (repository *MongoFundingRepository) FindByID(id string) *domain.Funding {
	bid := bson.ObjectIdHex(id)
	funding := domain.Funding{}

	data := fundingData{}
	if err := repository.Collection.FindId(bid).One(&data); err != nil {
		switch err {
		case mgo.ErrNotFound:
			return nil
		default:
			panic(err)
		}
	}

	data.gets(&funding)
	return &funding
}

// NewMongoFundingRepository create a new mongo db funding repository
func NewMongoFundingRepository(collection *mgo.Collection) usecases.FundingRepository {
	return &MongoFundingRepository{Collection: collection}
}

type fundingData struct {
	ID         bson.ObjectId `bson:"_id"`
	Name       string        `bson:"name"`
	Amount     float64       `bson:"amount"`
	ClosingDay int           `bson:"closing_day"`
}

func (data *fundingData) gets(funding *domain.Funding) {
	funding.ID = data.ID.Hex()
	funding.Name = data.Name
	funding.Amount = data.Amount
	funding.ClosingDay = data.ClosingDay
}

func (data *fundingData) puts(funding *domain.Funding) {
	if funding.ID != "" {
		data.ID = bson.ObjectIdHex(funding.ID)
	}
	data.Name = funding.Name
	data.Amount = funding.Amount
	data.ClosingDay = funding.ClosingDay
}
