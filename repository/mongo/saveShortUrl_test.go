package mongo

import (
	"context"
	"github.com/nqmt/gotime/v2"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
	"time"
)

func TestMongo_SaveShortUrl(t *testing.T) {
	suite := NewMongoRepoTest()
	suite.Setup()

	gotime.Freeze(time.Date(2020, 06, 6, 6, 6, 6, 6, time.UTC))
	err := suite.repo.SaveShortUrl("test04", "http://www.google.com", 1650334869)
	require.NoError(t, err)

	savedData := new(ShortUrlModel)
	err = suite.mongo.
		Database(suite.databaseName).
		Collection(suite.collectionName).
		FindOne(context.Background(), bson.D{{"_id", "test04"}}).
		Decode(savedData)
	require.NoError(t, err)

	expect := &ShortUrlModel{
		ID:        "test04",
		Url:       "http://www.google.com",
		Hit:       0,
		CreatedAt: 1591423566,
		ExpireAt:  1650334869,
	}
	require.Equal(t, expect, savedData)
}
