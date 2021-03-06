package main

import (
	"fmt"

	server "hackathon/client/src/endpoints"
	user "hackathon/client/src/userservice"
)

type Email struct {
	Email string `json:"email"`
}

func main() {
	fmt.Println("Welcome to Accident monitoring system")
	for selectOperation() {

	}
}

func selectOperation() bool {
	fmt.Println("Select the operation you want to perform")
	fmt.Println("1.Register User")
	// fmt.Println("2.Add Contact")
	fmt.Println("2.Send Signal")
	fmt.Println("3.Abort Signal")
	fmt.Println("4. Exit!")
	var i int
	fmt.Println("Enter number according to your choice of operation")
	fmt.Scanln(&i)
	switch i {
	case 1:
		user := user.RegisterUser()
		server.RegisterOnline(user)
	case 2:
		fmt.Println("enter your emailid")
		var email string
		fmt.Scanln(&email)
		server.SendAllert(email)
	case 3:
		fmt.Println("enter your emailid")
		var email string
		fmt.Scanln(&email)
		server.Abort(email)
	case 4:
		return false
	default:
		fmt.Println("selected Wrong Choice, Try again")
		fmt.Println("Do you exit")
		var resp string
		fmt.Scanln(&resp)
		if resp[0] != 'Y' && resp[0] != 'y' {
			return true
		}
		return false
	}

	return true
}
