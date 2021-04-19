package domain

import "github.com/nqmt/goerror"

var (
	ErrBadRequestValidateUrl = goerror.DefineBadRequest("BadRequestValidateUrl", "url not correct format")
	ErrBadRequestShortCode   = goerror.DefineBadRequest("ErrBadRequestShortCode", "short url correct format")
	ErrInternalServerRegex   = goerror.DefineInternalServerError("ErrInternalServerRegex", "internal server error")
)
