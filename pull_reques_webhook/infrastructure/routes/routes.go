package infrastructure

import (
	infrastructure "GoAir-WebHooks/pull_reques_webhook/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(e *gin.Engine) {
	routesPull := e.Group("/pullRequest")
	{
		routesPull.POST("/proccess", infrastructure.HandlePullRequestEvent)
	}
}
