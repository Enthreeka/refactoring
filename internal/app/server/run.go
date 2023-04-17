package server

import (
	"fmt"
	"net/http"

	"github.com/Enthreeka/refactoring/internal/apperror"
	"github.com/Enthreeka/refactoring/internal/config"
	controller "github.com/Enthreeka/refactoring/internal/user/contoller/http"
	"github.com/Enthreeka/refactoring/internal/user/usecase"
	"github.com/Enthreeka/refactoring/internal/user/usecase/repository"
	"github.com/Enthreeka/refactoring/pkg/logger"
	"github.com/go-chi/chi/v5"
)

func Run(config *config.Config, log *logger.Logger) {

	r := chi.NewRouter()

	apperror.Middleware(r)

	userRepository := repository.NewUser()

	userServiсe := usecase.NewUser(userRepository, log)

	userHandler := controller.NewHandler(userServiсe, log)

	r.Route("/", func(r chi.Router) {
		r.Get("/", userHandler.GeneralPage)
	})

	r.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Route("/users", func(r chi.Router) {
				r.Get("/", userHandler.SearchUsersHandler)
				r.Post("/", userHandler.CreateUserHandler)

				r.Route("/{id}", func(r chi.Router) {
					r.Get("/", userHandler.GetUserHandler)
					r.Patch("/", userHandler.UpdateUserHandler)
					r.Delete("/", userHandler.DeleteUserHandler)
				})
			})
		})
	})

	log.Info("Starting http server: %s:%s", config.Server.TypeServer, config.Server.Port)

	if err := http.ListenAndServe(fmt.Sprintf(":%s", config.Server.Port), r); err != nil {
		log.Fatal("Server listening failed:%s", err)
	}

}
