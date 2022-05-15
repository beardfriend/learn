package routes

import (
	"fmt"

	"boilerplate/libs"
	"boilerplate/models"

	"github.com/gin-gonic/gin"
)

func RouterGroup(c *gin.Engine) {
	test := c.Group("/test")
	{
		test.GET("", func(c *gin.Context) {
			// user := &models.User{Name: "helo"}
			// libs.Database().Save(user)
			user := models.User{ID: 1}
			var users models.User
			libs.Database().First(&user).Scan(&users)
			fmt.Print(users)
			c.JSON(200, gin.H{"message": "hello"})
		})
	}
}
