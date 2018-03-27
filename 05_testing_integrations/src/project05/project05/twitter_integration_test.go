// +build integration twitter_integration

package project05

import (
	"net/http"
	"testing"
)

func Test_website_is_up_twitter(t *testing.T) {
	destination := "https://twitter.com/robots.txt"
	resp, err := http.Get(destination)
	defer resp.Body.Close()

	if err != nil {
		t.Errorf("GET %s got [%+v]", destination, err)
	} else {
		if resp.StatusCode < 200 || resp.StatusCode > 299 {
			t.Errorf("GET %s got %d - %v", destination, resp.StatusCode, resp.Status)
		}
	}
}
