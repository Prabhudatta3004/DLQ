package repository

import (
	"github.com/Prabhudatta3004/DLQ/models"
	"gorm.io/gorm"
)

type MessageRepository interface {
	Create(message *models.Message) error
	GetByID(id uint) (*models.Message, error)
	GetByMessageID(messageID string) (*models.Message, error)
	GetAll() ([]models.Message, error)
	DeleteByID(id uint) error
	DeleteByMessageID(messageID string) error
	DeleteAll() error
}

type messageRepository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) MessageRepository {
	return &messageRepository{db: db}
}

func (r *messageRepository) Create(message *models.Message) error {
	return r.db.Create(message).Error
}

func (r *messageRepository) GetByID(id uint) (*models.Message, error) {
	var message models.Message
	err := r.db.First(&message, id).Error
	return &message, err
}

func (r *messageRepository) GetByMessageID(messageID string) (*models.Message, error) {
	var message models.Message
	err := r.db.Where("message_id = ?", messageID).First(&message).Error
	return &message, err
}

func (r *messageRepository) GetAll() ([]models.Message, error) {
	var messages []models.Message
	err := r.db.Order("created_at desc").Find(&messages).Error
	return messages, err
}

func (r *messageRepository) DeleteByID(id uint) error {
	return r.db.Delete(&models.Message{}, id).Error
}

func (r *messageRepository) DeleteByMessageID(messageID string) error {
	return r.db.Where("message_id = ?", messageID).Delete(&models.Message{}).Error
}

func (r *messageRepository) DeleteAll() error {
	return r.db.Exec("DELETE FROM messages").Error
}
