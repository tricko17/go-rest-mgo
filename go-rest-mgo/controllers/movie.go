package controllers

import (
	"fmt"
	"go-rest-mgo/forms"
	"go-rest-mgo/models"

	"github.com/gin-gonic/gin"
)

var movieModel = new(models.MovieModel)

type MovieController struct{}

func (movie *MovieController) Create(c *gin.Context) {
	var data forms.CreateMovieCommand
	if c.BindJSON(&data) != nil {
		c.JSON(406, gin.H{"message": "Invalid form", "form": data})
		c.Abort()
		return
	}

	err := movieModel.Create(data)
	if err != nil {
		c.JSON(406, gin.H{"message": "Movie could not be created", "error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"message": "Movie Created"})
}

func (movie *MovieController) Find(c *gin.Context) {
	list, err := movieModel.Find()
	if err != nil {
		c.JSON(404, gin.H{"message": "Find Error", "error": err.Error()})
		c.Abort()
	} else {
		c.JSON(200, gin.H{"data": list})
	}
}

func (movie *MovieController) Get(c *gin.Context) {
	id := c.Param("id")
	profile, err := movieModel.Get(id)
	if err != nil {
		c.JSON(404, gin.H{"message": "Movie not found", "error": err.Error()})
		c.Abort()
	} else {
		c.JSON(200, gin.H{"data": profile})
	}
}

func (movie *MovieController) Update(c *gin.Context) {
	id := c.Param("id")
	data := forms.UpdateMovieCommand{}

	if c.BindJSON(&data) != nil {
		c.JSON(406, gin.H{"message": "Invalid Parameters"})
		c.Abort()
		return
	}
	err := movieModel.Update(id, data)
	if err != nil {
		c.JSON(406, gin.H{"message": "Movie Could Not Be Updated", "error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{"message": "Movie Updated"})
}

func (movie *MovieController) Delete(c *gin.Context) {
	id := c.Param("id")
	fmt.Println(id)
	err := movieModel.Delete(id)
	if err != nil {
		c.JSON(406, gin.H{"message": "Movie Could Not Be Deleted", "error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{"message": "Movie Deleted"})
}
