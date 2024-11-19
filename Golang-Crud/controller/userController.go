package controllers

import (
	"golang-crud/config"
	"golang-crud/entities"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func ShowUser(c *gin.Context) {
	//get the post data
	var users []entities.User

	config.DB.Find(&users)
	//respond with them

	c.JSON(200, gin.H{
		"Users": users,
	})
}

func ShowUserByID(c *gin.Context) {
	//retrieve data from url parameter
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID Parameter"})
		return
	}

	// query into database based on ID
	var user entities.User
	if err := config.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// return the result value 
	c.JSON(200, gin.H{
		"User": user,
	})
}

func CreateUser(c *gin.Context) {
	//get data from req body
	var body struct {
		Name string		`json:"name"`
		Password string	`json:"password"`
		Email string	`json:"email"`
	}

	// Bind the JSON body to the struct
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	//create user
	user := entities.User{
		Name: body.Name, 
		Password: body.Password, 
		Email: body.Email, 
		Created_at: time.Now()}

	//return the value of the user
	result := config.DB.Create(&user)
	if result.Error != nil {
		c.JSON(500, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	// Return the created user
	c.JSON(200, gin.H{
		"message": "User created successfully",
		"user":    user,
	})
}

func UpdateUserByID(c *gin.Context) {
	// Get the ID from URL
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID Parameter"})
		return
	}

	// Get data from request body
	var body struct {
		Name     string `json:"name"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Find the user data for updating
	var user entities.User
	if err := config.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	// Update the data
	user.Name = body.Name
	user.Password = body.Password
	user.Email = body.Email


	// Return the data
	c.JSON(200, gin.H{
		"message": "User updated successfully",
		"user":    user,
	})
}


func DeleteUserByID (c *gin.Context) {
	// get request data by URL
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID User not found"})

	}

	// get the data body
	var body struct {
		Name     string `json:"name"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// find the data for deleting
	var user entities.User
	if err := config.DB.Delete(&user, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	c.JSON(200, gin.H{
		"message": "User deleted successfully",
		"user":    user,
	})
}	