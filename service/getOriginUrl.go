package service

import (
	domainShortCode "github.com/nqmt/short-url/domain"
)

func (s ShortUrl) GetOriginUrl(shortUrl string) (string, error) {
	domain := domainShortCode.ShortCode(shortUrl)
	if err := domain.Validate(); err != nil {
		return "", err
	}

	originUrl, err := s.mongoRepo.GetOriginUrl(shortUrl)
	if err != nil {
		return "", err
	}

	if err := s.mongoRepo.IncrementHit(shortUrl); err != nil {
		return "", err
	}

	return originUrl, nil
}
