package domain

import "time"

type Workflow struct {
	BadgeURL   string    `json:"badge_url"`   // URL del badge
	CreatedAt  time.Time `json:"created_at"`  // Fecha de creación
	HTMLURL    string    `json:"html_url"`    // URL HTML
	ID         int       `json:"id"`          // ID del workflow
	Name       string    `json:"name"`        // Nombre del workflow
	NodeID     string    `json:"node_id"`     // Identificador único
	Path       string    `json:"path"`        // Ruta
	State      string    `json:"state"`       // Estado del workflow
	UpdatedAt  time.Time `json:"updated_at"`  // Fecha de actualización
	URL        string    `json:"url"`         // URL API
}


type PullRequest struct {
	ID     int    `json:"id"`
	Number int    `json:"number"`
	URL    string `json:"url"`
}

type WorkflowRun struct {
	Conclusion          *string             `json:"conclusion"` // Puede ser null
	CreatedAt           time.Time           `json:"created_at"`
	Event               string              `json:"event"`
	HeadBranch          *string             `json:"head_branch"` // Puede ser null
	PullRequests        []PullRequest       `json:"pull_requests"`
	RunStartedAt        time.Time           `json:"run_started_at"`
	Status              string              `json:"status"`
}
