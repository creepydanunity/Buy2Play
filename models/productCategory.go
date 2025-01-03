package models

type ProductCategory struct {
	ID          uint   `json:"category_id"`
	Name        string `json:"category_name"`
	Description string `json:"category_description"`
	ImageURL    string `json:"category_image_url"`
}
