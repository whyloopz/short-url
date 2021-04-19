package blacklists

import "strings"

type Blacklist struct {
	list []string
}

func NewBlacklistRepo(blacklists string) *Blacklist {
	return &Blacklist{
		list: strings.Split(blacklists, ","),
	}
}
