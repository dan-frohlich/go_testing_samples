package service

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"strings"
)

//RobotFilter produce a filtered view of robots.txt
type RobotFilter interface {
	FilteredRobots(url string, allowed bool) ([]string, error)
}

type rf struct{}

//FilteredRobots - fun
func (rf *rf) FilteredRobots(url string, allowed bool) ([]string, error) {
	var resp *http.Response
	var err error

	var robotURL string
	if strings.HasSuffix(url, "/robots.txt") {
		robotURL = url
	} else {
		robotURL = url + "/robots.txt"
	}

	req, err := http.NewRequest("GET", robotURL, nil)
	if err != nil {
		log.Fatalln(err)
	}
	ua := `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36`
	// ua := `Golang_Spider_Bot/3.0`
	req.Header.Set("User-Agent", ua)
	req.Header.Set("Accept", "text/plain")

	client := &http.Client{}

	fmt.Printf("url: %+v\n", req)
	resp, err = client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	if err != nil {
		return nil, err
	}
	fmt.Printf("content length %d\n", resp.ContentLength)
	defer resp.Body.Close()
	r := bufio.NewReader(resp.Body)
	result := make([]string, 0)
	var predicate func(string) bool
	if allowed {
		predicate = func(line string) bool { return strings.HasPrefix(line, "Allow:") }
	} else {
		predicate = func(line string) bool { return strings.HasPrefix(line, "Disallow:") }
	}
	var line string
	for ; err != nil; line, err = r.ReadString('\n') {
		fmt.Println(line)
		if predicate(line) {
			result = append(result, line)
		}
	}
	fmt.Printf("%+v\n", resp.Header)

	return result, nil
}

//NewRobotFilter - foo
func NewRobotFilter() RobotFilter {
	return &rf{}
}
