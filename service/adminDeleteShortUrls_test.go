package service

import (
	"github.com/nqmt/short-url/repository/mocks"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestShortUrl_AdminDeleteShortUrls(t *testing.T) {
	mBlacklistRepo := &mocks.BlacklistRepo{}
	mMongoRepo := &mocks.MongoRepo{}

	// mock data
	mMongoRepo.On("SetExpireUrl", []string{"test001"}).Return(nil)

	sv := New(mBlacklistRepo, mMongoRepo, "root")
	err := sv.AdminDeleteShortUrls("root", &AdminDeleteShortUrlsInput{ShortUrl: []string{"test001"}})
	require.NoError(t, err)
}
