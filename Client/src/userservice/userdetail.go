package userservice

import "fmt"

type User struct {
	Id               string             `gorm:"primary_key" json:"id"`
	Email            string             `gorm:"unique" json:"email"`
	Name             string             `json:"name"`
	Address          string             `json:"address"`
	Phone            string             `json:"phone"`
	EmergencyContact []EmergencyContact `json:"ContactDetail"`
}

type EmergencyContact struct {
	Phone string `json:"phone"`
	Email string `json:"email"`
}

var UserDetail User

func RegisterUser() User {

	fmt.Println("Enter your name")
	fmt.Scanln(&UserDetail.Name)
	fmt.Println("Enter your emailid")
	fmt.Scanln(&UserDetail.Email)
	fmt.Println("Enter your address")
	fmt.Scanln(&UserDetail.Address)
	fmt.Println("Enter your phone Number")
	fmt.Scanln(&UserDetail.Phone)
	AddContacts()
	return UserDetail
}

func AddContacts() User {
	if UserDetail.Name == "" {
		fmt.Println("Your are not yet register. Register first!\n Redirecting for registration.")
		return RegisterUser()

	}
	var contactDetails []EmergencyContact
	for len(UserDetail.EmergencyContact) < 5 {
		var ec EmergencyContact
		fmt.Println("Please enter emergency contact phone number")
		fmt.Scanln(&ec.Phone)
		fmt.Println("Please enter emergency contact email")
		fmt.Scanln(&ec.Email)
		contactDetails = append(contactDetails, ec)
		if len(UserDetail.EmergencyContact) < 5 {
			fmt.Println("Do you want to add more Emergency contacts? Yes/No")
			var resp string
			fmt.Scanln(&resp)
			if resp[0] != 'Y' && resp[0] != 'y' {
				break
			}
		} else {
			fmt.Println("Your limit to add user is full")
		}
	}
	UserDetail.EmergencyContact = contactDetails
	return UserDetail
}
