package router

import (
	"github.com/go-chi/chi"

	"github.com/aaronmack/simple_app/app/app"
	"github.com/aaronmack/simple_app/app/requestlog"
	"github.com/aaronmack/simple_app/app/router/middleware"
)

func New(a *app.App) *chi.Mux {
	l := a.Logger()

	r := chi.NewRouter()

	// Routes for healthz
	r.Get("/healthz", app.HandleHealth)

	// Routes for APIs
	r.Route("/api/v1", func(r chi.Router) {
		r.Use(middleware.ContentTypeJson)

		// Routes for books
		r.Method("GET", "/books", requestlog.NewHandler(a.HandleListBooks, l))
		r.Method("POST", "/books", requestlog.NewHandler(a.HandleCreateBook, l))
		r.Method("GET", "/books/{id}", requestlog.NewHandler(a.HandleReadBook, l))
		r.Method("PUT", "/books/{id}", requestlog.NewHandler(a.HandleUpdateBook, l))
		r.Method("DELETE", "/books/{id}", requestlog.NewHandler(a.HandleDeleteBook, l))
	})

	r.Method("GET", "/", requestlog.NewHandler(a.HandleIndex, l))

	return r
}
