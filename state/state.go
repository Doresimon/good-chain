package state

import (
	"fmt"

	"github.com/Doresimon/good-chain/chain"
	"github.com/Doresimon/good-chain/console"
	"github.com/Doresimon/good-chain/crypto/hdk"
	"github.com/Doresimon/good-chain/middleware/application"
)

// State state tree
type State struct {
	OrgMap map[string]*application.Account
	Orgs   []string
}

// NewState ...
func NewState() *State {
	s := new(State)
	s.OrgMap = make(map[string]*application.Account)
	return s
}

// HandleBody ...
func (s *State) HandleBody(body *chain.Body) {
	switch body.Type {
	case "ORG":
		switch body.Action {
		case "CREATE":
			s.CreateOrg(body.ContentBytes)
		case "UPDATE":
			s.UpdateOrg()
		}

	case "ACCOUNT":
		switch body.Action {
		case "CREATE":
			s.CreateAcc(body.ContentBytes)
		case "UPDATE":
			s.UpdateAcc()
		}
	}
}

// CreateOrg ...
func (s *State) CreateOrg(content []byte) error {
	acc := application.ParseAccountCreation(content)

	if _, exist := s.OrgMap[acc.Name]; exist {
		return fmt.Errorf("org exist")
	}

	s.OrgMap[acc.Name] = acc
	s.Orgs = append(s.Orgs, acc.Name)

	return nil
}

// UpdateOrg ...
func (s *State) UpdateOrg() {
	console.Infof("UpdateOrg")
}

// CreateAcc ...
func (s *State) CreateAcc(content []byte) error {
	acc := application.ParseAccountCreation(content)

	path := hdk.NewPath(acc.Path)

	if _, exist := s.OrgMap[path.Root]; !exist {
		return fmt.Errorf("org not exist")
	}

	orgAcc := s.OrgMap[path.Root]
	parentAcc := orgAcc.GetDeepChild(path.ParentPath())

	if _, exist := parentAcc.ChildsMap[acc.Index]; exist {
		return fmt.Errorf("acc exist")
	}

	parentAcc.ChildsMap[acc.Index] = acc
	parentAcc.ChildsList = append(parentAcc.ChildsList, acc.Index)

	console.Infof("CreateAcc, name=%s, path=%s", acc.Name, acc.Path)
	return nil
}

// UpdateAcc ...
func (s *State) UpdateAcc() {
	console.Infof("UpdateAcc")

}
