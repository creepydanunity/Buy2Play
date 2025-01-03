package models

type ProductSubCategory struct {
	ID                uint            `json:"subcategory_id"`
	Name              string          `json:"subcategory_name"`
	Description       string          `json:"subcategory_description"`
	ImageURL          string          `json:"subcategory_image_url"`
	ProductCategoryID uint            `json:"product_category_id"`
	ProductCategory   ProductCategory `gorm:"foreignkey:ProductCategoryID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"product_category"`
}
