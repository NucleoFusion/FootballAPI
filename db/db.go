package db

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToDB() (mongo.Client, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		return mongo.Client{}, err
	}

	return *client, nil
}

func GetCollection(name string, client *mongo.Client) *mongo.Collection {
	return client.Database("premleague").Collection(name)
}
