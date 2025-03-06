package app

import (
	"net/http"

	handler "github.com/kushallabs/taskmanagement/src/handlers/project"

	"github.com/gin-gonic/gin"
	_ "github.com/kushallabs/taskmanagement/docs" // swagger docs
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitHandlerRouter() {
	r := gin.Default()

	r.GET("/ping1", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "in root",
		})
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	projectGroup := r.Group("/project")
	// all handlers related to project

	projectGroup.GET("/getProjectByID/:id", handler.GetProject)
	projectGroup.POST("/", handler.CreateProject)
	projectGroup.PUT("/:id", handler.UpdateProject) // New route
	projectGroup.DELETE("/:id", handler.DeleteProject)
	projectGroup.GET("/", handler.GetAllProjects)
	// 100s of apis

	r.Run(":8081") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
