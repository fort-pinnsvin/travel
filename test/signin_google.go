package test

import (
	gooauth2 "github.com/golang/oauth2"
	"github.com/martini-contrib/oauth2"
)

func Signin(tokens oauth2.Tokens) string {
	return tokens.Access()
}
