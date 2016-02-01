package router

import (
	"fmt"
	"net/http"

	"github.com/koding/kite"
)

func ProfileHandler(auth *kite.Client, w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("token")

	result, err := auth.Tell("profile", token)
	if err != nil {
		fmt.Println("Failed to get profile", err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.Write(result.Raw)
}
