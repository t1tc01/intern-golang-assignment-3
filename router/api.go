package router

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"gitlab.com/hedwig-phan/assignment-3/internal/feature"
	"gitlab.com/hedwig-phan/assignment-3/internal/util"
	"gitlab.com/hedwig-phan/assignment-3/middleware"
)

func SetRoutes(app *chi.Mux) {
	app.Get("/welcome", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	app.Post("/test-post", func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("Token admin: ", util.Ctx.Value("Authorization"))
		fmt.Println("Admin: ", util.Ctx.Value("Admin"))

		w.Write([]byte("test-post method"))
	})

	app.Post("/fetchgeo", func(w http.ResponseWriter, r *http.Request) {

		user := middleware.ForAdminContext(util.Ctx)
		fmt.Println(user)
		if user == nil || user.Role != 0 {
			w.Write([]byte("Access denied: only admin can use"))
		} else {
			err := feature.HandlerData(100)
			if err != nil {
				log.Fatal(err)
			}

			w.Write([]byte("Fetched 100 sample earthquake"))
		}

	})
}
