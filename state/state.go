package state

import (
	"fmt"
	"sort"

	"github.com/Doresimon/good-chain/chain"
	"github.com/Doresimon/good-chain/console"
	"github.com/Doresimon/good-chain/crypto/hdk"
	"github.com/Doresimon/good-chain/middleware/application"
)

// State state tree
type State struct {
	OrgMap      map[string]*application.Account  // Org.name => *Account
	RequestMap  map[string]*application.Request  // Request.hash.hex => *Request
	ResponseMap map[string]*application.Response // Response.hash.hex => *Response
	Orgs        []string                         // Org.name
}

// NewState ...
func NewState() *State {
	s := new(State)
	s.OrgMap = make(map[string]*application.Account)
	s.RequestMap = make(map[string]*application.Request)
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
			s.CreateReq(body.HashHexString(), body.ContentBytes)
		case "UPDATE":
			s.UpdateReq()
		}

	case "RESPONSE":
		switch body.Action {
		case "CREATE":
			s.CreateRes(body.HashHexString(), body.ContentBytes)
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

// CreateReq ...
func (s *State) CreateReq(hashHex string, content []byte) error {
	console.Infof("CreateReq ")
	req := application.ParseRequestCreation(content)
	path := hdk.NewPath(req.RequesterPath)
	orgAcc, exist := s.OrgMap[path.Root]
	if !exist {
		return fmt.Errorf("org not exist")
	}
	req.Requester = orgAcc.GetDeepChild(path.Indexes)
	req.TXHash = hashHex
	parent := orgAcc.GetDeepChild(path.ParentPath())
	parent.RequestList = append(parent.RequestList, req)

	s.RequestMap[req.TXHash] = req
	return nil
}

// CreateRes ...
func (s *State) CreateRes(hashHex string, content []byte) error {
	console.Infof("CreateRes")
	res := application.ParseResponseCreation(content)
	path := hdk.NewPath(res.ResponserPath)
	orgAcc, exist := s.OrgMap[path.Root]
	if !exist {
		return fmt.Errorf("org not exist")
	}
	res.Responser = orgAcc.GetDeepChild(path.Indexes)
	res.TXHash = hashHex

	requester := s.RequestMap[res.TXHash].Requester
	requester.ResponseList = append(requester.ResponseList, res)

	s.ResponseMap[res.TXHash] = res
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
