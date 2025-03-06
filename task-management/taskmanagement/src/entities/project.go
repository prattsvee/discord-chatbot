package entities

// tasks struct
type Task struct {
	ID          int
	Name        string
	Description string
	ProjectID   int
}

// define project struct with fields
type Project struct {
	ID          string `json:"id"`
	Name        string `json:"name" validate:"required, min = 3"`
	Description string `json:"description" validate:"required, min = 10"`
	Tasks       []Task `json:"task"`
}
