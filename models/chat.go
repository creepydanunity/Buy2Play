package models

import "time"

// Conversation represents a conversation between a user and the store admin
type Conversation struct {
	ID        uint      `json:"conversation_id"`
	OrderID   uint      `json:"order_id"`
	Order     Order     `gorm:"foreignkey:OrderID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"order"`
	UserID    uint      `json:"user_id"`
	User      User      `gorm:"foreignkey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`
	AdminID   uint      `json:"admin_id"`
	Admin     User      `gorm:"foreignkey:AdminID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"admin"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Message represents a message in a conversation
type Message struct {
	ID             uint         `json:"message_id"`
	ConversationID uint         `json:"conversation_id"`
	Conversation   Conversation `gorm:"foreignkey:ConversationID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"conversation,omitempty"`
	SenderID       uint         `json:"sender_id"`
	Sender         User         `gorm:"foreignkey:SenderID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"sender"`
	Content        string       `json:"content"`
	CreatedAt      time.Time    `gorm:"not null;default:current_timestamp" json:"created_at"`
}
