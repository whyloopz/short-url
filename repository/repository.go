package repository

//go:generate mockery --name=BlacklistRepo
type BlacklistRepo interface {
	GetBlacklists() []string
}

//go:generate mockery --name=MongoRepo
type MongoRepo interface {
	SaveShortUrl(shortUrl, originUrl string, expireAt int64) error
	GetOriginUrl(shortUrl string) (string, error)
	IncrementHit(shortUrl string) error
}
