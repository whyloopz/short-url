package domain

import (
	"crypto/md5"
	"fmt"
	"github.com/nqmt/goerror"
	"github.com/nqmt/gotime/v2"
	"regexp"
	"strconv"
)

var (
	ErrBadRequestValidateUrl = goerror.DefineBadRequest("BadRequestValidateUrl", "url not correct format")
	ErrInternalServerRegex   = goerror.DefineInternalServerError("ErrInternalServerRegex", "internal server error")
)

type ShortUrl struct {
	originUrl  string
	expireTime int
}

func NewShortUrl(OriginUrl string, expireTime int) *ShortUrl {
	return &ShortUrl{originUrl: OriginUrl, expireTime: expireTime}
}

func (s ShortUrl) ValidateUrl() error {
	match, err := regexp.MatchString(`https?://(www\.)?[-a-zA-Z0-9@:%._+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_+.~#?&/=]*)`, s.originUrl)
	if err != nil {
		return ErrInternalServerRegex.WithCause(err)
	}

	if !match {
		return ErrBadRequestValidateUrl
	}

	return nil
}

func (s ShortUrl) IsBlackList(blacklists []string) bool {
	for _, blacklist := range blacklists {
		if blacklist == s.originUrl {
			return true
		}
	}

	return false
}

func (s ShortUrl) GenShortUrl() string {
	composeUrl := strconv.FormatInt(gotime.Now().UnixNano(), 10) + "_" + s.originUrl
	encodeMd5 := md5.Sum([]byte(composeUrl))

	return fmt.Sprintf("%x", encodeMd5)[:6]
}

func (s ShortUrl) GetDefaultExpireAt() int {
	if s.expireTime == 0 {
		return 30
	}

	return s.expireTime
}

func (s ShortUrl) GetExpireAt() int64 {
	return gotime.
		Now().
		AddDate(0, 0, s.GetDefaultExpireAt()).
		Unix()
}
