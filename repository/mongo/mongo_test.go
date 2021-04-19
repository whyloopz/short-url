package mongo

import (
	"context"
	"github.com/nqmt/gotime/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type MongoRepoTestSuite struct {
	ctx            context.Context
	mongo          *mongo.Client
	repo           *Mongo
	databaseName   string
	collectionName string
}

func NewMongoRepoTest() *MongoRepoTestSuite {
	m := ConnectMongo("mongodb://localhost:7002")

	return &MongoRepoTestSuite{
		ctx:            context.Background(),
		mongo:          m,
		repo:           NewMongoRepo(m, "test", "shortUrl", 1),
		databaseName:   "test",
		collectionName: "shortUrl",
	}
}

func (ts MongoRepoTestSuite) TearDown() {
	if err := ts.mongo.Database("test").Drop(context.Background()); err != nil {
		panic(err)
	}
}

func (ts MongoRepoTestSuite) Seed() {
	gotime.Freeze(time.Date(2020, 06, 6, 6, 6, 6, 6, time.UTC))
	ts.repo.SaveShortUrl("test00", "http://www.google1.com", 1650334869)
	ts.repo.SaveShortUrl("test01", "http://www.google2.com", 1650334869)
	ts.repo.SaveShortUrl("test02", "http://www.google3.com", 1)
}

func (ts MongoRepoTestSuite) Setup() {
	ts.TearDown()
	ts.Seed()
}
