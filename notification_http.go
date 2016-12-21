package fofiwano

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/mitchellh/mapstructure"
	"github.com/pteich/go-timeout-httpclient"
)

type HTTP struct {
	URL        string `mapstructure:"url"`
	Method     string `mapstructure:"method"`
	ParamEvent string `mapstructure:"param_event"`
	ParamPath  string `mapstructure:"param_path"`
}

// NotifyHTTP sends a file change notification to HTTP endpoint
func (notifier *HTTP) Notify(event string, path string) error {

	httpClient := timeouthttp.NewClient(timeouthttp.Config{
		RequestTimeout: 5,
		ConnectTimeout: 5,
	})

	req, err := http.NewRequest(notifier.Method, notifier.URL, nil)
	if err != nil {
		return err
	}

	if notifier.Method == "GET" {
		q := req.URL.Query()
		q.Add(notifier.ParamEvent, event)
		q.Add(notifier.ParamPath, path)

		req.URL.RawQuery = q.Encode()
	}

	resp, err := httpClient.Do(req)
	if resp == nil {
		return errors.New("error reading response from " + req.URL.String())
	}
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	log.Printf("Event %s for %s pushed to %s\n", event, path, req.URL.String())

	if body, err := ioutil.ReadAll(resp.Body); err == nil {
		_ = body
	}

	return nil
}

func NewHTTPNotification(options interface{}) (*HTTP, error) {

	httpNotifier := &HTTP{
		Method:     "GET",
		ParamPath:  "path",
		ParamEvent: "event",
	}

	if err := mapstructure.Decode(options, &httpNotifier); err != nil {
		return nil, err
	}

	if httpNotifier.URL == "" {
		return nil, errors.New("HTTP URL missing")
	}

	return httpNotifier, nil
}
