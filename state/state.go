package state

import (
	"fmt"
	"strings"

	"github.com/Doresimon/good-chain/chain"
	"github.com/Doresimon/good-chain/console"
)

// State state tree
type State struct {
	orgMap map[string]*Account
	orgs   []string
}

// NewState ...
func NewState() *State {
	s := new(State)
	s.orgMap = make(map[string]*Account)
	return s
}

// HandleBody ...
func (s *State) HandleBody(body *chain.Body) {
	switch body.Type {
	case "ORG":
		switch body.Action {
		case "create":
			s.CreateOrg(body.Content)
		case "update":
			s.UpdateOrg()
		}

	case "ACCOUNT":
		switch body.Action {
		case "create":
			s.CreateAcc(body.Content)
		case "update":
			s.UpdateAcc()
		}
	}
}

// CreateOrg ...
func (s *State) CreateOrg(content *chain.Content) error {
	acc := NewAccount()
	acc.name = content.Name
	acc.path = content.Name
	acc.extra = content.Extra

	if _, exist := s.orgMap[acc.name]; exist {
		return fmt.Errorf("org exist")
	}

	s.orgMap[acc.name] = acc
	s.orgs = append(s.orgs, acc.name)

	return nil
}

// UpdateOrg ...
func (s *State) UpdateOrg() {
	console.Infof("UpdateOrg")
}

// CreateAcc ...
func (s *State) CreateAcc(content *chain.Content) error {

	acc := NewAccount()
	acc.name = content.Name
	acc.path = content.Path
	acc.index = content.Index
	acc.extra = content.Extra

	paths := strings.Split(acc.path, "/")
	if _, exist := s.orgMap[paths[0]]; !exist {
		return fmt.Errorf("org not exist")
	}

	parentAcc := s.orgMap[paths[0]]
	for i := 1; i < len(paths)-1; i++ {
		parentAcc = parentAcc.GetChild(paths[i])
	}

	if _, exist := parentAcc.childs[acc.index]; exist {
		return fmt.Errorf("acc exist")
	}

	parentAcc.childs[acc.index] = acc
	parentAcc.childsList = append(parentAcc.childsList, acc.index)

	console.Infof("CreateAcc, name=%s, path=%s", acc.name, acc.path)
	return nil
}

// UpdateAcc ...
func (s *State) UpdateAcc() {
	console.Infof("UpdateAcc")

}
