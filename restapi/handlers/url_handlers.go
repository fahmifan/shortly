package handlers

import (
	"net/http"

	"github.com/fahmifan/shortly/gen/models"
	"github.com/fahmifan/shortly/gen/restapi/operations/urls"
	"github.com/fahmifan/shortly/repository/sqlite"
	"github.com/go-openapi/runtime/middleware"
	"github.com/oklog/ulid"
)

type URLHandler struct {
	Context *Context
}

func NewURLHandler(u *URLHandler) *URLHandler {
	return u
}

func (u *URLHandler) List() urls.ListURLsHandlerFunc {
	return func(params urls.ListURLsParams, i interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		userID := ulid.ULID{}
		listURLs, err := u.Context.URLRepository.ListByUserID(ctx, userID, sqlite.ListFilter{})
		if err != nil {
			return urls.NewListURLsDefault(http.StatusInternalServerError).WithPayload(&models.Error{
				Code: http.StatusInternalServerError,
			})
		}
		return urls.NewListURLsOK().WithPayload(serializeListURLs(listURLs))
	}
}

func serializeListURLs(listURLs []sqlite.URL) []*models.URL {
	urls := make([]*models.URL, len(listURLs))
	for i := range listURLs {
		urls[i] = serializeURL(&listURLs[i])
	}

	return urls
}

func serializeURL(in *sqlite.URL) *models.URL {
	if in == nil {
		return nil
	}

	return &models.URL{
		ID:       in.ID.String(),
		IsPublic: &in.IsPublic,
		Original: &in.Original,
		Shorten:  in.Shorten,
	}
}
