package service

import (
	"strings"
	"testing"
)

func Test_service_returns_mashup(t *testing.T) {

	var tmock TweetAPI
	var cmock Catalog

	cmock = &CatalogMock{
		search: func(query string) []SearchHit {
			return []SearchHit{
				SearchHit{ProductID: "0001"},
				SearchHit{ProductID: "0002"},
				SearchHit{ProductID: "0003"},
			}
		},
		get: func(id string) Product {
			var name string
			switch id {
			case "0001":
				name = "Foo Bars"
			default:
				name = "Other Bars"
			}
			return Product{Name: name, ID: id}
		},
	}

	tmock = &TweetAPIMock{
		query: func(query string) []Tweet {
			return []Tweet{Tweet{Text: "Foo Bars are the best!"}}
		},
	}

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
