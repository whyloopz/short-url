package service

type CreateShortUrlInput struct {
	Url        string `json:"url" validate:"required"`
	ExpireTime int    `json:"expireTime" validate:"gte=0,lte=365"`
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
