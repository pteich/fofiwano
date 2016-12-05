package fofiwano

import (
	"errors"

	"github.com/mitchellh/mapstructure"
	"github.com/pteich/slackstatus"
)

func NotifySlack(options interface{}, event string, path string) error {

	var slackmsg slackstatus.Message
	if err := mapstructure.Decode(options, &slackmsg); err != nil {
		return err
	}

	if slackmsg.WebhookURL == "" {
		return errors.New("Slack webhook missing")
	}

	return slackmsg.Send(event+" - "+path, slackstatus.ColorGood)
}
