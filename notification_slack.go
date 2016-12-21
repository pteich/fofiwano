package fofiwano

import (
	"errors"

	"github.com/mitchellh/mapstructure"
	"github.com/pteich/slackstatus"
	"log"
)

type Slack struct {
	slackstatus.Message
}

// Slack.Notify sends a file change notification to Slack
func (slacknotifier *Slack) Notify(event string, path string) error {

	// TODO add templates for better control over messages
	err := slacknotifier.Send(event+" - "+path, slackstatus.ColorGood)
	if err == nil {
		log.Printf("Event %s for %s pushed to Slack channel %s\n", event, path, slacknotifier.Channel)
	}
	return err
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
