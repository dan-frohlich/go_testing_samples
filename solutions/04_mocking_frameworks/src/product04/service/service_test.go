package service

import (
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
)

func Test_service_returns_mashup(t *testing.T) {

	// var cmock Catalog
	// var tmock TweetAPI
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	cmock := NewMockCatalog(mockCtrl)
	tmock := NewMockTweetAPI(mockCtrl)

	cmock.EXPECT().Search(gomock.Any()).Return(
		[]SearchHit{
			SearchHit{ProductID: "0001"},
			SearchHit{ProductID: "0002"},
			SearchHit{ProductID: "0003"},
		})

	cmock.EXPECT().Get("0001").Return(Product{ID: "0001", Name: "Foo Bars"}).MinTimes(0)
	cmock.EXPECT().Get("0002").Return(Product{ID: "0002", Name: "Other Bars"}).MinTimes(0)
	cmock.EXPECT().Get("0003").Return(Product{ID: "0003", Name: "Other Bars"}).MinTimes(0)
	tmock.EXPECT().Query(gomock.Any()).Return([]Tweet{Tweet{Text: "Foo Bars are great!"}}).MinTimes(0)

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
