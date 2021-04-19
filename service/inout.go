package service

import "github.com/nqmt/short-url/repository/mongo"

type CreateShortUrlInput struct {
	Url        string `json:"url" validate:"required"`
	ExpireTime int    `json:"expireTime" validate:"gte=0,lte=365"`
}

type CreateShortUrlOutput struct {
	ShortUrl string `json:"shortUrl"`
}

func (o *CreateShortUrlOutput) SetHostName(hostname string) {
	o.ShortUrl = hostname + "/" + o.ShortUrl
}

type GetShortUrlOutput struct {
	OriginUrl string `json:"originUrl"`
}

type AdminGetShortUrlOutput struct {
	Urls []*mongo.ShortUrlModel `json:"urls"`
}

type AdminDeleteShortUrlsInput struct {
	ShortUrl []string `json:"shortUrl" validate:"required,gt=0"`
}
