package handlers

import "github.com/fahmifan/shortly/repository/sqlite"

type Context struct {
	JWTSecretKey  string
	URLRepository *sqlite.URLRepository
}

func NewContext(ctx *Context) *Context {
	return ctx
}
