// Package gletter is a simple lib to send e-mails using Golang SMTP
package gletter

import (
	"log"
	"net/smtp"
	"strconv"
)

// MailConfig set the basic config to send
// Server:        SMTP Server. Eg: smtp.server.com
// Port:          SMTP Server port. Eg: 25
// MailSender:    Who is sending the e-mail. Eg: my@email.com
// MailReceiver:  Who is going to get the e-mail
// MailPassword:  Password of MailSender
// Message:       Set e-mail message
// Subject:       Message Subject
type MailConfig struct {
	Server, MailSender, MailPassword, MailReceiver, Subject, Message string
	Port                                                             int
}

// SetConfig store the configs in the Struct to send by SendMail function
func (Config *MailConfig) SetConfig(s, mp, ms, mr, sbj string, p int, msg string) {

	Config.Server = s        // Set server
	Config.MailPassword = mp // Set password
	Config.MailSender = ms   // Set who is sending
	Config.MailReceiver = mr // Set who is receiving
	Config.Port = p          // Set the SMTP port
	Config.Message = msg     // Set the message to be send
	Config.Subject = sbj     // Set the message subject

}

// SendMail get the parameters by the user, put in the struct and use SetMail to send
func SendMail(s, mp, ms, mr, sbj string, p int, msg string) {

	// Load the Struct in Config variable
	Config := MailConfig{}

	// Get the values and send to the Struct
	Config.SetConfig(s, mp, ms, mr, sbj, p, msg)

	// Setup the e-mail to be ready to send
	SetMail(Config)

}

// SetMail use the parameters on the Struct to send e-mails
func SetMail(Config MailConfig) {

	// Get the Auth
	mauth := smtp.PlainAuth(
		"",
		Config.MailSender,   // Sender e-mail
		Config.MailPassword, // Sender password
		Config.Server,       // Server. Eg: smtp.google.com
	)

	emailFrom := []string{Config.MailReceiver}
	emailHeader := "To:" + Config.MailReceiver + "\r\nSubject:" + Config.Subject + "\r\nMIME-Version: 1.0\r\nContent-Type: text/html; charset=ISO-8891-1\r\n\r\n"

	emailBody := []byte(emailHeader + Config.Message)
	// Send the e-mail
	err := smtp.SendMail(
		Config.Server+":"+strconv.Itoa(Config.Port), // Server + Port
		mauth,             // Get the Auth setup
		Config.MailSender, // Get who is sending the e-mail
		emailFrom,         // Get who will receive the e-mail
		emailBody,         // Get the message
	)

	if err != nil {
		log.Fatal(err) // Log if error
	} else {
		log.Println("E-Mail send to: " + Config.MailReceiver) // Log if succeful and display who receive the e-mail
	}

}
