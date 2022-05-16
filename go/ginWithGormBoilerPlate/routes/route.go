package routes

import (
	"fmt"
	"strings"

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

		test.GET("/string", func(c *gin.Context) {
			oldString := "안녕|하세|요"
			newString := "요|sd|하세|감자|수박"

			if len(newString) == 0 {
				c.JSON(400, gin.H{"msg": "적어도 한 자 이상의 단어를 입력하세요"})
				return
			}

			respString := oldString
			addedString := ""
			if len(oldString) == 0 {
				respString = newString
			} else {
				oldSplit := strings.Split(oldString, "|")
				newSplit := strings.Split(newString, "|")

				for i := 0; i < len(newSplit); i++ {
					isAdd := false
					for j := 0; j < len(oldSplit); j++ {
						if newSplit[i] == oldSplit[j] {
							isAdd = false
							break
						}
						isAdd = true
					}
					if isAdd {
						if len(addedString) == 0 {
							addedString += newSplit[i]
							continue
						}
						respString += "|" + newSplit[i]
						addedString += "|" + newSplit[i]
					}
				}
			}

			c.JSON(200, gin.H{"msg": addedString})
		})
	}
}
