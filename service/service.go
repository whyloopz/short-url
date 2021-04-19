package service

import (
	"github.com/go-playground/validator"
	"github.com/nqmt/goerror"
	"github.com/nqmt/short-url/repository"
)

var (
	ErrBadRequestValidateInput = goerror.DefineBadRequest("ErrBadRequestValidateInput", "input not correct format")
	ErrBadRequestBlacklistUrl  = goerror.DefineBadRequest("ErrBadRequestBlacklistUrl", "url is blacklist")
)

type Service interface {
	CreateShortUrl(input *CreateShortUrlInput) (*CreateShortUrlOutput, error)
	GetOriginUrl(shortUrl string) (string, error)
	AdminGetShortUrls(token, searchShortUrl, searchUrl string) ([]string, error)
	AdminDeleteShortUrls(token string, shortUrl []string) error
}

type ShortUrl struct {
	validate      *validator.Validate
	blacklistRepo repository.BlacklistRepo
	mongoRepo     repository.MongoRepo
	adminToken    string
}

func New(blacklistRepo repository.BlacklistRepo, MongoRepo repository.MongoRepo, adminToken string) Service {
	return &ShortUrl{
		validate:      validator.New(),
		blacklistRepo: blacklistRepo,
		mongoRepo:     MongoRepo,
		adminToken:    adminToken,
	}
}
