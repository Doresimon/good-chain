package state

// Account ...
type Account struct {
	name       string
	index      string
	path       string
	pk         string
	childs     map[string]*Account
	childsList []string
	extra      map[string]string
	tasks      []string
}

// NewAccount ...
func NewAccount() *Account {
	s := new(Account)
	s.childs = make(map[string]*Account)
	s.extra = make(map[string]string)
	return s
}

// GetChild ...
func (acc *Account) GetChild(index string) *Account {
	if _, ok := acc.childs[index]; ok {
		return acc.childs[index]
	}
	return nil
}
