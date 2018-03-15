package providers

import (
	"errors"
	"fmt"
	"os"

	"github.com/pandelisz/serverless-go-forms/lib/types"
	mailgun "gopkg.in/mailgun/mailgun-go.v1"
)

// Mailgun contains all mailgun related sending
type mailgunProvider struct {
}

var Mailgun = mailgunProvider{}

// Send sends the message via mailgun
func (provider *mailgunProvider) Send(msg types.ContactBasic) error {

	var domain, APIKey, publicAPIKey, sender, recipient string

	if val, ok := os.LookupEnv("MG_DOMAIN"); ok {
		domain = val
	} else {
		return errors.New("MG_DOMAIN not set")
	}
	if val, ok := os.LookupEnv("MG_API_KEY"); ok {
		APIKey = val
	} else {
		return errors.New("MG_DOMAIN not set")
	}
	if val, ok := os.LookupEnv("MG_PUBLIC_API_KEY"); ok {
		publicAPIKey = val
	} else {
		return errors.New("MG_DOMAIN not set")
	}
	if val, ok := os.LookupEnv("MG_SENDER"); ok {
		sender = val
	} else {
		return errors.New("MG_DOMAIN not set")
	}
	if val, ok := os.LookupEnv("MG_RECIPIENT"); ok {
		recipient = val
	} else {
		return errors.New("MG_DOMAIN not set")
	}

	mg := mailgun.NewMailgun(domain, APIKey, publicAPIKey)
	message := mg.NewMessage(
		sender,
		"New contact form message",
		fmt.Sprintf("Sender: %s\nReferer: %s\nMessage:\n\n%s", msg.From, msg.Referer, msg.Message),
		recipient)
	_, _, err := mg.Send(message)
	if err != nil {
		fmt.Print(err)
		return err
	}
	return nil

}
