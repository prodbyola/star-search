// app's routes and controllers
package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"gorm.io/gorm"
)

type handler struct {
	DB  *gorm.DB
	RDS redis.Conn
}

func RegisterRoutes(r *gin.Engine, DB *gorm.DB, RDS redis.Conn) {
	H := &handler{
		DB:  DB,
		RDS: RDS,
	}

	r.GET("/get-movies", H.GetMovies)
	r.POST("/add-comment", H.AddComment)
	r.GET("/movie-comments/:movie_id", H.GetMovieComments)
	r.GET("/movie-characters/:movie_id", H.GetCharacters)
}
