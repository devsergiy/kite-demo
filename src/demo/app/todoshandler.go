package app

import "github.com/koding/kite"

func TodosHandler(auth *kite.Client, db *kite.Client, r *kite.Request) (interface{}, error) {
	// TODO: dial to auth and validate token

	// TODO: dial to DB and get items list
	result, _ := db.Tell("query", "todos")

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
