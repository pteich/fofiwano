package fofiwano

import (
	"errors"

	"github.com/mitchellh/mapstructure"
	"github.com/pteich/slackstatus"
)

type Slack struct {
	slackstatus.Message
}

// Slack.Notify sends a file change notification to Slack
func (slacknotifier *Slack) Notify(event string, path string) error {

	// TODO add templates for better control over messages
	return slacknotifier.Send(event+" - "+path, slackstatus.ColorGood)
}

func NewSlackNotification(options interface{}) (*Slack, error) {

	slacknotifier := new(Slack)

	if err := mapstructure.Decode(options, &slacknotifier.Message); err != nil {
		return nil, err
	}

	if slacknotifier.WebhookURL == "" {
		return nil, errors.New("Slack webhook missing")
	}

	return slacknotifier, nil
}
