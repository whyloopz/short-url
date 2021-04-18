package repository

type Blacklist struct{}

func NewBlacklist() *Blacklist {
	return &Blacklist{}
}

func (b Blacklist) GetBlacklists() ([]string, error) {
	return []string{}, nil
}
