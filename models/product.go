package models

import "buy2play/config"

type Product struct {
	ID                   uint               `json:"product_id"`
	Name                 string             `json:"product_name"`
	Price                int                `json:"product_price"`
	Description          string             `json:"product_description"`
	Type                 config.ProductType `gorm:"type:enum('auto', 'manual')" json:"product_type"`
	ImageURL             string             `json:"product_image_url"`
	ProductSubCategoryID uint               `json:"product_subcategory_id"`
	ProductSubCategory   ProductSubCategory `gorm:"foreignkey:ProductSubCategoryID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"product_sub_category"`
}
