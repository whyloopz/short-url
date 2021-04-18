package repository

type Mongo struct{}

func NewMongo() *Mongo {
	return &Mongo{}
}

func (m Mongo) SaveShortUrl(shortUrl, originUrl string, expireAt int64) error {
	return nil
}
