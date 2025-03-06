package mappers

import (
	controllers "github.com/kushallabs/taskmanagement/src/controllers/project"
	entities "github.com/kushallabs/taskmanagement/src/entities"
)

func ProjectControllerToProjectHandler(controllersProject controllers.Project) entities.Project {
	return entities.Project{
		ID:          controllersProject.ID,
		Name:        controllersProject.Name,
		Description: controllersProject.Description,
	}
}
