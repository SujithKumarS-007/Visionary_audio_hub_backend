package service

import (
	"log"
	"github.com/sfreiberg/gotwilio"
)

func SendSMSforVerification(mobileno, otp string) {
    log.Println("Inside send SMS................")
	// Twilio Account SID and Auth Token
	accountSid := "AC8ac8382f79ce8e5e425cdb0bc04171a9"
	authToken := "f494c021283c29d2a94e85a19838a300"

	// Create a Twilio client
	client := gotwilio.NewTwilioClient(accountSid, authToken)

	// Sender's phone number (Twilio number)
	from := "+16592014360"

	// Recipient's phone number
	to := "+91" + mobileno

	// Message to send
	opt_store := "We are Sending the SMS for verfication of your Mobile No and Email.Please use this OTP to complete the verification process.If you did not request this verification, please ignore this." + otp
	message := opt_store
	// Send the message
	resp, ex, err := client.SendSMS(from, to, message, "", "")
	if err != nil {
		log.Printf("Error sending SMS: %s", err)
		return
	}
	if ex != nil {
		log.Printf("Exception sending SMS: %s", ex.Message)
		return
	}
	log.Printf("SMS Sent! SID: %s, Status: %s", resp.Sid, resp.Status)
}
