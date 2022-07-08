package main

import (
	"fmt"
	"os"

	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

func main() {
	from := os.Getenv("+91222112")
	to := os.Getenv("+918299661294")

	client := twilio.NewRestClient()

	params := &openapi.CreateCallParams{}
	params.SetTo(to)
	params.SetFrom(from)
	params.SetUrl("http://twimlets.com/holdmusic?Bucket=com.twilio.music.ambient")

	resp, err := client.Api.CreateCall(params)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Call Status: " + *resp.Status)
		fmt.Println("Call Sid: " + *resp.Sid)
		fmt.Println("Call Direction: " + *resp.Direction)
	}
}
