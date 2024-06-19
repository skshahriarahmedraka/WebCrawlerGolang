package utils

import (
	"github.com/k0kubun/pp/v3"
	"github.com/tebeka/selenium"
)

func FetchRecommendedRate(webDriver selenium.WebDriver) string {
	cssPathRecommendedRate := ".BVRRBuyAgainPercentage > span:nth-child(1)"

	rating, err := webDriver.FindElement(selenium.ByCSSSelector, cssPathRecommendedRate)
	if err != nil {
		pp.Println("🚀 ~ file: title.go ~ line 12 ~ funcFetchTitle ~ err : ", err)
		return ""
	}
	text, err := rating.Text()
	if err != nil {
		pp.Println("🚀 ~ file: title.go ~ line 17 ~ funcFetchTitle ~ err : ", err)
		return ""
	}
	return text 

}
