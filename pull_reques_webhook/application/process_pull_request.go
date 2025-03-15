package aplication

import (
	"GoAir-WebHooks/pull_reques_webhook/domain"
	"encoding/json"
	"fmt"
	"log"
)

func ProcessPullRequestEvent(rawData []byte) (int, string) {
	var eventPayload domain.PullRequestEventPayload

	if err := json.Unmarshal(rawData, &eventPayload); err != nil {
		return 403, " "
	}

	message := createMessage(eventPayload.PullRequest.Base.Ref, eventPayload.PullRequest.Head.Ref, eventPayload.PullRequest.User.Login, eventPayload.Repository.FullName)

	if (eventPayload.Action == "closed") {
		log.Printf("Se hace el pull hacia la rama: %s \n", eventPayload.PullRequest.Base.Ref)
		log.Printf("El pull viene de la rama:  %s \n", eventPayload.PullRequest.Head.Ref)
		log.Printf("Usuario: %s \n", eventPayload.PullRequest.User.Login)
		log.Printf("Nombre del repositorio: %s \n", eventPayload.Repository.FullName)
	}

	return 200, message
}

func createMessage(base string, head string, user string, repository string) string {
	return fmt.Sprint("Se hace el pull hacia la rama: "+ base + "\n",
					   "El pull viene de la rama: "+ head + "\n", 
					   "Usuario: "+ user + "\n",
					   "Nombre del repositorio: "+ repository + "\n")
}