package repository

import "fmt"

type Project struct {
	ID          string
	Name        string
	Description string
}

func NewProjectRepo() Project {
	return Project{}
}

var projectMap = make(map[string]Project)

func (p Project) CreateProject() error {
	// create project
	projectMap[p.ID] = p

	fmt.Println("Project created", projectMap)
	return nil
}

func (p Project) GetProjectByID() (*Project, error) {
	// get project by id
	project, ok := projectMap[p.ID]
	if !ok {
		return nil, fmt.Errorf("Project not found")
	}
	return &project, nil
}

func (p Project) GetAllProjects() ([]Project, error) {
	var projects []Project
	for _, project := range projectMap {
		projects = append(projects, project)
	}
	return projects, nil
}

func (p Project) UpdateProject() error {
	_, exists := projectMap[p.ID]

	if !exists {
		return fmt.Errorf("Project not found")

	}
	projectMap[p.ID] = p
	return nil
}

func (p Project) DeleteProject() error {
	_, exists := projectMap[p.ID]

	if !exists {
		return fmt.Errorf("Project not found")
	}
	delete(projectMap, p.ID)
	return nil
}
