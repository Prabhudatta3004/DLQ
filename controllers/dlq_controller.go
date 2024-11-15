package controllers

import (
	"net/http"
	"strconv"

	"github.com/Prabhudatta3004/DLQ/models"
	"github.com/Prabhudatta3004/DLQ/services"
	"github.com/Prabhudatta3004/DLQ/utils"
	"github.com/gin-gonic/gin"
)

type DLQController struct {
	service services.DLQService
}

func NewDLQController(service services.DLQService) *DLQController {
	return &DLQController{service: service}
}

func (ctrl *DLQController) AddMessage(c *gin.Context) {
	var msg models.Message
	if err := c.ShouldBindJSON(&msg); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := ctrl.service.AddMessage(&msg); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}
	utils.RespondWithSuccess(c, gin.H{"message": "Message added to DLQ", "data": msg})
}

func (ctrl *DLQController) GetMessageByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid ID")
		return
	}
	msg, err := ctrl.service.GetMessageByID(uint(id))
	if err != nil {
		utils.RespondWithError(c, http.StatusNotFound, "Message not found")
		return
	}
	utils.RespondWithSuccess(c, msg)
}

func (ctrl *DLQController) GetMessageByMessageID(c *gin.Context) {
	messageID := c.Param("message_id")
	msg, err := ctrl.service.GetMessageByMessageID(messageID)
	if err != nil {
		utils.RespondWithError(c, http.StatusNotFound, "Message not found")
		return
	}
	utils.RespondWithSuccess(c, msg)
}

func (ctrl *DLQController) GetAllMessages(c *gin.Context) {
	messages, err := ctrl.service.GetAllMessages()
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithSuccess(c, messages)
}

func (ctrl *DLQController) DeleteMessageByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid ID")
		return
	}
	if err := ctrl.service.DeleteMessageByID(uint(id)); err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithSuccess(c, gin.H{"message": "Message deleted from DLQ"})
}

func (ctrl *DLQController) DeleteMessageByMessageID(c *gin.Context) {
	messageID := c.Param("message_id")
	if err := ctrl.service.DeleteMessageByMessageID(messageID); err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithSuccess(c, gin.H{"message": "Message deleted from DLQ"})
}

func (ctrl *DLQController) ClearMessages(c *gin.Context) {
	if err := ctrl.service.ClearMessages(); err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithSuccess(c, gin.H{"message": "All messages cleared from DLQ"})
}
