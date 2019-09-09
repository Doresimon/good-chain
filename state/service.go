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
	ret := make([]map[string]interface{}, len(s.state.orgs), len(s.state.orgs))

	fmt.Printf(" len(s.state.orgs)=%d\n", len(s.state.orgs))

	for i, orgName := range s.state.orgs {
		org := s.state.orgMap[orgName]
		retOrg := make(map[string]interface{})
		retOrg["name"] = org.name
		retOrg["child_len"] = len(org.childsList)
		retOrg["task_len"] = len(org.tasks)

		ret[i] = retOrg
	}

	retBytes, _ := json.Marshal(ret)

	fmt.Printf("%s\n", retBytes)

	return retBytes
}

// AccList ...
func (s *Service) AccList(orgName string) []byte {
	org := s.state.orgMap[orgName]
	ret := make([]map[string]interface{}, len(org.childsList), len(org.childsList))

	for i, childName := range org.childsList {
		child := org.childs[childName]
		tmp := make(map[string]interface{})
		tmp["name"] = child.name
		tmp["path"] = child.path
		tmp["index"] = child.index
		tmp["extra"] = child.extra

		ret[i] = tmp
	}

	retBytes, _ := json.Marshal(ret)

	fmt.Printf("%s\n", retBytes)

	return retBytes
}
