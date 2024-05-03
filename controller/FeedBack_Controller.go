package controller

import (
	"audiohub/models"
	"audiohub/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InsertFeedback(c *gin.Context) {
	var feedback models.Feedback

	if err := c.BindJSON(&feedback); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	result, err := service.InstertFeedback(feedback)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": result})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": result})
}

// Delete FeedBack
func Deletefeedback(c *gin.Context) {
	var feedback models.FeedbackDB

	if err := c.BindJSON(&feedback); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	log.Println(feedback)
	result := service.Deletefeedback(feedback)
	c.JSON(http.StatusOK, result)

}

// Get All FeedBacks
func GetFeedbacks(c *gin.Context) {
	var token models.Token
	if err := c.BindJSON(&token); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	data,message,err := service.GetFeedBacks(token)
	if err != nil{
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{"error": message})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": data})

}
