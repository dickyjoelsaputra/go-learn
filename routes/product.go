package routes

import (
	"go_learn/handlers"
	"go_learn/pkg/mysql"
	"go_learn/repositories"

	"github.com/gorilla/mux"
)

func ProductRoutes(r *mux.Router) {
  productRepository := repositories.RepositoryProduct(mysql.DB)
  h := handlers.HandlerProduct(productRepository)

  r.HandleFunc("/products", h.FindProducts).Methods("GET")
  r.HandleFunc("/product/{id}", h.GetProduct).Methods("GET")
  r.HandleFunc("/product", h.CreateProduct).Methods("POST")
}