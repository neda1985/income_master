package mdb


import (
"context"
"go.mongodb.org/mongo-driver/mongo"
"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)
type MongoDBRepository struct {
	client     *mongo.Client     // MongoDB client
	database   *mongo.Database   // MongoDB database
	collection *mongo.Collection // MongoDB collection (optional)
}

func NewMongoDBRepository() (*MongoDBRepository, error) {
	// Set up MongoDB connection
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	// Set up MongoDB database and collection
	database := client.Database("mydb")
	collection := database.Collection("money_collection")

	return &MongoDBRepository{
		client:     client,
		database:   database,
		collection: collection,
	}, nil
}



func (m MongoDBRepository) AddMoney(ctx context.Context, userId string, amount int, createdAt time.Time) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (m MongoDBRepository) SpendMoney(tx context.Context, userId string, amount int, createdAt time.Time) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (m MongoDBRepository) restMoney(ctx context.Context, userId string) (int error) {
	//TODO implement me
	panic("implement me")
}
