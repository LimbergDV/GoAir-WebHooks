package main

import (
	iPR "GoAir-WebHooks/pull_reques_webhook/infrastructure/routes"
	iWF "GoAir-WebHooks/workflow_webhook/infrastructure/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	iPR.RegisterRouter(r)
	iWF.RegisterRouter(r)

	r.Run()
}