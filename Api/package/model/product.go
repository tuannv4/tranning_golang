package model

import (
	"api/package/config"

	"gorm.io/gorm"
)

var db *gorm.DB

type Product struct {
	gorm.Model
	ItemName        string  `gorm:"size:255;not null;" json:"itemName"`
	UnitPrice       float32 `gorm:"" json:"unitprice"`
	Amount          float32 `gorm:"" json:"amount"`
	ItemStatus      int     `gorm:"" json:"itemstatus"`
	ItemDescription string  `gorm:"default:null" json:"itemdescription"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Product{})
}

func (p *Product) CreateProduct() *Product {
	db.Create(&p)
	return p
}

func GetListProduct() []Product {
	var products []Product
	db.Find(&products)
	return products
}

func GetProductById(Id int64) (*Product, *gorm.DB) {
	var product Product
	db := db.Where("ID=?", Id).Find(&product)
	return &product, db
}

func (p *Product) UpdateProduct() *Product {
	db.Save(&p)
	return p
}

func DeleteProduct(Id int64) Product {
	var product Product
	db.Delete(&Product{}, Id)
	return product
}
