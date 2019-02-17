package notification

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewHTTPNotificationGet(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		assert.Equal(t, "file", r.URL.Query().Get("event"))
		assert.Equal(t, "testfile.txt", r.URL.Query().Get("path"))

		fmt.Fprintln(w, "Ok")
	}))
	defer ts.Close()

	httpOptions := Options{"method": "get", "url": ts.URL}

	testNotification, err := NewHTTPNotification(httpOptions)
	assert.NoError(t, err)

	assert.Equal(t, http.MethodGet, testNotification.Method)
	assert.Equal(t, httpOptions["url"], testNotification.URL)

	err = testNotification.Notify("file", "testfile.txt")
	assert.NoError(t, err)
}

func TestNewHTTPNotificationPost(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		payload := &HTTPPaylod{}

		defer r.Body.Close()

		err := json.NewDecoder(r.Body).Decode(payload)
		assert.NoError(t, err)

		assert.Equal(t, "file", payload.Event)
		assert.Equal(t, "testfile.txt", payload.Path)

		fmt.Fprintln(w, "Ok")
	}))
	defer ts.Close()

	httpOptions := Options{"method": "post", "url": ts.URL}

	testNotification, err := NewHTTPNotification(httpOptions)
	assert.NoError(t, err)

	assert.Equal(t, http.MethodPost, testNotification.Method)
	assert.Equal(t, httpOptions["url"], testNotification.URL)

	err = testNotification.Notify("file", "testfile.txt")
	assert.NoError(t, err)
}
