package main

import (
	"log"

	"starsearch/common/db"
	"starsearch/common/redis"
	ctrl "starsearch/controllers"
	"starsearch/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Unable to load .env file")
	}

	router := gin.Default()
	router.Use(middlewares.Cors())
	DB := db.Connect()
	RDS := redis.Connect()
	ctrl.RegisterRoutes(router, DB, RDS)
	router.Run(":8000")
}
