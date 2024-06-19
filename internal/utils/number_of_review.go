package utils

import (
	"github.com/k0kubun/pp/v3"
	"github.com/tebeka/selenium"
)

func FetchNumberOfReviews(webDriver selenium.WebDriver) string {
	// 	cssPathProductName := "html body div div div div main div div div div div div h1.itemTitle.test-itemTitle"

	numOfReview, err := webDriver.FindElement(selenium.ByCSSSelector, "span.BVRRNumber.BVRRBuyAgainTotal")
	if err != nil {
		pp.Println("🚀 ~ file: number_of_review.go ~ line 13 ~ funcFetchNumberOfReviews ~ err : ", err)
		return ""
	}
	text, err := numOfReview.Text()
	if err != nil {
		pp.Println("🚀 ~ file: number_of_review.go ~ line 18 ~ funcFetchNumberOfReviews ~ err : ", err)
		return ""
	}
	return text

}
