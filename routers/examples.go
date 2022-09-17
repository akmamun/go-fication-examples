package routers

import (
	"github.com/go-chi/chi/v5"
	"go-fication-examples/controllers"
	"go-fication-examples/infra/database"
	"go-fication-examples/repository"
)

func ExamplesRoutes(router *chi.Mux, db *database.DB) {
	repo := repository.NewGormRepository(db)
	exampleCtrl := controllers.NewExampleHandler(repo)
	router.Group(func(r chi.Router) {
		r.Get("/test/", exampleCtrl.GetExamplesListData)
		r.Post("/test/", exampleCtrl.CreateData)
		r.Get("/test/{id}", exampleCtrl.GetOne)
		r.Patch("/test/{id}", exampleCtrl.GetOneAndUpdate)

	})
}
