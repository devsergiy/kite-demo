package router

import (
	"fmt"
	"net/http"

	"github.com/koding/kite"
)

func TodosHandler(app *kite.Client, w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("token")

	result, err := app.Tell("todos", token)
	if err != nil {
		fmt.Println("Failed to get todos", err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.Write(result.Raw)
}
