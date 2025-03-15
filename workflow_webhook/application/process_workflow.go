package aplication

import (
	"GoAir-WebHooks/workflow_webhook/domain"
	"encoding/json"
	"fmt"
)

func ProcessWorkflowEvent(rawData []byte) (int, string) {
	var eventPayload domain.WebhookPayload
	var message string

	if err := json.Unmarshal(rawData, &eventPayload); err != nil {
		return 403, " "
	}


	// Verificar que action sea "completed"
	if eventPayload.Action == "completed" {
		// Verificar que WorkflowRun y Conclusion no sean nil antes de acceder a ellas
		if eventPayload.WorkflowRun != nil && eventPayload.WorkflowRun.Conclusion != nil {
			if *eventPayload.WorkflowRun.Conclusion == "success" {
				message = createMessage("Éxito en la operación", eventPayload.Repository.URL)
				return 200, message
			}
		}
	}

	message = createMessage("La acción no se completó con éxito", eventPayload.Repository.URL)
	return 200, message
	
}

func createMessage(msg string, repo string) string {
	return fmt.Sprint(msg + "\nRepositorio: " + repo)
}