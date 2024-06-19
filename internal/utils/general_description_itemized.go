package utils

import (
	"github.com/k0kubun/pp/v3"
	"github.com/tebeka/selenium"
)

func FetchGeneralDescriptionItemized(webDriver selenium.WebDriver) []string {

	cssPathGeneralDescriptionItemized := "html body div div div div main div div div div div ul li.articleFeaturesItem"

	itemized := []string{}

	items, err := webDriver.FindElements(selenium.ByCSSSelector, cssPathGeneralDescriptionItemized)
	if err != nil {
		pp.Println("🚀 ~ file: general_description_itemized.go ~ line 16 ~ funcFetchGeneralDescriptionItemized ~ err : ", err)
		return itemized
	}
	// text, err := Des.Text()
	// if err != nil {
	// 	pp.Println("🚀 ~ file: general_description.go ~ line 18 ~ funcFetchGeneralDescription ~ err : ", err)
	// 	return ""
	// }
	for _, item := range items {
		text, err := item.Text()
		if err != nil {
			pp.Println("🚀 ~ file: general_description_itemized.go ~ line 29 ~ funcFetchGeneralDescriptionItemized ~ err : ", err)
			continue
		}
		itemized = append(itemized, text)
	}

	return itemized

}
