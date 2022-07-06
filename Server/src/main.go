package main

import (
	alert "hackathon/server/src/Alert"
	"hackathon/server/src/person"
	"net/http"

	"github.com/gin-gonic/gin"
)

type personContact struct {
	IdEmail string `gorm:"unique" json:"idemail"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
}

var Users []person.Person

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
	for _, j := range Users {

		if j.Email == input.Email {
			c.JSON(http.StatusBadRequest, "User already exists!")
			return
		}
	}
	Users = append(Users, person.Person{
		Name:  input.Name,
		Email: input.Email,
		Phone: input.Phone,
	})
	c.JSON(http.StatusOK, "Registered Succefully!")
}

func addContact(c *gin.Context) {
	var input personContact
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var i int
	var j person.Person
	for i, j = range Users {
		if j.Email == input.IdEmail {
			j.EmergencyContact = append(j.EmergencyContact, person.EmergencyContact{
				Phone: input.Phone,
				Email: input.Email,
			})
			Users[i].EmergencyContact = j.EmergencyContact
		}
	}
	c.JSON(http.StatusOK, Users[i])
}

func accident(c *gin.Context) {
	var input person.Person
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, j := range Users {
		if j.Email == input.Email {
			alert.SendAlert(j)
		}
	}
	c.JSON(http.StatusOK, Users)
}

func abort(c *gin.Context) {
	var input person.Person
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, j := range Users {
		if j.Email == input.Email {
			alert.AbortAlert(j)
		}
	}
	c.JSON(http.StatusOK, "Alert aborted!")
}

func main() {
	router := gin.Default()
	// router.GET("/albums", getAlbums)
	router.POST("/register", registerUser)
	router.POST("/addcontact", addContact)
	router.GET("/accident", accident)
	router.GET("/abort", abort)

	router.Run("localhost:8080")
}
