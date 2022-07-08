package endpoints

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	user "hackathon/client/src/userservice"
)

func RegisterOnline(data user.User) {
	fmt.Println(data)
	dataenc, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err.Error())
	}

	response, err := http.Post("http://localhost:8080/register", "application/json", bytes.NewBuffer(dataenc))

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseData))
}

type Email struct {
	Email string `json:"email"`
}

func SendAllert(email string) {
	dataenc, err := json.Marshal(Email{
		Email: email,
	})
	if err != nil {
		fmt.Println("Error occured", err.Error())
	}
	fmt.Println("Allert send start")
	response, err := http.Post("http://localhost:8080/allert", "application/json", bytes.NewBuffer(dataenc))

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	fmt.Println("error send end")
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseData))
}

func UpdateContacts(data user.User) {
	fmt.Println(data)
	dataenc, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err.Error())
	}

	response, err := http.Post("http://localhost:8080/register", "application/json", bytes.NewBuffer(dataenc))

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseData))
}
