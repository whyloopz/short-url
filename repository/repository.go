package repository

type BlacklistRepo interface {
	GetBlacklists() ([]string, error)
}

type MongoRepo interface {
	SaveShortUrl(shortUrl, originUrl string, expireAt int64) error
}
