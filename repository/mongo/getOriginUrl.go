package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func (m Mongo) GetOriginUrl(shortUrl string) (string, error) {
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(m.timeout)*time.Second)

	model := new(ShortUrlModel)
	if err := m.shortUrlCollection.FindOne(ctx, bson.D{{"_id", shortUrl}}).Decode(model); err != nil {
		if err == mongo.ErrNoDocuments {
			return "", ErrNotFoundShortUrl
		}
		return "", ErrInternalServerGetOriginUrlMongoDB.WithCause(err)
	}

	if model.ExpireAt < time.Now().Unix() {
		return "", ErrGoneShortUrlExpired
	}

	return model.Url, nil
}
