package routes

import (
	"go_learn/handlers"
	"go_learn/pkg/mysql"
	"go_learn/repositories"

	"github.com/gorilla/mux"
)

func ProfileRoutes(r *mux.Router) {
  profileRepository := repositories.RepositoryProfile(mysql.DB)
  h := handlers.HandlerProfile(profileRepository)

  r.HandleFunc("/profile/{id}", h.GetProfile).Methods("GET")
}