package handlers

import "github.com/fahmifan/shortly/repository/sqlite"

type Context struct {
	URLRepository *sqlite.URLRepository
}

func NewContext(ctx *Context) *Context {
	return ctx
}
