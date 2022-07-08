package alert

import (
	"bytes"
	"encoding/json"
	"fmt"
	"hackathon/server/src/person"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func SendAlert(email string) error {
	//init chan
	for _, j := range person.Users {

		if j.Email == email {
			fmt.Println(j)
			sendEmailAllert(j.EmergencyContact[0].Email)
		}
	}
	return fmt.Errorf("can't find person detail, person is not registered with us")
}

// cf45291ef134a27d9bf2f4ebb0b568eb
// AC02ed0937bb6cd586beff558d0e3a6026
// func sendWhatsappMessage(j person.Person) {
// // time.Sleep(10 * time.Second)
// from := os.Getenv("+912121212121")
// to := os.Getenv("+918299661294")
// fmt.Println("swm1")
// client := twilio.NewRestClient()

// params := &openapi.CreateMessageParams{}
// params.SetTo(to)
// params.SetFrom(from)
// params.SetBody("Hello there")
// fmt.Println("swm2")
// resp, err := client.Api.CreateMessage(params)
// if err != nil {
// 	fmt.Println("swm2.err")
// 	fmt.Println(err.Error())
// } else {
// 	fmt.Println("swm2.else")
// 	response, _ := json.Marshal(*resp)
// 	fmt.Println("Response: " + string(response))
// }

//wait for 15 sec
//if abort quit
//else send.
// }

func AbortAlert(u person.Person) {

}

type email struct {
	Tomail    string `json:"to_email"`
	Fromemail string `json:"from_email"`
	Subject   string `json:"subject"`
	Message   string `json:"message"`
}

func sendEmailAllert(emailid string) {
	data := email{
		Tomail:    emailid,
		Fromemail: "admin@lazycoderz.com",
		Subject:   "Accident Alert!",
		Message:   "Help me!",
	}
	dataenc, err := json.Marshal(data)
	if err != nil {
		fmt.Println("error occured", err.Error())
		return
	}
	response, err := http.Post("http://lazycoderz.com/phpmailer.php", "application/json", bytes.NewBuffer(dataenc))

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
