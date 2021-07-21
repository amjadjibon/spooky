// Package email Sending Email Using Smtp in Golang
package email

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"net/mail"
	"net/smtp"
	"os"
	"strings"
)

// sendEmailSMTP Main function
func sendEmailSMTP(from, password, host, port, msg string, toList []string) {
	// We can't send strings directly in mail,
	// strings need to be converted into slice bytes
	body := []byte(msg)

	// PlainAuth uses the given username and password to
	// authenticate to host and act as identity.
	// Usually identity should be the empty string,
	// to act as username.
	auth := smtp.PlainAuth("", from, password, host)

	// SendMail uses TLS connection to send the mail
	// The email is sent to all address in the toList,
	// the body should be of type bytes, not strings
	// This returns error if any occurred.
	err := smtp.SendMail(host+":"+port, auth, from, toList, body)

	// handling the errors
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Successfully sent mail to all user in toList")
}

func isValid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func parseDomainNameFromEmail(email string) string {
	if !isValid(email) {
		return ""
	}
	components := strings.Split(email, "@")
	return components[1]
}

func smtpHost(domain string) string {
	switch domain {
	case "google.com":
		return "smtp.gmail.com"
	default:
		return ""
	}
}

func SendEmail(c *cli.Context) error {
	from := c.String("from")
	pass := c.String("pass")
	to := c.String("to")
	msg := c.String("message")
	host := "smtp.google.com"
	port := "587"

	var toList []string
	toList = append(toList, to)

	//fmt.Println(smtpHost(parseDomainNameFromEmail(to)))
	//
	//if smtpHost(to) == "google.com" {
	//	host = "smtp.google.com"
	//} else {
	//	return errors.New("not supported domain")
	//}

	sendEmailSMTP(from, pass, host, port, msg, toList)

	return nil
}