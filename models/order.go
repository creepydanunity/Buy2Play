package models

import (
	"buy2play/config"
	"time"
)

type Order struct {
	ID         uint          `gorm:"primary_key" json:"order_id"`
	Timestamp  time.Time     `json:"timestamp"`
	TotalPrice int           `json:"total_price"`
	Products   []Product     `json:"products"`
	Status     config.Status `json:"status" gorm:"type:enum('pending', 'approved', 'rejected');default:'pending'"`
	UserID     uint          `json:"user_id"`
	User       User          `gorm:"foreignkey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`
}
