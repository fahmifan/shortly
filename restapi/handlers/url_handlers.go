package handlers

import (
	"github.com/fahmifan/shortly/gen/models"
	"github.com/fahmifan/shortly/gen/restapi/operations/urls"
	"github.com/go-openapi/runtime/middleware"
)

type URLHandler struct {
	Context *Context
}

func NewURLHandler(u *URLHandler) *URLHandler {
	return u
}

func (u *URLHandler) List() urls.ListURLsHandlerFunc {
	return func(urls.ListURLsParams) middleware.Responder {
		return urls.NewListURLsOK().WithPayload([]*models.URL{
			{ID: "123", Shorten: "test"},
		})
	}
}
