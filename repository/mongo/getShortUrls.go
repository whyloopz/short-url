package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func (m Mongo) GetShortUrls(searchShortUrl, searchUrl string) ([]*ShortUrlModel, error) {
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(m.timeout)*time.Second)

	regexes := make(bson.A, 0, 2)

	if searchShortUrl != "" {
		regexes = append(regexes, bson.M{
			"_id": bson.M{
				"$regex": searchShortUrl,
			},
		})
	}
	if searchUrl != "" {
		regexes = append(regexes, bson.M{
			"url": bson.M{
				"$regex": searchUrl,
			},
		})
	}

	filter := bson.D{{}}
	if len(regexes) != 0 {
		filter = bson.D{{"$or", regexes}}
	}

	cursor, err := m.shortUrlCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var shortUrls []*ShortUrlModel
	if err := cursor.All(ctx, &shortUrls); err != nil {
		return []*ShortUrlModel{}, err
	}

	return shortUrls, nil
}
