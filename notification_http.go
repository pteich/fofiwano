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
	URI        string `mapstructure:"uri"`
	Method     string `mapstructure:"method"`
	ParamEvent string `mapstructure:"param_event"`
	ParamPath  string `mapstructure:"param_path"`
	httpClient *http.Client
}

// NotifyHTTP sends a file change notification to HTTP endpoint
func (notifier *HTTP) Notify(event string, path string) error {

	req, err := http.NewRequest(notifier.Method, notifier.URI, nil)
	if err != nil {
		return err
	}

	if notifier.Method == "GET" {
		q := req.URL.Query()
		q.Add(notifier.ParamEvent, event)
		q.Add(notifier.ParamPath, path)

		req.URL.RawQuery = q.Encode()
	}

	resp, err := notifier.httpClient.Do(req)
	if resp == nil {
		return errors.New("error reading response from " + req.URL.String())
	}
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	log.Printf("Gepusht an %s\n", req.URL.String())

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

	httpNotifier.httpClient = timeouthttp.NewClient(timeouthttp.Config{
		ConnectTimeout: 5,
		RequestTimeout: 5,
	})

	if err := mapstructure.Decode(options, &httpNotifier); err != nil {
		return nil, err
	}

	if httpNotifier.URI == "" {
		return nil, errors.New("HTTP URI missing")
	}

	return httpNotifier, nil
}
