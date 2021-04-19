package service

import (
	"github.com/nqmt/gotime/v2"
	"github.com/nqmt/short-url/repository/mocks"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestShortUrl_CreateShortUrl_Success(t *testing.T) {
	mBlacklistRepo := &mocks.BlacklistRepo{}
	mMongoRepo := &mocks.MongoRepo{}

	// mock data
	mBlacklistRepo.On("GetBlacklists").Return([]string{}, nil)
	mMongoRepo.On("SaveShortUrl", "c91cfc", "http://www.test.com", int64(1588565044)).Return(nil)
	gotime.Freeze(time.Date(2020, 4, 4, 4, 4, 4, 4, time.UTC))

	sv := New(mBlacklistRepo, mMongoRepo, "")

	// input & output expect
	input := &CreateShortUrlInput{
		Url:        "http://www.test.com",
		ExpireTime: 30,
	}
	expect := &CreateShortUrlOutput{ShortUrl: "c91cfc"}

	shortUrl, err := sv.CreateShortUrl(input)
	require.NoError(t, err)
	require.Equal(t, expect, shortUrl)
}
