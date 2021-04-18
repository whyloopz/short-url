package service

import (
	"github.com/nqmt/short-url/domain"
)

func (s ShortUrl) CreateShortUrl(input *CreateShortUrlInput) (*CreateShortUrlOutput, error) {
	if err := s.validate.Struct(input); err != nil {
		return nil, ErrBadRequestValidateInput.WithCause(err)
	}

	shortUrlDomain := domain.NewShortUrl(input.Url)
	if err := shortUrlDomain.ValidateUrl(); err != nil {
		return nil, err
	}

	blacklists, err := s.blacklistRepo.GetBlacklists()
	if err != nil {
		return nil, err
	}
	if shortUrlDomain.IsBlackList(blacklists) {
		return nil, ErrBadRequestValidateInput
	}

	shortUrl := shortUrlDomain.GenShortUrl()
	if err := s.mongoRepo.SaveShortUrl(shortUrl, input.Url, input.GetExpireAt()); err != nil {
		return nil, ErrBadRequestValidateInput
	}

	return &CreateShortUrlOutput{
		ShortUrl: shortUrl,
	}, nil
}
