package services

import (
	"errors"

	"github.com/Prabhudatta3004/DLQ/models"
	"github.com/Prabhudatta3004/DLQ/repository"
)

type DLQService interface {
	AddMessage(message *models.Message) error
	GetMessageByID(id uint) (*models.Message, error)
	GetMessageByMessageID(messageID string) (*models.Message, error)
	GetAllMessages() ([]models.Message, error)
	DeleteMessageByID(id uint) error
	DeleteMessageByMessageID(messageID string) error
	ClearMessages() error
}

type dlqService struct {
	repo repository.MessageRepository
}

func NewDLQService(repo repository.MessageRepository) DLQService {
	return &dlqService{repo: repo}
}

func (s *dlqService) AddMessage(message *models.Message) error {
	existingMessage, _ := s.repo.GetByMessageID(message.MessageID)
	if existingMessage != nil && existingMessage.ID != 0 {
		return errors.New("message with this MessageID already exists")
	}
	return s.repo.Create(message)
}

func (s *dlqService) GetMessageByID(id uint) (*models.Message, error) {
	return s.repo.GetByID(id)
}

func (s *dlqService) GetMessageByMessageID(messageID string) (*models.Message, error) {
	return s.repo.GetByMessageID(messageID)
}

func (s *dlqService) GetAllMessages() ([]models.Message, error) {
	return s.repo.GetAll()
}

func (s *dlqService) DeleteMessageByID(id uint) error {
	return s.repo.DeleteByID(id)
}

func (s *dlqService) DeleteMessageByMessageID(messageID string) error {
	return s.repo.DeleteByMessageID(messageID)
}

func (s *dlqService) ClearMessages() error {
	return s.repo.DeleteAll()
}
