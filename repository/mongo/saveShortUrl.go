package mongo

import (
	"context"
	"github.com/nqmt/gotime/v2"
	"time"
)

func (m Mongo) SaveShortUrl(shortUrl, originUrl string, expireAt int64) error {
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(m.timeout)*time.Second)

	model := &ShortUrlModel{
		ID:        shortUrl,
		Url:       originUrl,
		Hit:       0,
		CreatedAt: gotime.NowUnix(),
		ExpireAt:  expireAt,
	}

	_, err := m.shortUrlCollection.InsertOne(ctx, model)
	if err != nil {
		return ErrInternalServerSaveShortUrlMongoDB.WithCause(err)
	}

	return nil
}
