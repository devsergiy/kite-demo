// Service number #3 The app: this service should expose 1 endpoint1 to return
// a small to-do action list and should only be access with a valid JWT, it should
// get the data from the DB Accessor service

package main

import (
	"fmt"

	conf "demo/config"
	"demo/kitewrapper"

	"github.com/koding/kite"
)

func main() {
	routerConfig, err := conf.NewConfig("app")
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

	dbKite, err := k.FindAndDial("db")
	if err != nil {
		fmt.Println("Failed to dial db service", err)
		// return
	}

	// Add our handler method
	k.HandleFunc("todos", func(r *kite.Request) (interface{}, error) {
		// TODO: dial to DB and get items list
		_ = dbKite

		return []string{"one", "two"}, nil
	})

	k.Run()
}
