package main

import (
	"os"
	"sms-app/twilio"
)

func init() {
	//setting env variables
	os.Setenv("TO", "+923110030006")
	os.Setenv("BODY", "This is my test sms body")
}

func main() {
	//getting env variables
	to := os.Getenv("TO")
	body := os.Getenv("BODY")

	sms := twilio.SMS{
		To:   to,
		Body: body,
	}
	response := twilio.SendSMS(sms)
	println(response)
}
