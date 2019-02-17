package notification

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"github.com/pteich/go-timeout-httpclient"
)

type HTTP struct {
	URL        string `mapstructure:"url"`
	Method     string `mapstructure:"method"`
	ParamEvent string `mapstructure:"param_event"`
	ParamPath  string `mapstructure:"param_path"`
}

type HTTPPaylod struct {
	Event string `json:"event"`
	Path  string `json:"path"`
}

// NotifyHTTP sends a file change notification to HTTP endpoint
func (notifier *HTTP) Notify(event string, path string) error {

	httpClient := timeouthttp.NewClient(timeouthttp.Config{
		RequestTimeout: 5,
		ConnectTimeout: 5,
	})

	body := &bytes.Buffer{}

	if notifier.Method == http.MethodPost {
		payload := HTTPPaylod{
			Event: event,
			Path:  path,
		}
		err := json.NewEncoder(body).Encode(payload)
		if err != nil {
			return err
		}
	}

	req, err := http.NewRequest(notifier.Method, notifier.URL, body)
	if err != nil {
		return err
	}

	if notifier.Method == http.MethodGet {
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

func NewHTTPNotification(options Options) (*HTTP, error) {

	httpNotifier := &HTTP{
		Method:     http.MethodGet,
		ParamPath:  "path",
		ParamEvent: "event",
	}

	if err := mapstructure.Decode(options, &httpNotifier); err != nil {
		return nil, err
	}

	if httpNotifier.URL == "" {
		return nil, errors.New("HTTP URL missing")
	}

	httpNotifier.Method = strings.ToUpper(httpNotifier.Method)

	return httpNotifier, nil
}
