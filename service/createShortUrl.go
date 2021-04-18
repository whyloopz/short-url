package service

func (s ShortUrl) CreateShortUrl(url string, _ int64) (string, error) {
	return url, nil
}
