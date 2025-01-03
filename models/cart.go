package models

type CartItem struct {
	ID        uint    `json:"cart_id"`
	Quantity  uint    `json:"quantity"`
	ProductID uint    `json:"product_id"`
	Product   Product `gorm:"foreignkey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"product"`
	UserID    uint    `json:"user_id"`
	User      User    `gorm:"foreignkey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`
}
