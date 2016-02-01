package app

import (
	"errors"

	"github.com/koding/kite"
)

var BadToken = errors.New("Bad Token")

func TodosHandler(auth *kite.Client, db *kite.Client, r *kite.Request) (interface{}, error) {
	token, err := r.Args.One().String()
	if err != nil {
		return nil, BadToken
	}

	result, _ := auth.Tell("validateToken", token)
	if !result.MustBool() {
		return nil, BadToken
	}

	result, _ = db.Tell("query", "todos")

	var (
		arr, _ = result.Slice()
		todos  []string
	)
	for _, a := range arr {
		val, _ := a.String()
		todos = append(todos, val)
	}

	return todos, nil
}
