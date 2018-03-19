package service

import (
	"strings"
	"testing"
)

func Test_service_returns_mashup(t *testing.T) {

	var cmock Catalog
	var tmock TweetAPI
	// mocker := gomock.NewController(t)
	// cmock := NewMockCatalog(mocker)
	// tmock := NewMockTweetAPI(mocker)

	service := NewService(cmock, tmock)
	query := "Good Bars"
	mashUps := service.Mashup(query)

	if len(mashUps) != 3 {
		t.Errorf("expected 3 mashups but got [%d]", len(mashUps))
	}

	expectedProductID := "0001"
	expectedProductName := "Foo Bars"

	foorBarMashUp := mashUps[0]
	actualProductID := foorBarMashUp.Product.ID
	if actualProductID != expectedProductID {
		t.Errorf("expected [%s] but got [%s]", expectedProductID, actualProductID)
	}

	actualProductName := foorBarMashUp.Product.Name
	if actualProductName != expectedProductName {
		t.Errorf("expected [%s] but got [%s]", expectedProductName, actualProductName)
	}

	if len(foorBarMashUp.Tweets) != 1 {
		t.Errorf("expected 1 mashups but got [%d]", len(foorBarMashUp.Tweets))
	}

	actualTweetText := foorBarMashUp.Tweets[0].Text
	if !strings.Contains(actualTweetText, expectedProductName) {
		t.Errorf("expect a tweet with [%s] but got [%s]", expectedProductName, actualTweetText)
	}

}
