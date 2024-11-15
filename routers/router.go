package routers

import (
	"github.com/Prabhudatta3004/DLQ/controllers"
	"github.com/Prabhudatta3004/DLQ/repository"
	"github.com/Prabhudatta3004/DLQ/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	// Initialize components
	messageRepo := repository.NewMessageRepository(db)
	dlqService := services.NewDLQService(messageRepo)
	dlqController := controllers.NewDLQController(dlqService)

	// Define routes
	dlq := router.Group("/dlq")
	{
		dlq.POST("/message", dlqController.AddMessage)
		dlq.GET("/message/id/:id", dlqController.GetMessageByID)
		dlq.GET("/message/message_id/:message_id", dlqController.GetMessageByMessageID)
		dlq.GET("/messages", dlqController.GetAllMessages)
		dlq.DELETE("/message/id/:id", dlqController.DeleteMessageByID)
		dlq.DELETE("/message/message_id/:message_id", dlqController.DeleteMessageByMessageID)
		dlq.DELETE("/messages", dlqController.ClearMessages)
	}

	return router
}
