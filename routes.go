package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// Get all todos
	r.GET("/todos", func(c *gin.Context) {
		var todos []Todo
		db.Find(&todos)

		c.JSON(200, todos)
	})

	// Create a new todo
	r.POST("/todos", func(c *gin.Context) {
		var todo Todo
		c.BindJSON(&todo)

		db.Create(&todo)

		c.JSON(200, todo)
	})

	// Get a specific todo
	r.GET("/todos/:id", func(c *gin.Context) {
		var todo Todo
		db.First(&todo, c.Param("id"))

		c.JSON(200, todo)
	})

	// Update a todo
	r.PUT("/todos/:id", func(c *gin.Context) {
		var todo Todo
		db.First(&todo, c.Param("id"))

		c.BindJSON(&todo)

		db.Save(&todo)

		c.JSON(200, todo)
	})

	// Delete a todo
	r.DELETE("/todos/:id", func(c *gin.Context) {
		var todo Todo
		db.Delete(&todo, c.Param("id"))

		c.Status(204)
	})

	return r
}
