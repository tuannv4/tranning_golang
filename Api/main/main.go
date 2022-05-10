package main

import (
	"api/package/model"
	"api/package/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func init() {
	var listProduct = []model.Product{
		{ItemName: "product1", UnitPrice: 10, Amount: 15, ItemStatus: 1, ItemDescription: "ok"},
		{ItemName: "product2", UnitPrice: 11, Amount: 16, ItemStatus: 0, ItemDescription: "ok"},
		{ItemName: "product3", UnitPrice: 12, Amount: 17, ItemStatus: 1, ItemDescription: "ok"},
	}
	for _, v := range listProduct {
		v.CreateProduct()
	}
}
func main() {
	r := mux.NewRouter()
	routes.ProductRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8082", r))
}
