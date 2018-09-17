package chain

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
)

type ChainConfig struct {
	UID  *big.Int `json:"uid"`
	Name string   `json:"name"`
}

var defaultConfig *ChainConfig
var defaultPath string

func init() {
	defaultConfig = new(ChainConfig)
	defaultConfig.Name = "default chain -. -"
	defaultConfig.UID = big.NewInt(0)

	defaultPath = "./chain.config"
}

func (this *ChainConfig) read(path string) {
	fmt.Println("ChainConfig.read(path string)")

	dat, err := ioutil.ReadFile(path)
	config := new(ChainConfig)

	if err != nil {
		data, _ := json.Marshal(defaultConfig)
		config = defaultConfig

		fmt.Printf("Create new %s file \n", path)
		ioutil.WriteFile(path, data, 0777)
	} else {
		fmt.Printf("Read from %s file \n", path)
		_ = json.Unmarshal(dat, config)
	}

	this.Name = config.Name
	this.UID = config.UID
}

func (this *ChainConfig) readDefault() {
	fmt.Println("ChainConfig.readDefault()")

	dat, err := ioutil.ReadFile(defaultPath)
	config := new(ChainConfig)

	if err != nil {
		data, _ := json.Marshal(defaultConfig)
		config = defaultConfig

		fmt.Println("Create new chain.config file")

		ioutil.WriteFile(defaultPath, data, 0777)
	} else {
		fmt.Println("Read from chain.config file")
		_ = json.Unmarshal(dat, config)
	}

	this.Name = config.Name
	this.UID = config.UID
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
