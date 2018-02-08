package main

import (
	"go-rest-mgo/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		movie := new(controllers.MovieController)
		v1.POST("/movies", movie.Create)
		v1.GET("/movies", movie.Find)
		v1.GET("/movies/:id", movie.Get)
		v1.PUT("/movies/:id", movie.Update)
		v1.DELETE("/movies/:id", movie.Delete)
	}

	router.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "Not Found")
	})

	router.Run()
}
