package config

import (
	"github.com/BurntSushi/toml"
	"log"
	"path/filepath"
	"os"
	"path"
)

const configFile = "config.toml"

var configuration Config

func getConfigFilePath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return path.Join(dir, configFile)
}

func LoadConfiguration() {
	if _, err := toml.DecodeFile(getConfigFilePath(), &configuration); err != nil {
		log.Fatalf("%v", err)
	}
}

func GetConfiguration() Config {
	return configuration
}
