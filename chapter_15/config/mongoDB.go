package config

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoConfig() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI("mongodb+srv://barb:12345678xyz@mycluster.rflsp.mongodb.net/?retryWrites=true&w=majority&appName=mycluster")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		return nil, err
	}

	err = client.Ping(context.Background(), nil)

	if err != nil {
		return nil, err
	}

	return client, nil
}
