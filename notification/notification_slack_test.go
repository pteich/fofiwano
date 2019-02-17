package notification

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSlackNotification(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		body, err := ioutil.ReadAll(r.Body)
		assert.NoError(t, err)
		defer r.Body.Close()

		assert.Equal(t, []byte(`{"attachments":[{"fallback":null,"color":"good","pretext":null,"author_name":null,"author_link":null,"author_icon":null,"title":null,"title_link":null,"text":"file - testfile.txt","image_url":null,"fields":null,"footer":"","footer_icon":null}]}`), body)

		fmt.Fprintln(w, "Ok")
	}))
	defer ts.Close()

	slackOptions := Options{"webhook_url": ts.URL}

	testNotification, err := NewSlackNotification(slackOptions)
	assert.NoError(t, err)

	err = testNotification.Notify("file", "testfile.txt")
	assert.NoError(t, err)

}
