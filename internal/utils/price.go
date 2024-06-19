package utils

import (
	"github.com/k0kubun/pp/v3"
	"github.com/tebeka/selenium"
)

func FetchPrice(webDriver selenium.WebDriver) string {
// 	cssPathPricing := "html body div div div div main div div div div div div.articlePrice p.price-text span.price-value.test-price-value"

	price, err := webDriver.FindElement(selenium.ByCSSSelector, ".price-value")
	if err != nil {
		pp.Println("🚀 ~ file: price.go ~ line 12 ~ funcFetchPrice ~ err : ", err)
		return ""
	}
	text, err := price.Text()
	if err != nil {
		pp.Println("🚀 ~ file: price.go ~ line 17 ~ funcFetchPrice ~ err : ", err)
		return ""
	}
	return text

}
