package providers

import (
	"errors"
	"fmt"
	"os"

	"github.com/nlopes/slack"
	"github.com/pandelisz/serverless-go-forms/lib/types"
)

type slackProvider struct {
}

var Slack = slackProvider{}

// Send sends the message via slack
func (provider *slackProvider) Send(msg types.ContactBasic) error {

	var token, channel string

	if val, ok := os.LookupEnv("SLACK_TOKEN"); ok {
		token = val
	} else {
		return errors.New("SLACK_TOKEN not set")
	}
	if val, ok := os.LookupEnv("SLACK_CHANNEL_ID"); ok {
		channel = val
	} else {
		return errors.New("SLACK_CHANNEL_ID not set")
	}

	api := slack.New(token)
	params := slack.PostMessageParameters{}
	_, _, err := api.PostMessage(channel, fmt.Sprintf("*Sender*: %s\n*Referer*: %s\n*Message*:\n```%s```", msg.From, msg.Referer, msg.Message), params)
	if err != nil {
		fmt.Printf("%s\n", err)
		return err
	}

	return nil

}
