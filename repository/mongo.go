package repository

import (
	"context"
	"github.com/nqmt/gotime/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
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

func NewMongo(mongo *mongo.Client, database, collection string, insertTimeout int) *Mongo {
	return &Mongo{
		shortUrlCollection: mongo.Database(database).Collection(collection),
		insertTimeout:      insertTimeout,
	}
}

func (m Mongo) SaveShortUrl(shortUrl, originUrl string, expireAt int64) error {
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(m.insertTimeout)*time.Second)

	model := &ShortUrlModel{
		ID:        shortUrl,
		Url:       originUrl,
		Hit:       0,
		CreatedAt: gotime.NowUnix(),
		ExpireAt:  expireAt,
	}

	_, err := m.shortUrlCollection.InsertOne(ctx, model)
	if err != nil {
		return err
	}

	return nil
}
