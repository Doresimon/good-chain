package http

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Doresimon/good-chain/chain"
	"github.com/Doresimon/good-chain/console"
	"github.com/Doresimon/good-chain/crypto/bls"
	"github.com/Doresimon/good-chain/crypto/hdk"
	"github.com/Doresimon/good-chain/crypto/rand"
	"github.com/Doresimon/good-chain/state"
)

var port = "80"
var s = new(Service)
var msaterKey *bls.PrivateKey
var masterChainCode []byte

// Service ...
type Service struct {
	port string
	cs   *chain.Service
	ss   *state.Service
}

// NewService ...
func NewService(cs *chain.Service, ss *state.Service) *Service {
	// key setup
	key := []byte("good chain key")
	seed, _ := rand.Bytes(256)
	var err error
	msaterKey, masterChainCode, err = hdk.GenerateMasterKey(key, seed)
	if err != nil {
		panic(err)
	}

	// service setup
	s.port = port
	s.cs = cs
	s.ss = ss

	// api setup
	http.HandleFunc("/create-org", newAccount)
	http.HandleFunc("/create-account", newAccount)
	http.HandleFunc("/read-org-list", readOrgList)
	http.HandleFunc("/read-account-list", readAccList)
	http.HandleFunc("/create-request", newRequest)
	http.HandleFunc("/create-response", newReponse)
	go http.ListenAndServe(":"+s.port, nil)

	// print info
	console.Title("Http Info")
	console.Infof("port = %s", s.port)
	console.Infof("msaterKey = %x", msaterKey.Bytes())
	console.Infof("masterChainCode = %x", masterChainCode)
	console.Title("---------")
	return s
}

type accountInfo struct {
	Path  string            `json:"path"`
	Extra map[string]string `json:"extra"`
}

func newAccount(w http.ResponseWriter, r *http.Request) {
	bodyBytes, _ := ioutil.ReadAll(r.Body)

	var log chain.Log

	err := json.Unmarshal(bodyBytes, &log)
	if err != nil {
		panic(err)
	}

	logBodyBytes, _ := json.Marshal(log.Body)
	hash := sha256.Sum256(logBodyBytes)
	log.Hash = hash[:]

	s.cs.AddLog(&log)

	t, _ := json.Marshal(log.Body)
	console.Infof("%s", t)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	w.Write([]byte("ok"))
}
func newRequest(w http.ResponseWriter, r *http.Request) {
	bodyBytes, _ := ioutil.ReadAll(r.Body)

	console.Infof("%s", "newRequest")
	var log chain.Log

	err := json.Unmarshal(bodyBytes, &log)
	if err != nil {
		panic(err)
	}

	logBodyBytes, _ := json.Marshal(log.Body)
	hash := sha256.Sum256(logBodyBytes)
	log.Hash = hash[:]

	s.cs.AddLog(&log)

	t, _ := json.Marshal(log.Body)
	console.Infof("%s", t)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	w.Write([]byte("ok"))
}
func newReponse(w http.ResponseWriter, r *http.Request) {
	bodyBytes, _ := ioutil.ReadAll(r.Body)

	console.Infof("%s", "newReponse")
	var log chain.Log

	err := json.Unmarshal(bodyBytes, &log)
	if err != nil {
		panic(err)
	}

	logBodyBytes, _ := json.Marshal(log.Body)
	hash := sha256.Sum256(logBodyBytes)
	log.Hash = hash[:]

	s.cs.AddLog(&log)

	t, _ := json.Marshal(log.Body)
	console.Infof("%s", t)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	w.Write([]byte("ok"))
}

func readOrgList(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("readOrgList\n")
	res := s.ss.OrgList()
	// s.ss.
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	w.Write(res)
}

func readAccList(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("readAccList\n")
	res := s.ss.AccList("fudan")
	// s.ss.
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	w.Write(res)
}
