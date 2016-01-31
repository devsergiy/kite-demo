package router

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/koding/kite"
)

func LoginHandler(auth *kite.Client, w http.ResponseWriter, r *http.Request) {
	params := map[string]string{
		"user": r.FormValue("username"),
		"pass": r.FormValue("password"),
	}

	result, err := auth.Tell("login", params)
	if err != nil {
		fmt.Println("Failed to login:", err)

		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	token, _ := result.String()
	json, _ := json.Marshal(map[string]string{"token": token})

	w.Write(json)
}
