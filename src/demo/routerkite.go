// Service number #1 The router: this service should be the first endpoint to the
// service, it should be able to route the request to the appropriate service.

package main

import (
	"fmt"
	"net/http"

	"github.com/koding/kite"

	conf "demo/config"
	"demo/kitewrapper"
	"demo/router"
)

func main() {
	routerConfig, err := conf.NewConfig("router")
	if err != nil {
		fmt.Println("Error getting config", err)
		return
	}

	k := kitewrapper.NewKiteWrapper(routerConfig)
	err = k.RegisterToKontrol()
	if err != nil {
		fmt.Println("Failed to register", err)
		return
	}

	var appKite, authKite *kite.Client

	go func() {
		var err error

		appKite, err = k.FindAndDial("app")
		if err != nil {
			fmt.Println("Failed to dial app service", err)
			// return
		}
	}()

	go func() {
		var err error

		authKite, err = k.FindAndDial("auth")
		if err != nil {
			fmt.Println("Failed to dial auth service", err)
			// return
		}
	}()

	k.HandleHTTPFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		router.LoginHandler(authKite, w, r)
	})

	k.HandleHTTPFunc("/profile", func(w http.ResponseWriter, r *http.Request) {
		router.ProfileHandler(authKite, w, r)
	})

	k.HandleHTTPFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		router.TodosHandler(appKite, w, r)
	})

	k.Run()
}
