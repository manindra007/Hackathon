package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type EmergencyContact struct {
	Phone string `json:"phone"`
	Email string `json:"email"`
}

type Person struct {
	Id               string             `gorm:"primary_key" json:"id"`
	Email            string             `gorm:"unique" json:"email"`
	Name             string             `json:"name"`
	Address          string             `json:"address"`
	Phone            string             `json:"phone"`
	EmergencyContact []EmergencyContact `json:"ContactDetail"`
}

func registerUser(c *gin.Context) {
	var input Person
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(input)
	c.JSON(http.StatusOK, "Registered Succefully!")
}

func accident(c *gin.Context) {
	var input Person
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(input)
	c.JSON(http.StatusOK, "accident happened Succefully!")
}

func main() {
	router := gin.Default()
	// router.GET("/albums", getAlbums)
	router.POST("/register", registerUser)
	router.POST("/accident", accident)

	router.Run("localhost:8080")
}
