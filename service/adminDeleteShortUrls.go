package service

func (s ShortUrl) AdminDeleteShortUrls(token string, input *AdminDeleteShortUrlsInput) error {
	if token != s.adminToken {
		return ErrUnauthorizedAdminGetShortUrls
	}

	if len(input.ShortUrl) < 1 {
		return ErrNotFoundShortUrlInput
	}

	return s.mongoRepo.SetExpireUrl(input.ShortUrl)
}
