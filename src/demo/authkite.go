// Service number #3 The app: this service should expose 1 endpoint1 to return
// a small to-do action list and should only be access with a valid JWT, it should
// get the data from the DB Accessor service

package main

import (
	"errors"
	"fmt"

	"demo/auth"
	conf "demo/config"
	"demo/kitewrapper"

	"github.com/koding/kite"
)

const (
	User = "username"
	Pass = "password"
)

var BadCredentials = errors.New("Bad credentials")

func main() {
	authorizer := auth.NewAuthorizer()

	authConfig, err := conf.NewConfig("auth")
	if err != nil {
		fmt.Println("Error getting config", err)
		return
	}

	k := kitewrapper.NewKiteWrapper(authConfig)
	err = k.RegisterToKontrol()
	if err != nil {
		fmt.Println("Failed to register", err)
		return
	}

	var dbKite *kite.Client

	go func() {
		var err error

		dbKite, err = k.FindAndDial("db_accessor")
		if err != nil {
			fmt.Println("Failed to dial db service", err)
			// return
		}
	}()

	// Add our handler method
	k.HandleFunc("login", func(r *kite.Request) (interface{}, error) {
		var (
			params, _ = r.Args.One().Map()
			user, _   = params["user"].String()
			pass, _   = params["pass"].String()
		)

		if user == User && pass == Pass {
			return authorizer.Token, nil
		}

		return nil, BadCredentials
	})

	k.HandleFunc("profile", func(r *kite.Request) (interface{}, error) {
		token, err := r.Args.One().String()
		if err != nil {
			return nil, BadCredentials
		}

		if !authorizer.Validate(token) {
			return nil, BadCredentials
		}

		result, err := dbKite.Tell("query", "profile")
		if err != nil {
			return nil, err
		}

		var (
			profileMap, _ = result.Map()
			profile       = make(map[string]string)
		)
		for name, value := range profileMap {
			val, _ := value.String()
			profile[name] = val
		}

		return profile, nil
	})

	k.HandleFunc("validateToken", func(r *kite.Request) (interface{}, error) {
		token, err := r.Args.One().String()

		if err != nil {
			return false, BadCredentials
		}

		return authorizer.Validate(token), nil
	})

	k.Run()
}
