// +build integration google_integration

package project05

import (
	"net/http"
	"project05/service"
	"testing"
)

func Test_website_is_up_google(t *testing.T) {
	destination := "http://www.google.com/robots.txt"
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

func Test_google_RobotFilter(t *testing.T) {
	f := service.NewRobotFilter()
	if allowed, err := f.FilteredRobots("http://www.google.com", true); err != nil {
		t.Errorf("filter had error: %+v", err)
	} else if len(allowed) < 1 {
		t.Errorf("found no allowed entries.")
	}

}
