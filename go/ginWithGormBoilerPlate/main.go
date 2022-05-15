package main

import (
	"log"

	"boilerplate/libs"
	"boilerplate/models"
	"boilerplate/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db := libs.Database()
	db.AutoMigrate(&models.User{})
	r := gin.Default()

	routes.RouterGroup(r)
	r.Run()
}
