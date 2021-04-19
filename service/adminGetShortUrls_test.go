package service

import (
	"github.com/nqmt/short-url/repository/mocks"
	"github.com/nqmt/short-url/repository/mongo"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestShortUrl_AdminGetShortUrls(t *testing.T) {
	mBlacklistRepo := &mocks.BlacklistRepo{}
	mMongoRepo := &mocks.MongoRepo{}

	// mock data
	mockData := []*mongo.ShortUrlModel{
		{"1", "url", 0, 0, 0},
		{"1", "url", 0, 0, 0},
	}
	mMongoRepo.On("GetShortUrls", "test01", "https://").Return(mockData, nil)

	sv := New(mBlacklistRepo, mMongoRepo, "root")
	got, err := sv.AdminGetShortUrls("root", "test01", "https://")

	expect := 2
	require.NoError(t, err)
	require.Equal(t, expect, len(got))
}
