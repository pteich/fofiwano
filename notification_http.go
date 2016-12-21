package fofiwano

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/mitchellh/mapstructure"
	"github.com/pteich/go-timeout-httpclient"
)

type NotifyHTTPOptions struct {
	URI        string `mapstructure:"uri"`
	Method     string `mapstructure:"method"`
	ParamEvent string `mapstructure:"param_event"`
	ParamPath  string `mapstructure:"param_path"`
}

// NotifyHTTP sends a file change notification to HTTP endpoint
func NotifyHTTP(options interface{}, event string, path string) error {

	httpClient := timeouthttp.NewClient(timeouthttp.Config{
		ConnectTimeout: 5,
		RequestTimeout: 5,
	})

	httpOptions := NotifyHTTPOptions{
		Method:     "GET",
		ParamPath:  "path",
		ParamEvent: "event",
	}

	if err := mapstructure.Decode(options, &httpOptions); err != nil {
		return err
	}

	if httpOptions.URI == "" {
		return errors.New("HTTP URI missing")
	}

	req, err := http.NewRequest(httpOptions.Method, httpOptions.URI, nil)
	if err != nil {
		return err
	}

	if httpOptions.Method == "GET" {
		q := req.URL.Query()
		q.Add(httpOptions.ParamEvent, event)
		q.Add(httpOptions.ParamPath, path)

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

	log.Printf("Gepusht an %s\n", req.URL.String())

	if body, err := ioutil.ReadAll(resp.Body); err == nil {
		_ = body
	}

	return nil
}
