package infrastructure

import (
	infrastructure "GoAir-WebHooks/workflow_webhook/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(e *gin.Engine) {
	routesPull := e.Group("/actions")
	{
		routesPull.POST("/proccess", infrastructure.HandleActionsEvent)
	}
}
