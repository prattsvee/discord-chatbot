package handler

import (
	"fmt"
	"net/http"

	controllers "github.com/kushallabs/taskmanagement/src/controllers/project"
	"github.com/kushallabs/taskmanagement/src/entities"

	"github.com/gin-gonic/gin"
)

// create project, get project, get projects, update project, delete project

// GetProjects returns all projects
func GetProject(c *gin.Context) {

	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	fmt.Println(id, "is from handlers")

	projectController := controllers.Project{
		ID: id,
	}

	resp, err := projectController.GetProjectByID()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)

}

func CreateProject(c *gin.Context) {

	var project entities.Project

	// concert json inpt from post and set data to project object from entity paakge
	if err := c.ShouldBindJSON(&project); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(project, "is from handlers")

	projectController := controllers.NewProjectController(project.ID, project.Name, project.Description)

	// TODO return error from create project in controller and handle it here
	projectController.CreateProject()

	c.JSON(http.StatusOK, project)

	// we get data from POST method
}

func GetAllProjects(c *gin.Context) {
	projectController := controllers.Project{}

	projects, err := projectController.GetAllProjects()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, projects)
}

func UpdateProject(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	var project entities.Project
	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	project.ID = id
	projectController := controllers.NewProjectController(
		project.ID,
		project.Name,
		project.Description,
	)

	if err := projectController.UpdateProject(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, project)
}
func DeleteProject(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	projectController := controllers.Project{
		ID: id,
	}

	if err := projectController.DeleteProject(); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Project deleted successfully"})
}
