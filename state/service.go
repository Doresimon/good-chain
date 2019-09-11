package state

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/Doresimon/good-chain/chain"
	"github.com/Doresimon/good-chain/console"
)

var stateTime = time.Second * 2

// Service state service
type Service struct {
	next  uint64
	cs    *chain.Service
	state *State
}

// NewService creates a new state service
func NewService(cs *chain.Service) *Service {
	s := new(Service)
	s.next = 0
	s.cs = cs
	s.state = NewState()

	go s.Tick()

	return s
}

// Update ...
func (s *Service) Update() {
	curBN := s.cs.C.BN()

	for i := s.next; i < curBN; i++ {
		console.Infof("state.Service: state tree follows #%x block", i)

		block := s.cs.C.ReadBlock(i)

		for _, log := range block.Logs {
			s.state.HandleBody(log.Body)
		}

		for _, orgName := range s.state.Orgs {
			org := s.state.OrgMap[orgName]
			fmt.Printf("name=%s, extra=%s\n", org.Name, org.Extra)
			fmt.Printf("%s\n", org)
		}
	}
	s.next = curBN
}

// Tick ...
func (s *Service) Tick() {
	tiktok := time.NewTicker(stateTime)

	for {
		select {
		case <-tiktok.C:
			s.Update()
		}
	}
}

// OrgList ...
func (s *Service) OrgList() []byte {
	ret := make([]map[string]interface{}, len(s.state.Orgs), len(s.state.Orgs))

	fmt.Printf(" len(s.state.orgs)=%d\n", len(s.state.Orgs))

	for i, orgName := range s.state.Orgs {
		org := s.state.OrgMap[orgName]
		retOrg := make(map[string]interface{})
		retOrg["name"] = org.Name
		retOrg["child_len"] = len(org.ChildsList)
		retOrg["task_len"] = len(org.Tasks)

		ret[i] = retOrg
	}

	retBytes, _ := json.Marshal(ret)

	fmt.Printf("%s\n", retBytes)

	return retBytes
}

// AccList ...
func (s *Service) AccList(orgName string) []byte {
	org := s.state.OrgMap[orgName]
	ret := make([]map[string]interface{}, len(org.ChildsList), len(org.ChildsList))

	for i, childName := range org.ChildsList {
		child := org.ChildsMap[childName]
		tmp := make(map[string]interface{})
		tmp["name"] = child.Name
		tmp["path"] = child.Path
		tmp["index"] = child.Index
		tmp["extra"] = child.Extra

		ret[i] = tmp
	}

	retBytes, _ := json.Marshal(ret)

	fmt.Printf("%s\n", retBytes)

	return retBytes
}
