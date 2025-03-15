package aplication

import (
	"GoAir-WebHooks/workflow_webhook/domain"
	"encoding/json"
	"fmt"
)

func ProcessWorkflowEvent(rawData []byte) (int, string) {
	var eventPayload domain.WebhookPayload
	var workflow_run domain.WorkflowRun
	var message string

	if err := json.Unmarshal(rawData, &eventPayload); err != nil {
		return 403, " "
	}


	if (*&eventPayload.Action == "completed") {
		if *workflow_run.Conclusion == "success" {
			message = createMessage("Éxito en la operación", eventPayload.Repository.URL)
			return 200, message
		}
	} 

	message = createMessage("La acción no se completo con éxito", eventPayload.Repository.URL)
	return 200, message

}

func createMessage(msg string, repo string) string {
	return fmt.Sprint(msg + "\nRepositorio: " + repo)
}