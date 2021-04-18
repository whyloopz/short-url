package handler

type CreateShortUrlInput struct {
	Url        string `json:"url" validate:"required"`
	ExpireTime int64  `json:"expireTime"`
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
