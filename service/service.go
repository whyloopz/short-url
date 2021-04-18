package service

type Service interface {
	CreateShortUrl(url string, expireTime int64) (string, error)
	GetShortUrl(shortUrl string) (string, error)
	AdminGetShortUrls(token, searchShortUrl, searchUrl string) ([]string, error)
	AdminDeleteShortUrls(token string, shortUrl []string) error
}

type ShortUrl struct{}

func New() Service {
	return &ShortUrl{}
}
