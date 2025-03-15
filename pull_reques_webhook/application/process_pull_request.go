package aplication

import (
	"GoAir-WebHooks/pull_reques_webhook/domain"
	"encoding/json"
	"fmt"
	"log"
)

func ProcessPullRequestEvent(rawData []byte) (int, string) {
	var eventPayload domain.PullRequestEventPayload
	var message string
	// Deserializa el JSON recibido en la estructura de datos
	if err := json.Unmarshal(rawData, &eventPayload); err != nil {
		return 403, " "
	}

	// Verifica la acción del Pull Request usando if else
	if eventPayload.Action == "closed" {
		// Cuando el PR es cerrado
		log.Printf("Se hace el pull hacia la rama: %s \n", eventPayload.PullRequest.Base.Ref)
		log.Printf("El pull viene de la rama:  %s \n", eventPayload.PullRequest.Head.Ref)
		log.Printf("Usuario: %s \n", eventPayload.PullRequest.User.Login)
		log.Printf("Nombre del repositorio: %s \n", eventPayload.Repository.FullName)
		
		// Verifica si el PR fue fusionado
		if eventPayload.PullRequest.Merged != nil && *eventPayload.PullRequest.Merged {
			log.Printf("El Pull Request fue fusionado exitosamente\n")
			// Crear mensaje de Pull Request fusionado
			message = createMessage("Pull Request Fusionado Exitosamente", eventPayload.PullRequest.Base.Ref, eventPayload.PullRequest.Head.Ref, eventPayload.PullRequest.User.Login, eventPayload.Repository.FullName)
		} else {
			log.Printf("El Pull Request fue cerrado pero no fusionado\n")
			// Crear mensaje de Pull Request cerrado pero no fusionado
			message = createMessage("Pull Request Cerrado sin Fusionar", eventPayload.PullRequest.Base.Ref, eventPayload.PullRequest.Head.Ref, eventPayload.PullRequest.User.Login, eventPayload.Repository.FullName)
		}

	} else if eventPayload.Action == "reopened" {
		// Cuando el PR es reabierto
		log.Printf("El Pull Request ha sido reabierto\n")
		log.Printf("Se hace el pull hacia la rama: %s \n", eventPayload.PullRequest.Base.Ref)
		log.Printf("El pull viene de la rama:  %s \n", eventPayload.PullRequest.Head.Ref)
		log.Printf("Usuario: %s \n", eventPayload.PullRequest.User.Login)
		log.Printf("Nombre del repositorio: %s \n", eventPayload.Repository.FullName)
		// Crear mensaje de Pull Request reabierto
		message = createMessage("Pull Request Reabierto", eventPayload.PullRequest.Base.Ref, eventPayload.PullRequest.Head.Ref, eventPayload.PullRequest.User.Login, eventPayload.Repository.FullName)

	} else if eventPayload.Action == "ready_for_review" {
		// Cuando el PR está listo para revisión
		log.Printf("El Pull Request está listo para revisión\n")
		log.Printf("Se hace el pull hacia la rama: %s \n", eventPayload.PullRequest.Base.Ref)
		log.Printf("El pull viene de la rama:  %s \n", eventPayload.PullRequest.Head.Ref)
		log.Printf("Usuario: %s \n", eventPayload.PullRequest.User.Login)
		log.Printf("Nombre del repositorio: %s \n", eventPayload.Repository.FullName)
		// Crear mensaje de Pull Request listo para revisión
		message = createMessage("Pull Request Listo para Revisión", eventPayload.PullRequest.Base.Ref, eventPayload.PullRequest.Head.Ref, eventPayload.PullRequest.User.Login, eventPayload.Repository.FullName)

	} else {
		log.Printf("Acción no reconocida: %s\n", eventPayload.Action)
		// En caso de acción no reconocida, se puede generar un mensaje genérico
		message = "Acción de Pull Request no reconocida."
	}

	return 200, message
}

func createMessage(Title string, base string, head string, user string, repository string) string {
	return fmt.Sprint( "Titulo:" + Title + "\n",
					   "Se hace el pull hacia la rama: "+ base + "\n",
					   "El pull viene de la rama: "+ head + "\n", 
					   "Usuario: "+ user + "\n",
					   "Nombre del repositorio: "+ repository + "\n")
} 