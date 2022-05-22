package handlers

import (
	"math/rand"
	"net/http"

	"github.com/fahmifan/shortly/gen/models"
	"github.com/fahmifan/shortly/gen/restapi/operations/urls"
	"github.com/fahmifan/shortly/repository/sqlite"
	"github.com/fahmifan/shortly/ulids"
	"github.com/go-openapi/runtime/middleware"
	"github.com/jxskiss/base62"
	"github.com/rs/zerolog/log"
)

type URLHandler struct {
	Context *Context
}

func NewURLHandler(u *URLHandler) *URLHandler {
	return u
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
		Shorten:  &in.Shorten,
	}
}

func (u *URLHandler) List() urls.ListURLsHandlerFunc {
	return func(params urls.ListURLsParams) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		userID := ulids.ULID{}
		listURLs, err := u.Context.URLRepository.ListByUserID(ctx, userID, sqlite.ListFilter{})
		if err != nil {
			log.Err(err).Msg("internal")
			return urls.NewListURLsDefault(http.StatusInternalServerError).WithPayload(&models.Error{
				Code: http.StatusInternalServerError,
			})
		}
		return urls.NewListURLsOK().WithPayload(serializeListURLs(listURLs))
	}
}

func (u *URLHandler) Create() urls.CreateURLHandlerFunc {
	return func(params urls.CreateURLParams) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		shorten := ""
		if *params.URL.Shorten == "" {
			shorten = string(base62.FormatUint(uint64(rand.Uint32())))
		}
		url := sqlite.URL{
			ID:       ulids.New(),
			Original: *params.URL.Original,
			IsPublic: *params.URL.IsPublic,
			Shorten:  shorten,
		}
		err := u.Context.URLRepository.Create(ctx, &url)
		if err != nil {
			log.Err(err).Msg("create")
			return urls.NewCreateURLInternalServerError().WithPayload(&models.Error{
				Code: 500,
			})
		}
		return urls.NewCreateURLOK().WithPayload(serializeURL(&url))
	}
}
