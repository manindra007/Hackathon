package main

import (
	"fmt"
	alert "hackathon/server/src/Alert"
	"hackathon/server/src/person"
	"net/http"

	"github.com/gin-gonic/gin"
)

func registerUser(c *gin.Context) {
	var input person.Person
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if input.Email == "" {
		c.JSON(http.StatusBadRequest, "Email field can't be nil")
		return
	}
	for _, j := range person.Users {

		if j.Email == input.Email {
			c.JSON(http.StatusBadRequest, "User already exists!")
			return
		}
	}
	fmt.Println(input)

	person.Users = append(person.Users, input)
	fmt.Println(person.Users)
	c.JSON(http.StatusOK, "Registered Succefully!")
}

type Email struct {
	Email string `json:"email"`
}

func allert(c *gin.Context) {
	fmt.Println("here")
	var input Email
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if resp := alert.SendAlert(input.Email); resp {
		c.JSON(http.StatusOK, "Alert Send!")
	} else {
		c.JSON(http.StatusOK, "Alert aborted!")
	}
}

func abort(c *gin.Context) {
	fmt.Println("here")
	var input Email
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	alert.AbortAlert(input.Email)
	c.JSON(http.StatusOK, "Alert aborted!")
}

func main() {
	router := gin.Default()
	router.POST("/register", registerUser)
	// router.POST("/updateUser", updateUser)
	router.POST("/allert", allert)
	router.POST("/abort", abort)

	router.Run("localhost:8080")
}
