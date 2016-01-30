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

type config struct {
	Kites map[string]KiteConfig
}

const CONFIG_PATH = "../conf/config.yml"

func NewConfig(kiteName string) (*KiteConfig, error) {
	config := &config{}

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

	kiteConf, ok := config.Kites[kiteName]
	if !ok {
		fmt.Println("No such kite config", kiteName)
		return nil, err
	}

	return &kiteConf, nil
}
