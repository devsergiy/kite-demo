package router

import (
	"fmt"
	"net/http"

	"github.com/koding/kite"
)

func LoginHandler(auth *kite.Client, w http.ResponseWriter, r *http.Request) {
	// AUTH: login and return correct JWT

	params := map[string]string{
		"user": r.FormValue("username"),
		"pass": r.FormValue("password"),
	}

	result, err := auth.Tell("login", params)
	if err != nil {
		fmt.Println("Failed to login", err)
		return
	}

	fmt.Println(result) // TODO
}
