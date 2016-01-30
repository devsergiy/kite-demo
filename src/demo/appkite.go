// Service number #3 The app: this service should expose 1 endpoint1 to return
// a small to-do action list and should only be access with a valid JWT, it should
// get the data from the DB Accessor service

package main

import (
	"fmt"

	"demo/app"
	conf "demo/config"
	"demo/kitewrapper"

	"github.com/koding/kite"
)

func main() {
	appConfig, err := conf.NewConfig("app")
	if err != nil {
		fmt.Println("Error getting config", err)
		return
	}

	k := kitewrapper.NewKiteWrapper(appConfig)
	err = k.RegisterToKontrol()
	if err != nil {
		fmt.Println("Failed to register", err)
		return
	}

	authKite, err := k.FindAndDial("auth")
	if err != nil {
		fmt.Println("Failed to dial auth service", err)
		// return
	}

	dbKite, err := k.FindAndDial("db")
	if err != nil {
		fmt.Println("Failed to dial db service", err)
		// return
	}

	// Add our handler method
	k.HandleFunc("todos", func(r *kite.Request) (interface{}, error) {
		return app.TodosHandler(authKite, dbKite, r)
	})

	k.Run()
}
