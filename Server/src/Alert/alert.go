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
	"time"
)

var mp map[string]bool

type email struct {
	Tomail    string `json:"to_email"`
	Fromemail string `json:"from_email"`
	Subject   string `json:"subject"`
	Message   string `json:"message"`
}

func SendAlert(email string) bool {
	//init chan
	resp := false
	for _, j := range person.Users {

		if j.Email == email {
			fmt.Println(j)
			resp = sendEmailAllert(j.EmergencyContact[0].Email)
		}
	}
	return resp
}

func abortEmailAlert(u string) bool {
	mp[u] = false
	fmt.Println(mp)
	return true
}

func AbortAlert(email string) bool {
	//init chan
	res := false
	for _, j := range person.Users {

		if j.Email == email {
			fmt.Println(j)
			abortEmailAlert(j.EmergencyContact[0].Email)
			res = true
		}
	}
	return res
}

func sendEmailAllert(emailid string) bool {

	data := email{
		Tomail:    emailid,
		Fromemail: "admin@lazycoderz.com",
		Subject:   "Accident Alert!",
		Message:   "Help me!",
	}
	dataenc, err := json.Marshal(data)
	if err != nil {
		fmt.Println("error occured", err.Error())
		return false
	}

	mp = make(map[string]bool)
	mp[emailid] = true
	fmt.Println(mp)
	time.Sleep(5 * time.Second)
	var response *http.Response
	var resp bool
	if v := mp[emailid]; v {
		fmt.Println("send email")
		response, err = http.Post("http://lazycoderz.com/phpmailer.php", "application/json", bytes.NewBuffer(dataenc))
		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}
		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(responseData))
		resp = true
	} else {
		fmt.Println("Aborted send email")
		resp = false
	}
	delete(mp, emailid)

	return resp
}
