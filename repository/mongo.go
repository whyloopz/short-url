package repository

import (
	"context"
	"github.com/nqmt/goerror"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

var (
	ErrInternalServerSaveShortUrlMongoDB = goerror.DefineInternalServerError("ErrInternalServerSaveShortUrlMongoDB", "internal server error")
)

func ConnectMongo(url string) *mongo.Client {
	var err error
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		panic(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}

	return client
}

type ShortUrlModel struct {
	ID        string `bson:"_id"`
	Url       string `bson:"url"`
	Hit       int    `bson:"hit"`
	CreatedAt int64  `bson:"createAt"`
	ExpireAt  int64  `bson:"expireAt"`
}

type Mongo struct {
	shortUrlCollection *mongo.Collection
	insertTimeout      int
}

func NewMongoRepo(mongo *mongo.Client, database, collection string, insertTimeout int) *Mongo {
	return &Mongo{
		shortUrlCollection: mongo.Database(database).Collection(collection),
		insertTimeout:      insertTimeout,
	}
}
