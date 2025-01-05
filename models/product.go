package models

type ProductType string

const (
	Auto   ProductType = "auto"
	Manual ProductType = "manual"
)

type Product struct {
	ID                   uint               `json:"product_id"`
	Name                 string             `json:"product_name"`
	Price                int                `json:"product_price"`
	Description          string             `json:"product_description"`
	Type                 ProductType        `json:"product_type"`
	ImageURL             string             `json:"product_image_url"`
	ProductSubCategoryID uint               `json:"product_subcategory_id"`
	ProductSubCategory   ProductSubCategory `gorm:"foreignkey:ProductSubCategoryID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"product_sub_category"`
}
