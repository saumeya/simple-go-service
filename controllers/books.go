package controllers

import (
	"net/http"
	"simple-go-service/models"

	"github.com/gin-gonic/gin"
)

type CreateMovieInput struct {
	Title        string `json:"title" binding:"required"`
	Poster_Path  string `json:poster_path"`
	Language     string `json:"language":`
	Overview     string `json:"overview"`
	Release_Date string `json:"release_date`
}

type UpdateMovieInput struct {
	Title        string `json:"title"`
	Poster_Path  string `json:poster_path"`
	Language     string `json:"language":`
	Overview     string `json:"overview"`
	Release_Date string `json:"release_date`
}

func FindMovies(c *gin.Context) {
	var movies []models.Movie
	models.DB.Find(&movies)

	c.JSON(http.StatusOK, gin.H{"data": movies})

}

func CreateMovie(c *gin.Context) {
	var input CreateMovieInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//create a book
	movie := models.Movie{Title: input.Title, Poster_Path: input.Poster_Path, Language: input.Language, Overview: input.Overview, Release_Date: input.Release_Date}
	models.DB.Create(&movie)

	c.JSON(http.StatusOK, gin.H{"data": movie})
}

func FindMovie(c *gin.Context) {
	var movie models.Movie

	if err := models.DB.Where("id = ?", c.Param("id")).First(&movie).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": movie})
}

func UpdateMovie(c *gin.Context) {

	var movie models.Movie
	if err := models.DB.Where("id = ?", c.Param("id")).First(&movie).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input UpdateMovieInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&movie).Updates(models.Movie{Title: input.Title, Poster_Path: input.Poster_Path, Language: input.Language, Overview: input.Overview, Release_Date: input.Release_Date})

	c.JSON(http.StatusOK, gin.H{"data": movie})
}

func DeleteMovie(c *gin.Context) {
	// Get model if exist
	var movie models.Movie
	if err := models.DB.Where("id = ?", c.Param("id")).First(&movie).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&movie)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
