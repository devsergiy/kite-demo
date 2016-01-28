// Service number #1 The router: this service should be the first endpoint to the
// service, it should be able to route the request to the appropriate service.

package main

import (
	"fmt"
	"net/http"

	conf "demo/config"
	"demo/kitewrapper"
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

	appKite, err := k.FindAndDial("app")
	if err != nil {
		fmt.Println("Failed to dial app service", err)
		return
	}

	authKite, err := k.FindAndDial("auth")
	if err != nil {
		fmt.Println("Failed to dial auth service", err)
		return
	}

	k.HandleHTTPFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		// AUTH: login and return correct JWT

		params := map[string]string{
			"user": r.FormValue("username"),
			"pass": r.FormValue("password"),
		}

		result, err := authKite.Tell("login", params)
		if err != nil {
			fmt.Println("Failed to login", err)
			return
		}

		fmt.Println(result) // TODO
	})

	k.HandleHTTPFunc("/profile", func(w http.ResponseWriter, r *http.Request) {
		result, err := authKite.Tell("profile")
		if err != nil {
			fmt.Println("Failed to get profile", err)
			return
		}

		fmt.Println(result) // TODO
	})

	k.HandleHTTPFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		result, err := appKite.Tell("todos")
		if err != nil {
			fmt.Println("Failed to get todos", err)
			return
		}

		fmt.Println(result) // TODO
	})

	k.Run()
}
