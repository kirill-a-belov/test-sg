package models

import "github.com/jinzhu/gorm"

type ByURLResponse struct {
	Products []*Product `json:"products"`
}

type ByIDResponse struct {
	Product *Product `json:"product"`
}

type Product struct {
	gorm.Model `json:"-"`
	PID        string `gorm:"unique;not null" json:pid"`
	URL        string `json:"url"`
	Title      string `json:"title"`
	Price      string `json:"price"`
	Image      string `json:"image"`
	IsInStock  bool   `json:"is_in_stock"`
}

func (Product) TableName() string {
	return "requests"
}
