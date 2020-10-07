package router
import (
    "github.com/go-chi/chi"
    "github.com/aaronmack/simple_app/app/app"
)
func New() *chi.Mux {
    r := chi.NewRouter()
    r.MethodFunc("GET", "/", app.HandleIndex)
    return r
}