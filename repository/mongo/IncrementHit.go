package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func (m Mongo) IncrementHit(shortUrl string) error {
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(m.timeout)*time.Second)

	find := bson.D{{"_id", shortUrl}}
	queryUpdate := bson.D{{"$inc", bson.D{{"hit", 1}}}}
	_, err := m.shortUrlCollection.UpdateOne(ctx, find, queryUpdate)
	if err != nil {
		return ErrInternalServerIncrementHitMongoDB.WithCause(err)
	}

	return nil
}
