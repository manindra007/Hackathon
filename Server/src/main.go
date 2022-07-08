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
	// person.Users = append(person.Users, person.Person{
	// 	Name:    input.Name,
	// 	Email:   input.Email,
	// 	Phone:   input.Phone,
	// 	Address: input.Address,
	// })
	person.Users = append(person.Users, input)
	fmt.Println(person.Users)
	c.JSON(http.StatusOK, "Registered Succefully!")
}

func addContact(c *gin.Context) {
	var input person.PersonContact
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var i int
	var j person.Person
	for i, j = range person.Users {
		if j.Email == input.IdEmail {
			j.EmergencyContact = append(j.EmergencyContact, person.EmergencyContact{
				Phone: input.Phone,
				Email: input.Email,
			})
			person.Users[i].EmergencyContact = j.EmergencyContact
		}
	}
	c.JSON(http.StatusOK, person.Users[i])
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

	alert.SendAlert(input.Email)
	c.JSON(http.StatusOK, person.Users)
}

func abort(c *gin.Context) {
	var input person.Person
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, j := range person.Users {
		if j.Email == input.Email {
			alert.AbortAlert(j)
		}
	}
	c.JSON(http.StatusOK, "Alert aborted!")
}

func main() {
	router := gin.Default()
	router.POST("/register", registerUser)
	router.POST("/addcontact", addContact)
	router.POST("/allert", allert)
	router.GET("/abort", abort)

	router.Run("localhost:8080")
}
