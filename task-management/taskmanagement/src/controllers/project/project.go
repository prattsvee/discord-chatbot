package controllers

import (
	"errors"
	"fmt"

	"github.com/google/uuid"

	repository "github.com/kushallabs/taskmanagement/src/repository/db/inmemory/maps/project"
)

type Project struct {
	ID          string
	Name        string
	Description string
}

func NewProjectController(id, name, description string) *Project {
	return &Project{
		ID:          id,
		Name:        name,
		Description: description,
	}
}

func (p Project) CreateProject() error {

	// validate project fields
	if p.Name == "" {
		return errors.New("name is required")
	}

	// create project
	p.ID = uuid.NewString()

	fmt.Println(p, "is from controller")

	// call repository layer
	projRepo := repository.Project{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
	}

	return projRepo.CreateProject()
}

func (p Project) GetProjectByID() (*Project, error) {
	// get project by id
	projRepo := repository.Project{
		ID: p.ID,
	}

	fmt.Println(p.ID, "is from controller")
	project, err := projRepo.GetProjectByID()

	if err != nil {
		return nil, err
	}
	fmt.Println(project, "is from controller")

	// convert repository project to controller project
	projectController := &Project{
		ID:          project.ID,
		Name:        project.Name,
		Description: project.Description,
	}

	return projectController, nil
}

func (p Project) GetAllProjects() ([]Project, error) {
	projRepo := repository.Project{}
	projects, err := projRepo.GetAllProjects()
	if err != nil {
		return nil, err
	}

	var projectControllers []Project
	for _, proj := range projects {
		projectControllers = append(projectControllers, Project{
			ID:          proj.ID,
			Name:        proj.Name,
			Description: proj.Description,
		})
	}
	return projectControllers, nil
}

func (p Project) UpdateProject() error {
	if p.Name == "" {
		return errors.New("Name is required")
	}
	projRepo := repository.Project{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
	}
	return projRepo.UpdateProject()
}

func (p Project) DeleteProject() error {

	if p.ID == "" {
		return errors.New("ID is required")
	}
	projRepo := repository.Project{
		ID: p.ID,
	}
	return projRepo.DeleteProject()
}
