package mongo

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMongo_GetShortUrls(t *testing.T) {
	suite := NewMongoRepoTest()
	suite.Setup()

	shortUrls, err := suite.repo.GetShortUrls("", "")

	expect := 3
	require.NoError(t, err)
	require.Equal(t, expect, len(shortUrls))
}
