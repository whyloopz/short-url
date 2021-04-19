package repository

func (b Blacklist) GetBlacklists() []string {
	return b.list
}
