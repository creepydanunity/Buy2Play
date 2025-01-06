package models

import "time"

type Status string

const (
	Pending   Status = "pending"
	Approved  Status = "approved"
	Completed Status = "completed"
	Rejected  Status = "rejected"
)

type Order struct {
	ID         uint        `gorm:"primary_key" json:"order_id"`
	Timestamp  time.Time   `json:"timestamp"`
	TotalPrice int         `json:"total_price"`
	OrderItems []OrderItem `gorm:"foreignkey:OrderID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"order_items"`
	Status     Status      `json:"status" gorm:"default:'pending'"`
	UserID     uint        `json:"user_id"`
	User       User        `gorm:"foreignkey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`
}

type OrderItem struct {
	ID        uint    `gorm:"primary_key" json:"order_item_id"`
	OrderID   uint    `json:"order_id"`
	Quantity  int     `json:"quantity"`
	ProductID uint    `json:"product_id"`
	Product   Product `gorm:"foreignkey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"product"`
}
