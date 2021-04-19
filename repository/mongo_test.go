package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepoTestSuite struct {
	ctx            context.Context
	mongo          *mongo.Client
	repo           *Mongo
	databaseName   string
	collectionName string
}

func NewMongoRepoTest() *MongoRepoTestSuite {
	m := ConnectMongo("mongodb://localhost:6002")

	return &MongoRepoTestSuite{
		ctx:            context.Background(),
		mongo:          m,
		repo:           NewMongoRepo(m, "test", "shortUrl", 1),
		databaseName:   "test",
		collectionName: "shortUrl",
	}
}

func (ts *MongoRepoTestSuite) TearDown() {
	if err := ts.mongo.Database("test").Drop(context.Background()); err != nil {
		panic(err)
	}
}

func (ts *MongoRepoTestSuite) Setup() {
	ts.TearDown()
}
