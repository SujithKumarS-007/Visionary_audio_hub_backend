package controller

import (
	"audiohub/models"
	"audiohub/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Admin Login
func AdminLogin(c *gin.Context) {
	var login models.AdminData
	if err := c.BindJSON(&login); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	token, result := service.AdminLoginCheck(&login)
	if result != 5 {
		c.JSON(http.StatusOK, gin.H{"result": result})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

// Create Admin
func CreateAdmin(c *gin.Context) {
	var admin models.AdminSignup
	if err := c.BindJSON(&admin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	result, data := service.CreateAdmin(admin)
	if result == "Created Successfully" {
		c.JSON(http.StatusOK, gin.H{"result": data})
		return
	}
	c.JSON(http.StatusOK, gin.H{"error": result})
}

// Get Every Data as Single
func GetData(c *gin.Context) {
	var data models.Getdata
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	log.Println(data)
	result,message, err := service.GetData(data)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{"error": message})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": result})
}

// Get all Customer Data
func GetallCustomerdata(c *gin.Context) {
	var token models.Token
	if err := c.BindJSON(&token); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	data, message, err := service.GetallCustomerdata(token)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": message})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": data})
}
// Admin Page Details
func GetAllDetailsForAdmin(c *gin.Context) {
	var token models.Token
	if err := c.BindJSON(&token); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	data,message,err := service.AdminNeededData(token)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{"error": message})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"message": data})
}

// Update All things by Admin
func Update(c *gin.Context) {
	var update models.Update
	if err := c.BindJSON(&update); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	result,err := service.Update(update)
	if err != nil{
		c.JSON(http.StatusOK, gin.H{"error": result})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": result})	
}

// Delete User
func DeleteUser(c *gin.Context) {
	var delete models.Delete
	if err := c.BindJSON(&delete); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	result := service.DeleteUser(delete)
	c.JSON(http.StatusOK, result)
}

// Delete User
func DeleteHistorybyadmin(c *gin.Context) {
	var delete models.Delete
	if err := c.BindJSON(&delete); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	result := service.DeleteHistorybyadmin(delete)
	c.JSON(http.StatusOK, result)
}

// BLock Customer
func Block(c *gin.Context) {
	var data models.Block
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	log.Println(data)
	result, err := service.Block(data)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{"error": result})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": result})
}

// Shut Down
func ShutDown(c *gin.Context) {
	var update models.ShutDown
	if err := c.BindJSON(&update); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	result,err := service.ShutDown(update)
	if err != nil{
		c.JSON(http.StatusOK, gin.H{"error": result})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": result})	
}

// Shut Down
func ClearDB(c *gin.Context) {
	var update models.ClearDB
	if err := c.BindJSON(&update); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	result,err := service.ClearDB(update)
	if err != nil{
		c.JSON(http.StatusOK, gin.H{"error": result})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": result})	
}


