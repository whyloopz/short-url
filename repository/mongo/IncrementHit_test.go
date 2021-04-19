package mongo

import (
	"context"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

func TestMongo_IncrementHit_Success(t *testing.T) {
	suite := NewMongoRepoTest()
	suite.Setup()

	err := suite.repo.IncrementHit("test00")
	require.NoError(t, err)
	err = suite.repo.IncrementHit("test00")
	require.NoError(t, err)
	err = suite.repo.IncrementHit("test00")
	require.NoError(t, err)

	incrementHitModel := new(ShortUrlModel)
	err = suite.mongo.
		Database(suite.databaseName).
		Collection(suite.collectionName).
		FindOne(context.Background(), bson.D{{"_id", "test00"}}).
		Decode(incrementHitModel)
	require.NoError(t, err)

	expect := 3
	require.Equal(t, expect, incrementHitModel.Hit)
}
