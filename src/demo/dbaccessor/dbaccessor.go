package dbaccessor

import (
	"demo/helpers"
	"errors"
	"fmt"

	"github.com/koding/kite"
	yaml "gopkg.in/yaml.v2"
)

const (
	FixturePath = "../fixtures/models.yml"
)

type DBAccessor struct {
	User  map[string]string
	Todos []string
}

func NewDBAccessor() *DBAccessor {
	a := &DBAccessor{}

	yamlFile, _ := helpers.GetYamlContent(FixturePath)
	err := yaml.Unmarshal(yamlFile, a)
	if err != nil {
		fmt.Println("Error parsing user fixture")
	}

	return a
}

func (a *DBAccessor) QueryHandler(r *kite.Request) (interface{}, error) {
	entity, _ := r.Args.One().String()

	if entity == "todos" {
		return a.Todos, nil
	} else if entity == "user" {
		return a.User, nil
	}

	return nil, errors.New("Bad entity")
}
