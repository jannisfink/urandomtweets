package config

import (
	"github.com/BurntSushi/toml"
	"log"
)

const configFile = "config.toml"

var configuration Config

func LoadConfiguration() {
	if _, err := toml.DecodeFile(configFile, &configuration); err != nil {
		log.Fatalf("%v", err)
	}
}

func GetConfiguration() Config {
	return configuration
}
