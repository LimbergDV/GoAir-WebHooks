package domain

type WebhookPayload struct {
    Action      string       `json:"action"`
    Repository  Repository   `json:"repository"`
    Sender      Sender       `json:"sender"`
    Workflow    Workflow     `json:"workflow"`
    WorkflowRun *WorkflowRun  `json:"workflow_run"`
}

type Repository struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
    URL  string `json:"url"`
}

type Sender struct {
    Login string `json:"login"`
    ID    int    `json:"id"`
    URL   string `json:"url"`
}

type Workflow struct {
    ID        int    `json:"id"`
    Name      string `json:"name"`
    State     string `json:"state"`
    BadgeURL  string `json:"badge_url"`
    CreatedAt string `json:"created_at"`
    UpdatedAt string `json:"updated_at"`
    URL       string `json:"url"`
}

type WorkflowRun struct {
    ID          int     `json:"id"`
    Name        string  `json:"name"`
    Status      string  `json:"status"`
    Conclusion  *string `json:"conclusion"`
    CreatedAt   string  `json:"created_at"`
    RunNumber   int     `json:"run_number"`
    RunStartedAt string `json:"run_started_at"`
    WorkflowID  int     `json:"workflow_id"`
    WorkflowURL string  `json:"workflow_url"`
}