package main

import (
	"log"
	"net/smtp"
	"os"
)

func main() {
	sendGmail(
		"alice@example.com",
		"test",
		"This mail was sent by golang.",
	)
}

func sendGmail(to string, subject string, body string) {
	user := os.Getenv("GMAIL_USER")
	password := os.Getenv("GMAIL_PASSWORD")
	name := os.Getenv("MAIL_DISPLAY_NAME")
	from := os.Getenv("MAIL_FROM")

	auth := smtp.PlainAuth("", user, password, "smtp.gmail.com")

	msg := []byte("" +
		"From: " + name + " <" + from + ">\r\n" +
		"To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		body + "\r\n")

	if err := smtp.SendMail("smtp.gmail.com:587", auth, from, []string{to}, msg); err != nil {
		log.Fatalln("smtp error:", err)
	}
	log.Println("success")
}
