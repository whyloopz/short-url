package mongo

import (
	"context"
	"github.com/nqmt/gotime/v2"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
	"time"
)

func TestMongo_ExpireUrl_Success(t *testing.T) {
	suite := NewMongoRepoTest()
	suite.Setup()

	gotime.Freeze(time.Date(1999, 1, 1, 1, 1, 1, 1, time.UTC))
	err := suite.repo.SetExpireUrl([]string{"test01"})
	require.NoError(t, err)

	setExpire := new(ShortUrlModel)
	err = suite.mongo.
		Database(suite.databaseName).
		Collection(suite.collectionName).
		FindOne(context.Background(), bson.D{{"_id", "test01"}}).
		Decode(setExpire)
	require.NoError(t, err)

	expect := int64(915152461)
	require.Equal(t, expect, setExpire.ExpireAt)
}
