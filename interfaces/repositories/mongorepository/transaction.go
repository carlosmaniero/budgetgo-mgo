package mongorepository

import (
	"time"

	"github.com/carlosmaniero/budgetgo/domain"
	"github.com/carlosmaniero/budgetgo/usecases"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// MongoTransactionRepository is the mongodb implementation of the transaction
// repository
type MongoTransactionRepository struct {
	Collection *mgo.Collection
}

// Store a transaction inside the database
func (repository *MongoTransactionRepository) Store(transaction *domain.Transaction) string {
	bid := bson.NewObjectId()

	model := transactionData{}
	model.puts(transaction)
	model.ID = bid

	repository.Collection.Insert(model)
	return bid.Hex()
}

// FindByID returns one transaction to the database
func (repository *MongoTransactionRepository) FindByID(id string) *domain.Transaction {
	bid := bson.ObjectIdHex(id)
	transaction := domain.Transaction{}
	data := transactionData{}

	if err := repository.Collection.FindId(bid).One(&data); err != nil {
		switch err {
		case mgo.ErrNotFound:
			return nil
		default:
			panic(err)
		}
	}

	data.gets(&transaction)

	return &transaction
}

type transactionData struct {
	ID          bson.ObjectId `bson:"_id"`
	Description string        `bson:"description"`
	Amount      float64       `bson:"amount"`
	Date        time.Time     `bson:"date"`
	Funding     *fundingData  `bson:"funding"`
}

func (data *transactionData) puts(transaction *domain.Transaction) {
	if transaction.ID != "" {
		data.ID = bson.ObjectIdHex(transaction.ID)
	}
	data.Description = transaction.Description
	data.Amount = transaction.Amount
	data.Date = transaction.Date
	funding := fundingData{}
	funding.puts(transaction.Funding)
	data.Funding = &funding
}

func (data *transactionData) gets(transaction *domain.Transaction) {
	transaction.ID = data.ID.Hex()
	transaction.Description = data.Description
	transaction.Amount = data.Amount
	transaction.Date = data.Date
	funding := domain.Funding{}
	data.Funding.gets(&funding)
	transaction.Funding = &funding
}

// NewMongoTransactionRepository create a new mongo db transaction repository
func NewMongoTransactionRepository(collection *mgo.Collection) usecases.TransactionRepository {
	return &MongoTransactionRepository{Collection: collection}
}
