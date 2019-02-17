package notification

import (
	"errors"
	"log"

	"github.com/mitchellh/mapstructure"
	"github.com/pteich/slackstatus"
)

type Slack struct {
	slackstatus.Message
	WebhookURL string `mapstructure:"webhook_url"`
	Channel    string `mapstructure:"channel"`
	Username   string `mapstructure:"username"`
	IconEmoji  string `mapstructure:"icon_emoji"`
	Footer     string `mapstructure:"footer"`
}

// Slack.Notify sends a file change notification to Slack
func (slacknotifier *Slack) Notify(event string, path string) error {

	// TODO add templates for better control over messages
	err := slacknotifier.Send(event+" - "+path, slackstatus.COLOR_GOOD)
	if err == nil {
		log.Printf("Event %s for %s pushed to Slack channel %s\n", event, path, slacknotifier.Channel)
	}
	return err
}

func NewSlackNotification(options Options) (*Slack, error) {

	slacknotifier := new(Slack)

	if err := mapstructure.Decode(options, &slacknotifier); err != nil {
		return nil, err
	}

	if slacknotifier.WebhookURL == "" {
		return nil, errors.New("Slack webhook missing")
	}

	slacknotifier.Message.WebhookURL = slacknotifier.WebhookURL
	slacknotifier.Message.Username = slacknotifier.Username
	slacknotifier.Message.Channel = slacknotifier.Channel
	slacknotifier.Message.IconEmoji = slacknotifier.IconEmoji
	slacknotifier.Message.Footer = slacknotifier.Footer

	return slacknotifier, nil
}
