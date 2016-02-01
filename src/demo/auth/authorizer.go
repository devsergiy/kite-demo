package auth

import jwt "github.com/dgrijalva/jwt-go"

var (
	FooClaim = "bar"
	SignKey  = []byte("sdfhgsdjagdfghjsgd")
)

type Authorizer struct {
	Token    string
	claimFoo string
}

func NewAuthorizer() *Authorizer {
	auth := &Authorizer{}
	auth.setupToken()

	return auth
}

func (a *Authorizer) setupToken() {
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims["foo"] = FooClaim
	tokenString, err := token.SignedString(SignKey)

	if err != nil {
		panic(err)
	}

	a.Token = tokenString
	a.claimFoo = FooClaim
}

func (a *Authorizer) Validate(myToken string) bool {
	token, err := jwt.Parse(myToken, func(token *jwt.Token) (interface{}, error) {
		return SignKey, nil
	})
	foo := token.Claims["foo"]

	if err == nil && token.Valid && foo == FooClaim {
		return true
	}

	return false
}
