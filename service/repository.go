package service

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

type mongoRepository struct {
	timeout  int64
	database string
	client   *mongo.Client
}



func newMongoClient(dbString string, timeout time.Duration) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbString))
	if err != nil {
		return nil, err
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	return client, nil
}

func NewMongoRepository(dbString, db string, timeout int64) (MongoRepository, error) {

	client, err := newMongoClient(dbString, time.Duration(timeout)*time.Second)
	if err != nil {
		return nil, err
	}
	mongoRepo := &mongoRepository{
		timeout:  10,
		database: db,
		client:   client,
	}
	return mongoRepo, nil
}
