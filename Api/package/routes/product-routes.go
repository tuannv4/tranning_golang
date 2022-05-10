package routes

import (
	"api/package/controller"

	"github.com/gorilla/mux"
)

var ProductRoutes = func(route *mux.Router) {
	route.HandleFunc("/product/", controller.CreateProduct).Methods("POST")
	route.HandleFunc("/product/", controller.GetProducts).Methods("GET")
	route.HandleFunc("/product/{productId}", controller.GetProductById).Methods("GET")
	route.HandleFunc("/product/{productId}", controller.UpdateProduct).Methods("PUT")
	route.HandleFunc("/product/{productId}", controller.DeleteProduct).Methods("DELETE")
}
