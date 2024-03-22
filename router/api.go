package router

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"gitlab.com/hedwig-phan/assignment-3/internal/feature"
)

func SetRoutes(app *chi.Mux) {
	app.Get("/welcome", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	app.Post("/test-post", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	app.Post("/fetchgeo", func(w http.ResponseWriter, r *http.Request) {
		err := feature.HandlerData(0)
		if err != nil {
			log.Fatal(err)
		}

		w.Write([]byte("welcome"))
	})
}
