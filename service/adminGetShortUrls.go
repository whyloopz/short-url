package service

import "github.com/nqmt/short-url/repository/mongo"

func (s ShortUrl) AdminGetShortUrls(token, searchShortUrl, searchUrl string) ([]*mongo.ShortUrlModel, error) {
	if token != s.adminToken {
		return nil, ErrUnauthorizedAdminGetShortUrls
	}

	return s.mongoRepo.GetShortUrls(searchShortUrl, searchUrl)
}
