package main

import (
	"github.com/gin-gonic/gin"
)

type Person struct {
	UserID string `uri:"id" binding:"required"`
	Name   string `uri:"name" binding:"required"`
}

func main() {
	route := gin.Default()
	route.GET("/:name/:id", func(c *gin.Context) {
		var person Person
		if err := c.ShouldBindUri(&person); err != nil {
			c.JSON(400, gin.H{"msg": err})
			return
		}
		c.JSON(200, gin.H{"name": person.Name, "id": person.UserID})
	})
	route.Run(":8080")
}
