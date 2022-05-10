package controller

import (
	"api/package/model"
	"api/package/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	newProducts := model.GetListProduct()
	res, _ := json.Marshal(newProducts)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetProductById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productId := vars["productId"]
	ID, err := strconv.ParseInt(productId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	product, _ := model.GetProductById(ID)
	res, _ := json.Marshal(product)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	productCreate := &model.Product{}
	utils.ParseBody(r, productCreate)
	p := productCreate.CreateProduct()
	res, _ := json.Marshal(p)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var updateProduct = &model.Product{}
	utils.ParseBody(r, updateProduct)
	vars := mux.Vars(r)
	productId := vars["productId"]
	ID, err := strconv.ParseInt(productId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	productDetails, db := model.GetProductById(ID)

	if updateProduct.ItemName != "" {
		productDetails.ItemName = updateProduct.ItemName
	}
	if updateProduct.UnitPrice != 0 {
		productDetails.UnitPrice = updateProduct.UnitPrice
	}
	if updateProduct.Amount != 0 {
		productDetails.Amount = updateProduct.Amount
	}
	if updateProduct.ItemStatus != 0 {
		productDetails.ItemStatus = updateProduct.ItemStatus
	}
	if updateProduct.ItemDescription != "" {
		productDetails.ItemDescription = updateProduct.ItemDescription
	}

	db.Save(&productDetails)
	//Convert object string -> json encodings
	res, _ := json.Marshal(productDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productId := vars["productId"]
	ID, err := strconv.ParseInt(productId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	p := model.DeleteProduct(ID)
	res, _ := json.Marshal(p)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
