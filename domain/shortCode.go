package domain

import "regexp"

type ShortCode string

func (s ShortCode) Validate() error {
	match, err := regexp.MatchString(`^([a-fA-F0-9]){6}$`, string(s))
	if err != nil {
		return ErrInternalServerRegex.WithCause(err)
	}

	if !match {
		return ErrBadRequestShortCode
	}

	return nil
}
