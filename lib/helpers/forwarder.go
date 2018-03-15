package helpers

import (
	"fmt"
	"strings"

	"github.com/pandelisz/serverless-go-forms/lib/providers"
	"github.com/pandelisz/serverless-go-forms/lib/types"
)

type Forwarder struct {
	Mailgun bool
	Slack   bool
	SMTP    bool
	Twilio  bool
}

func (b *Forwarder) Send(payload types.ContactBasic) error {
	var arrErrors []string
	if b.Mailgun {
		err := providers.Mailgun.Send(payload)
		if err != nil {
			arrErrors = append(arrErrors, err.Error())
		}
	}

	if len(arrErrors) > 0 {
		return fmt.Errorf(strings.Join(arrErrors, ";"))
	}

	return nil

}
