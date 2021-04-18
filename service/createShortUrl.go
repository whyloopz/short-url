package service

import (
	"github.com/nqmt/short-url/domain"
)

func (s ShortUrl) CreateShortUrl(input *CreateShortUrlInput) (*CreateShortUrlOutput, error) {
	if err := s.validate.Struct(input); err != nil {
		return nil, ErrBadRequestValidateInput.WithCause(err)
	}

	domainShortUrl := domain.NewShortUrl(input.Url, input.ExpireTime)
	if err := domainShortUrl.ValidateUrl(); err != nil {
		return nil, err
	}

	blacklists, err := s.blacklistRepo.GetBlacklists()
	if err != nil {
		return nil, err
	}
	if domainShortUrl.IsBlackList(blacklists) {
		return nil, ErrBadRequestBlacklistUrl
	}

	shortUrl := domainShortUrl.GenShortUrl()
	if err := s.mongoRepo.SaveShortUrl(shortUrl, input.Url, domainShortUrl.GetExpireAt()); err != nil {
		return nil, err
	}

	return &CreateShortUrlOutput{
		ShortUrl: shortUrl,
	}, nil
}
