package app

import "github.com/koding/kite"

func TodosHandler(auth *kite.Client, db *kite.Client, r *kite.Request) (interface{}, error) {
	// TODO: dial to auth and validate token

	// TODO: dial to DB and get items list

	return []string{"one", "two"}, nil
}
