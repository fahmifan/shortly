package main

import (
	"fmt"
	"os"

	"github.com/fahmifan/shortly/gen/restapi"
	"github.com/fahmifan/shortly/gen/restapi/operations"
	"github.com/fahmifan/shortly/repository/sqlite"
	"github.com/fahmifan/shortly/restapi/handlers"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Msg("starting server")
	if err := bootstrap(); err != nil {
		log.Err(err).Msg("")
		os.Exit(1)
	}
	log.Info().Msg("stopping server without error")
}

func bootstrap() error {
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		return fmt.Errorf("load swagger spec: %w", err)
	}

	api := operations.NewShortlyAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer closeErr(server.Shutdown)

	db, err := sqlite.Open("sqlitedb/shortly.db")
	if err != nil {
		return fmt.Errorf("open db: %w", err)
	}

	urlRepo := sqlite.NewURLRepository(&sqlite.URLRepository{
		DB: db,
	})

	ctxHandler := handlers.NewContext(&handlers.Context{
		URLRepository: urlRepo,
	})

	urlHandler := handlers.NewURLHandler(&handlers.URLHandler{Context: ctxHandler})
	authHandler := &handlers.AuthHandler{}

	api.UrlsListURLsHandler = urlHandler.List()
	api.UrlsCreateURLHandler = urlHandler.Create()
	_ = authHandler
	// api.HasRoleAuth = func(token string, scopes []string) (interface{}, error) {
	// 	return authHandler.HasRole(token, scopes)
	// }

	server.SetHandler(api.Serve(middleware.PassthroughBuilder))

	server.Port = 8080

	return server.Serve()
}

// wrap closer function and handle the error
func closeErr(fn func() error) {
	if err := fn(); err != nil {
		fmt.Println("error on close: ", err)
	}
}
