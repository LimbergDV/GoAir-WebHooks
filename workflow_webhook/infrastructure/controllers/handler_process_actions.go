package infrastructure

import (
	aplication "GoAir-WebHooks/workflow_webhook/application"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleActionsEvent(ctx *gin.Context) {

	eventType := ctx.GetHeader("X-GitHub-Event")
	// deliveryD := ctx.GetHeader("X-GitHub-Delivery")

	// log.Printf("Nuevo evento: %s con ID: %s", eventType, deliveryD)

	rawData, err := ctx.GetRawData()

	if err != nil {
		log.Printf("Error al leer el cuerpo de la solicitud: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "leer datos"})
	}

	var statusCode int

	switch eventType {
	case "ping":
		ctx.JSON(http.StatusOK, gin.H{"status": "success"})
	case "workflow_run":
		status, msg := aplication.ProcessWorkflowEvent(rawData)

		statusCode = status

		log.Print(msg)

		if msg != " " {
			SendMessageToDiscord(msg)
		}
	}

	switch statusCode {
	case 200:
		ctx.JSON(http.StatusOK, gin.H{"success": "Pull Request procesado con éxito"})
	case 403:
		log.Printf("Error al deserializar el payload del pull request: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al procesar el payload del pull request"})
	default:
		ctx.JSON(http.StatusOK, gin.H{"success": "Normal"})
	}

}
