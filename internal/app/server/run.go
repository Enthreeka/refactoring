package server

import (
	controller "github.com/Enthreeka/refactoring/internal/contoller/http"
	"github.com/Enthreeka/refactoring/internal/usecase"
	"github.com/Enthreeka/refactoring/internal/usecase/repository"
	"github.com/Enthreeka/refactoring/pkg/logger"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"time"
)

func Run(log *logger.Logger) {

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	//r.Use(apperror.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	userRepository := repository.NewUser()

	userServiсe := usecase.NewUser(userRepository, log)

	userHandler := controller.NewHandler(userServiсe, log)

	//r.Get("/", func(w http.ResponseWriter, r *http.Request) {
	//	w.Write([]byte(time.Now().String()))
	//})

	r.Route("/", func(r chi.Router) {
		r.Get("/", userHandler.GeneralPage)
	})

	r.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Route("/users", func(r chi.Router) {
				r.Get("/", userHandler.SearchUsers)
				r.Post("/", userHandler.CreateUser)

				r.Route("/{id}", func(r chi.Router) {
					r.Get("/", userHandler.GetUser)
					r.Patch("/", userHandler.UpdateUser)
					r.Delete("/", userHandler.DeleteUser)
				})
			})
		})
	})

	http.ListenAndServe(":3333", r)

}
