package service

import (
	"github.com/nqmt/short-url/repository/mocks"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestShortUrl_GetOriginUrl_Success(t *testing.T) {
	mBlacklistRepo := &mocks.BlacklistRepo{}
	mMongoRepo := &mocks.MongoRepo{}

	// mock data
	mMongoRepo.On("GetOriginUrl", "123546").Return("http://www.google.com", nil)
	mMongoRepo.On("IncrementHit", "123546").Return(nil)

	sv := New(mBlacklistRepo, mMongoRepo)

	input := "123546"
	expect := "http://www.google.com"

	shortUrl, err := sv.GetOriginUrl(input)
	require.NoError(t, err)
	require.Equal(t, expect, shortUrl)
}
