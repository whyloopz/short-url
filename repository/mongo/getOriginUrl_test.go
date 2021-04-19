package mongo

import (
	"github.com/nqmt/goerror"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMongo_GetOriginUrl_Success(t *testing.T) {
	suite := NewMongoRepoTest()
	suite.Setup()

	originUrl, err := suite.repo.GetOriginUrl("test00")

	expect := "http://www.google1.com"
	require.NoError(t, err)
	require.Equal(t, originUrl, expect)
}

func TestMongo_GetOriginUrl_NotFound(t *testing.T) {
	suite := NewMongoRepoTest()
	suite.Setup()

	originUrl, err := suite.repo.GetOriginUrl("test99")

	expect := ""
	require.Equal(t, ErrNotFoundShortUrl.Code, err.(*goerror.GoError).Code)
	require.Equal(t, originUrl, expect)
}

func TestMongo_GetOriginUrl_Gone(t *testing.T) {
	suite := NewMongoRepoTest()
	suite.Setup()

	originUrl, err := suite.repo.GetOriginUrl("test02")

	expect := ""
	require.Equal(t, ErrGoneShortUrlExpired.Code, err.(*goerror.GoError).Code)
	require.Equal(t, originUrl, expect)
}
