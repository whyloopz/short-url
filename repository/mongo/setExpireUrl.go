package mongo

import (
	"context"
	"github.com/nqmt/gotime/v2"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func (m Mongo) SetExpireUrl(urls []string) error {
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(m.timeout)*time.Second)

	find := bson.D{{"_id", bson.D{{"$in", urls}}}}
	queryUpdate := bson.D{{"$set", bson.D{{"expireAt", gotime.Now().Unix()}}}}

	_, err := m.shortUrlCollection.UpdateMany(ctx, find, queryUpdate)
	if err != nil {
		return ErrInternalServerSetExpireUrlMongoDB.WithCause(err)
	}

	return nil
}
