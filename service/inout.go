package service

import "time"

type CreateShortUrlInput struct {
	Url        string `json:"url" validate:"required"`
	ExpireTime int64  `json:"expireTime" validate:"gte=0,lte=365"`
}

func (i CreateShortUrlInput) GetDefaultExpireTime() int64 {
	if i.ExpireTime == 0 {
		return 30
	}
	return i.ExpireTime
}

func (i CreateShortUrlInput) GetExpireAt() int64 {
	return time.Now().AddDate(0, 0, int(i.GetDefaultExpireTime())).Unix()
}

type CreateShortUrlOutput struct {
	ShortUrl string `json:"shortUrl"`
}

type GetShortUrlOutput struct {
	OriginUrl string `json:"originUrl"`
}

type AdminGetShortUrlOutput struct {
	Urls []string `json:"urls"`
}
