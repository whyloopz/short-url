package service

import (
	"github.com/go-playground/validator"
	"github.com/nqmt/goerror"
	"github.com/nqmt/short-url/repository"
	"github.com/nqmt/short-url/repository/mongo"
)

var (
	ErrBadRequestValidateInput       = goerror.DefineBadRequest("ErrBadRequestValidateInput", "input not correct format")
	ErrNotFoundShortUrlInput         = goerror.DefineNotFound("ErrBadRequestValidateInput", "short url not found")
	ErrBadRequestBlacklistUrl        = goerror.DefineBadRequest("ErrBadRequestBlacklistUrl", "url is blacklist")
	ErrUnauthorizedAdminGetShortUrls = goerror.DefineUnauthorized("ErrBadRequestBlacklistUrl", "unauthorized")
)

type Service interface {
	CreateShortUrl(input *CreateShortUrlInput) (*CreateShortUrlOutput, error)
	GetOriginUrl(shortUrl string) (string, error)
	AdminGetShortUrls(token, searchShortUrl, searchUrl string) ([]*mongo.ShortUrlModel, error)
	AdminDeleteShortUrls(token string, input *AdminDeleteShortUrlsInput) error
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
