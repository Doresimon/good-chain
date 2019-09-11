package application

import (
	"encoding/json"
	"fmt"

	"github.com/Doresimon/good-chain/crypto/hdk"
)

// Account ...
type Account struct {
	Name       string
	Index      uint32
	Path       string
	PathX      *hdk.Path
	Pk         string
	ChildsMap  map[uint32]*Account
	ChildsList []uint32
	Extra      string
	Tasks      []string
}

// NewAccount ...
func NewAccount() *Account {
	s := new(Account)
	s.ChildsMap = make(map[uint32]*Account)
	return s
}

// GetChild ...
func (acc *Account) GetChild(index uint32) *Account {
	if _, ok := acc.ChildsMap[index]; ok {
		return acc.ChildsMap[index]
	}
	return nil
}

// GetDeepChild ...
func (acc *Account) GetDeepChild(indexes []uint32) *Account {
	fmt.Printf("%%GetDeepChild: %d\n", indexes)
	var child = acc
	for _, index := range indexes {
		_, ok := child.ChildsMap[index]
		if !ok {
			return nil
		}
		child = child.ChildsMap[index]
	}
	return child
}

// String ...
func (acc *Account) String() string {
	var str = fmt.Sprintf("%s: %s index=%d", acc.Path, acc.Name, acc.Index)
	for _, index := range acc.ChildsList {
		_, ok := acc.ChildsMap[index]
		if !ok {
			continue
		}
		str += "\n" + acc.ChildsMap[index].String()
	}
	return str
}

// AccountCreation ...
type AccountCreation struct {
	Name  string `json:"name"`
	Path  string `json:"path"`
	Extra string `json:"extra"`
}

// ParseAccountCreation ...
func ParseAccountCreation(contentBytes []byte) *Account {
	ac := new(AccountCreation)
	err := json.Unmarshal(contentBytes, ac)
	if err != nil {
		panic(err)
	}

	acc := NewAccount()
	acc.Path = ac.Path
	acc.PathX = hdk.NewPath(acc.Path)
	acc.Name = ac.Name
	acc.Extra = ac.Extra
	acc.Index = acc.PathX.LastIndex()

	return acc
}
