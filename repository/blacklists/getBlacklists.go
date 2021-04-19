package blacklists

func (b Blacklist) GetBlacklists() []string {
	return b.list
}
