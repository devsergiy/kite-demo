// Service number #1 The router: this service should be the first endpoint to the
// service, it should be able to route the request to the appropriate service.

package main

import (
	"fmt"
	"net/http"

	conf "demo/config"
	"demo/kitewrapper"

	"github.com/koding/kite"
)

var Db *kite.Kite

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

	// Add our handler method
	k.HandleFunc("todos", func(r *kite.Request)(interface{}, error) {
	})

	k.Run()
}

