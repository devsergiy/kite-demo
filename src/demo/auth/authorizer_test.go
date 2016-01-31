package auth_test

import (
	"testing"

	"demo/auth"
)

const (
	ExpectedToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIifQ.2mR6Oa_ZjoYZBK4Mayd6jYgSpe_z0HZQS_cBEEdkSjU"
)

func TestNewAuthorizer(t *testing.T) {
	auth := auth.NewAuthorizer()

	if auth.Token != ExpectedToken {
		t.Error("Expected"+ExpectedToken+", got ", auth.Token)
	}
}

func TestValidate(t *testing.T) {
	auth := auth.NewAuthorizer()

	if !auth.Validate(ExpectedToken) {
		t.Error("Expect to return true")
	}
}
