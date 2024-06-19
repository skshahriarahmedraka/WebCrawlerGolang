package utils

import (
	"github.com/k0kubun/pp/v3"
	"github.com/tebeka/selenium"
)

func FetchRating(webDriver selenium.WebDriver) string {

	rating := ""
	ratingElm, err := webDriver.FindElement(selenium.ByCSSSelector, "span.BVRRNumber.BVRRRatingNumber")
	if err != nil {
		pp.Println("🚀 ~ file: rating.go ~ line 13 ~ funcFetchRating ~ err : ", err)
		return ""
	}
	rating, _ = ratingElm.Text()
	return rating
}
