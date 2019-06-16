package chain

import (
	"encoding/json"
	"io/ioutil"
	"math/big"

	"github.com/Doresimon/good-chain/console"
)

// Config ...
type Config struct {
	Version uint16   `json:"version"`
	UID     *big.Int `json:"uid"`
	Name    string   `json:"name"`
}

var defaultConfig *Config
var defaultPath string

func init() {
	defaultConfig = new(Config)
	defaultConfig.Name = "default chain -. -"
	defaultConfig.UID = big.NewInt(0)
	defaultConfig.Version = 0x1000

	defaultPath = "./chain.config"
}

func (this *Config) read(path string) {
	console.Dev("Config.read(" + path + ")")

	dat, err := ioutil.ReadFile(path)
	config := new(Config)

	if err != nil {
		data, _ := json.MarshalIndent(defaultConfig, "", "\t")
		config = defaultConfig

		console.Info("Create new " + path + " file")
		ioutil.WriteFile(path, data, 0777)
	} else {

		console.Info("Read from " + path + " file")
		_ = json.Unmarshal(dat, config)
	}

	this.Name = config.Name
	this.UID = config.UID
}

func (this *Config) readDefault() {
	console.Dev("Config.readDefault()")

	dat, err := ioutil.ReadFile(defaultPath)
	config := new(Config)

	if err != nil {
		data, _ := json.MarshalIndent(defaultConfig, "", "\t")
		config = defaultConfig

		console.Info("Create new chain.config file")

		ioutil.WriteFile(defaultPath, data, 0777)
	} else {
		console.Info("Read from chain.config file")
		_ = json.Unmarshal(dat, config)
	}

	this.Name = config.Name
	this.UID = config.UID
}
