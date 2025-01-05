package models

import "time"

type Status string

const (
	Pending  Status = "pending"
	Approved Status = "approved"
	Rejected Status = "rejected"
)

type Order struct {
	ID         uint      `gorm:"primary_key" json:"order_id"`
	Timestamp  time.Time `json:"timestamp"`
	TotalPrice int       `json:"total_price"`
	Products   []Product `gorm:"many2many:order_products;" json:"products"`
	Status     Status    `json:"status" gorm:"default:'pending'"`
	UserID     uint      `json:"user_id"`
	User       User      `gorm:"foreignkey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`
}
