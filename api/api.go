package api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/isongjosiah/lernen-api/config"
	"github.com/isongjosiah/lernen-api/deps"
	"github.com/isongjosiah/lernen-api/tracing"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type API struct {
	Server *http.Server
	Config *config.Config
	Deps   *deps.Dependencies
}

func (a *API) Serve() error {
	a.Server = &http.Server{
		Addr:           ":8080", //TODO Josiah fix this with environment variables
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		Handler:        a.SetupServerHandler(), // instead of using the default http.Handler, we define our own handler here that does a couple of middleware checks
		MaxHeaderBytes: 1024 * 1024,
	}
	fmt.Println("API: runing...")

	return a.Server.ListenAndServe()
}

func (a *API) Shutdown() error {
	return a.Server.Shutdown(context.Background())
}

//type Handler func(w http.ResponseWriter, r *http.Request) *ServerResponse

func (a *API) SetupServerHandler() http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.RealIP)
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.Timeout(60 * time.Second))

	mux.Get("/", Home)
	mux.Mount("/user", a.UserRoutes(mux))
	mux.Mount("/auth", a.AuthRoutes(mux))
	mux.Mount("/course", a.CourseRoutes(mux))
	return mux
}
func Home(w http.ResponseWriter, r *http.Request) {
	if _, err := fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path); err != nil {
		return
	}
}

func decodeJSONBody(ctx *tracing.Context, body io.ReadCloser, target interface{}) error {
	//	check if something was delivered in the request body
	err := json.NewDecoder(body).Decode(&target)
	switch {
	case err == io.EOF:
		return errors.New("request body is empty")
	case err != nil:
		return errors.Wrap(err, "error parsing json body for request")
	}

	return nil
}
