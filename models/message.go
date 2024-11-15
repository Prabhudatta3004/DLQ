package models

import "time"

type Message struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	MessageID string    `gorm:"uniqueIndex;not null" json:"message_id" binding:"required"`
	Payload   string    `gorm:"type:text;not null" json:"payload" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
}
