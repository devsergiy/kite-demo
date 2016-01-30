// Service number # DB Accessor: this service should expose db content, there is
// no need to use a real database hardcoded data or read from a json|yaml file
// is ok.

package main

import (
	"fmt"

	conf "demo/config"
	"demo/dbaccessor"
	"demo/kitewrapper"

	"github.com/koding/kite"
)

func main() {
	dbConfig, err := conf.NewConfig("db")
	if err != nil {
		fmt.Println("Error getting config", err)
		return
	}

	k := kitewrapper.NewKiteWrapper(dbConfig)
	err = k.RegisterToKontrol()
	if err != nil {
		fmt.Println("Failed to register", err)
		return
	}

	db := dbaccessor.NewDBAccessor()

	// Add our handler method
	k.HandleFunc("query", func(r *kite.Request) (interface{}, error) {
		return db.QueryHandler(r)
	})

	k.Run()
}
