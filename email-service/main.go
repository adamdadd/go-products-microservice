package email_service

import (
	"log"
	"net/smtp"
)

func main () {
	auth := smtp.PlainAuth("","","", "")

	to := []string{"bueller@home.net"}
	msg := []byte("To user@email\r\n" +
		"Subject: Replace this dynamically" +
		"\r\n"+
		"Offers")
	err := smtp.SendMail("", auth, "", to, msg)
	if err != nil {
		log.Fatal(err)
	}
}
