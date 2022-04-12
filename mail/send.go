package mail

import (
	"errors"
	"net/mail"
	"net/smtp"
	"os"
	"strconv"

	"github.com/fldu/unisender/utils"
)

/* TODO:
- Attachment
- Unit tests (https://github.com/mocktools/go-smtp-mock?)
*/

func SendNotification(c utils.Config) error {
	server := c.Email.SMTPServer + strconv.Itoa(c.Email.SMTPPort)
	auth := smtp.PlainAuth("", c.Email.SMTPUsername, c.Email.SMTPPassword, c.Email.SMTPServer)

	err := validateEmailAddr(c.Email.From, c.Email.To)
	if err != nil {
		return errors.New(err.Error())
	}

	// Try to interpret the body parameter as a file. If it fails, writing body parameter as the body
	body, err := os.ReadFile(c.Email.Body)
	if err != nil {
		body = []byte(c.Email.Body)
	}

	err = smtp.SendMail(
		server,
		auth,
		c.Email.From,
		c.Email.To,
		body,
	)
	if err != nil {
		return errors.New("Failed to send email notification: " + err.Error())
	}
	return nil
}

func validateEmailAddr(f string, r []string) error {
	r = append(r, f)
	for _, j := range r {
		_, err := mail.ParseAddress(j)
		if err != nil {
			return errors.New("Email address: " + j + " is not compliant, aborting.")
		}
	}
	return nil
}
