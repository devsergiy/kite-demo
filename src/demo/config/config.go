package config

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	yaml "gopkg.in/yaml.v2"
)

type KiteConfig struct {
	Name    string
	Port    int
	Host    string
	Path    string
	Version string
}

type Config struct {
	Kontrol struct {
		URL  string
		User string
	}
	Router KiteConfig
	App    KiteConfig
	DB     KiteConfig
	Auth   KiteConfig
}

const CONFIG_PATH = "../../conf/config.yml"

func NewConfig() (*Config, error) {
	config := &Config{}

	filename, _ := filepath.Abs(CONFIG_PATH)
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("No such file %s", filename)
		return nil, err
	}

	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		fmt.Println("Error parsing config")
		return nil, err
	}

	return config, nil
}
