package main

import (
	"os"
	"sms-app/sendgrid"
	"sms-app/twilio"
)

func init() {
	//setting SMS env variables
	os.Setenv("TO", "+923110030006")
	os.Setenv("BODY", "This is my test sms body")

	//setting EMAIL env variables
	os.Setenv("EMAIL_FROM_NAME", "From Sample Name")
	os.Setenv("EMAIL_FROM", "no-reply@example.com")
	os.Setenv("EMAIL_TO_NAME", "Muhammad Fayyaz")
	os.Setenv("EMAIL_TO", "muhammad_fayyaz@live.com")
	os.Setenv("EMAIL_SUBJECT", "Sending with SendGrid is Fun")
	os.Setenv("EMAIL_BODY", "This is my test email content.")
}

func main() {
	//sendSMS()
	sendEmail()
}

func sendSMS() {
	sms := twilio.SMS{
		To:   os.Getenv("TO"),
		Body: os.Getenv("BODY"),
	}
	response := twilio.SendSMS(sms)
	println(response)
}

func sendEmail() {
	email := sendgrid.EMAIL{
		FromName: os.Getenv("EMAIL_FROM_NAME"),
		From:     os.Getenv("EMAIL_FROM"),
		ToName:   os.Getenv("EMAIL_TO_NAME"),
		To:       os.Getenv("EMAIL_TO"),
		Subject:  os.Getenv("EMAIL_SUBJECT"),
		Body:     os.Getenv("EMAIL_BODY"),
	}
	sendgrid.SendEmail(email)
}
