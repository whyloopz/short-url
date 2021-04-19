package mongo

import (
	"context"
	"github.com/nqmt/goerror"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

var (
	ErrInternalServerSaveShortUrlMongoDB       = goerror.DefineInternalServerError("ErrInternalServerSaveShortUrlMongoDB", "internal server error")
	ErrInternalServerGetShortUrlsFindMongoDB   = goerror.DefineInternalServerError("ErrInternalServerSaveShortUrlMongoDB", "internal server error")
	ErrInternalServerGetShortUrlsCursorMongoDB = goerror.DefineInternalServerError("ErrInternalServerSaveShortUrlMongoDB", "internal server error")
	ErrInternalServerGetOriginUrlMongoDB       = goerror.DefineInternalServerError("ErrInternalServerGetOriginUrlMongoDB", "internal server error")
	ErrInternalServerIncrementHitMongoDB       = goerror.DefineInternalServerError("ErrInternalServerIncrementHitMongoDB", "internal server error")
	ErrInternalServerSetExpireUrlMongoDB       = goerror.DefineInternalServerError("ErrInternalServerIncrementHitMongoDB", "internal server error")
	ErrGoneShortUrlExpired                     = goerror.DefineGone("ErrGoneShortUrlExpired", "url is expire")
	ErrNotFoundShortUrl                        = goerror.DefineGone("ErrNotFoundShortUrl", "not found url")
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
	ID        string `json:"_id" bson:"_id"`
	Url       string `json:"url" bson:"url"`
	Hit       int    `json:"hit" bson:"hit"`
	CreatedAt int64  `json:"createAt" bson:"createAt"`
	ExpireAt  int64  `json:"expireAt" bson:"expireAt"`
}

type Mongo struct {
	shortUrlCollection *mongo.Collection
	timeout            int
}

func NewMongoRepo(mongo *mongo.Client, database, collection string, insertTimeout int) *Mongo {
	return &Mongo{
		shortUrlCollection: mongo.Database(database).Collection(collection),
		timeout:            insertTimeout,
	}
}
