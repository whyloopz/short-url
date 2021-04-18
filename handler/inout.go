package handler

type CreateShortUrlOutput struct {
	ShortUrl string `json:"shortUrl"`
}

type GetShortUrlOutput struct {
	OriginUrl string `json:"originUrl"`
}

type AdminGetShortUrlOutput struct {
	Urls []string `json:"urls"`
}
