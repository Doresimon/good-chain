package main

import (
	"encoding/json"
	"time"

	"github.com/Doresimon/good-chain/chain"
	"github.com/Doresimon/good-chain/crypto/hash"
	"github.com/Doresimon/good-chain/crypto/hdk"
	"github.com/Doresimon/good-chain/middleware/application"
)

func newOrg() []byte {
	var err error

	content := new(application.OrgCreation)
	content.Name = "fudan"
	content.Extra = "复旦大学, 中国上海, 邯郸路220号"
	// content.Name = "tongji"
	// content.Extra = "同济大学, 中国上海, 地址不详"
	content.PublicKey = string(masterPrivKey.Public().Bytes())
	content.ChainCode = string(masterChainCode)

	body := new(chain.Body)
	body.Type = "ORG"
	body.Action = "CREATE"
	body.Timestamp = uint32(time.Now().Unix())
	body.ContentBytes, err = json.Marshal(content)
	if err != nil {
		panic(err)
	}

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}
	sigBytes, err := masterPrivKey.Sign(bodyBytes)
	if err != nil {
		panic(err)
	}

	log := chain.NewLog(masterPrivKey.Public().Bytes(), sigBytes, body)

	logBytes, err := log.Marshal()
	if err != nil {
		panic(err)
	}

	return logBytes
}

func newAccount() []byte {
	var err error

	content := new(application.AccountCreation)
	content.Name = "陈老师"
	content.Path = "fudan/0"
	content.Extra = "美美哒"

	body := new(chain.Body)
	body.Type = "ACCOUNT"
	body.Action = "CREATE"
	body.Timestamp = uint32(time.Now().Unix())
	body.ContentBytes, err = json.Marshal(content)
	if err != nil {
		panic(err)
	}

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}
	sigBytes, err := masterPrivKey.Sign(bodyBytes)
	if err != nil {
		panic(err)
	}

	log := chain.NewLog(masterPrivKey.Public().Bytes(), sigBytes, body)

	logBytes, err := log.Marshal()
	if err != nil {
		panic(err)
	}

	return logBytes
}

func newRequest() []byte {
	var err error
	path := "fudan/0"
	data := "申请打印成绩单"
	sk, cc, ok := hdk.Priv2Priv(masterPrivKey, masterChainCode, 0)
	_ = cc
	if !ok {
		panic("key generate failed")
	}

	content := new(application.RequestCreation)
	content.Signer = path
	content.Data = data
	content.Commit = hash.SHA256Hex([]byte(data))
	content.ECIESKey = ""

	body := new(chain.Body)
	body.Type = "REQUEST"
	body.Action = "CREATE"
	body.Timestamp = uint32(time.Now().Unix())
	body.ContentBytes, err = json.Marshal(content)
	if err != nil {
		panic(err)
	}

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}
	sigBytes, err := sk.Sign(bodyBytes)
	if err != nil {
		panic(err)
	}

	log := chain.NewLog(sk.Public().Bytes(), sigBytes, body)

	logBytes, err := log.Marshal()
	if err != nil {
		panic(err)
	}

	return logBytes
}
