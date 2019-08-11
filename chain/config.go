package chain

import (
	"encoding/json"
	"io/ioutil"
	"math/big"
	"time"

	"github.com/Doresimon/good-chain/console"
)

var blockTime = time.Second * 10
var logPoolSize = 1024
var blockPoolSize = 1024

// Config ...
type Config struct {
	Version uint16   `json:"version"`
	UID     *big.Int `json:"uid"`
	Name    string   `json:"name"`
}

var defaultConfig = Config{
	Version: 0x1000,
	UID:     big.NewInt(0),
	Name:    "default chain -. -",
}
var defaultPath = "./chain.json"

func init() {
	// defaultConfig.Name = "default chain -. -"
	// defaultConfig.UID = big.NewInt(0)
	// defaultConfig.Version = 0x1000

	// defaultPath = "./chain.json"
}

func (cfg *Config) read(path string) {
	console.Dev("Config.read(" + path + ")")

	buf, err := ioutil.ReadFile(path)
	var config Config

	if err != nil {
		data, _ := json.MarshalIndent(defaultConfig, "", "\t")
		config = defaultConfig

		console.Info("Create new " + path + " file")
		ioutil.WriteFile(path, data, 0777)
	} else {

		console.Info("Read from " + path + " file")
		_ = json.Unmarshal(buf, &config)
	}

	cfg.Name = config.Name
	cfg.UID = config.UID
}

func (cfg *Config) readDefault() {
	console.Dev("Config.readDefault()")

	dat, err := ioutil.ReadFile(defaultPath)
	var config Config

	if err != nil {
		data, _ := json.MarshalIndent(defaultConfig, "", "\t")
		config = defaultConfig

		console.Info("Create new chain.config file")

		ioutil.WriteFile(defaultPath, data, 0777)
	} else {
		console.Info("Read from chain.config file")
		_ = json.Unmarshal(dat, &config)
	}

	cfg.Name = config.Name
	cfg.UID = config.UID
}
