package state

import (
	"fmt"
	"sort"

	"github.com/Doresimon/good-chain/chain"
	"github.com/Doresimon/good-chain/console"
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

	case "REQUEST":
		switch body.Action {
		case "CREATE":
			s.CreateReq(body.ContentBytes)
		case "UPDATE":
			s.UpdateReq()
		}

	case "RESPONSE":
		switch body.Action {
		case "CREATE":
			s.CreateRes(body.ContentBytes)
		case "UPDATE":
			s.UpdateRes()
		}
	}
}

// CreateOrg ...
func (s *State) CreateOrg(content []byte) error {
	acc := application.ParseAccountCreation(content)
	acc.Path = acc.Name

	if _, exist := s.OrgMap[acc.Name]; exist {
		return fmt.Errorf("org exist")
	}

	s.OrgMap[acc.Name] = acc
	s.Orgs = append(s.Orgs, acc.Name)

	return nil
}

// CreateAcc ...
func (s *State) CreateAcc(content []byte) error {
	acc := application.ParseAccountCreation(content)
	path := acc.PathX

	orgAcc, exist := s.OrgMap[path.Root]
	if !exist {
		// return fmt.Errorf("org not exist")
		panic("org not exist")
	}

	parentAcc := orgAcc.GetDeepChild(path.ParentPath())
	if parentAcc == nil {
		// return fmt.Errorf("parent account not valid")
		panic("parent account not valid")
	}

	_, exist = parentAcc.ChildsMap[acc.Index]
	if exist {
		panic("acc exist")
	}

	parentAcc.ChildsMap[acc.Index] = acc
	parentAcc.ChildsList = append(parentAcc.ChildsList, acc.Index)
	sort.Sort(uint32Slice(parentAcc.ChildsList))

	return nil
}

func (s *State) CreateReq(content []byte) error {
	console.Infof("CreateReq not implemented")
	return nil
}
func (s *State) CreateRes(content []byte) error {
	console.Infof("CreateRes not implemented")
	return nil
}

func (s *State) UpdateAcc() { console.Infof("UpdateAcc not implemented") }
func (s *State) UpdateOrg() { console.Infof("UpdateOrg not implemented") }
func (s *State) UpdateReq() { console.Infof("UpdateReq not implemented") }
func (s *State) UpdateRes() { console.Infof("UpdateRes not implemented") }

// uint32Slice is used to sort uint32 array
type uint32Slice []uint32

func (s uint32Slice) Len() int           { return len(s) }
func (s uint32Slice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s uint32Slice) Less(i, j int) bool { return s[i] < s[j] }
