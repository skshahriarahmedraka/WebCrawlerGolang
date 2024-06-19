package utils

import (
	"github.com/k0kubun/pp/v3"
	"github.com/tebeka/selenium"
)

func FetchGeneralDescription(webDriver selenium.WebDriver) string {
	cssPathGeneralDescription := "html body div div div div main div div div div div div.description_part.details.test-itemComment-descriptionPart"

	Des, err := webDriver.FindElement(selenium.ByCSSSelector, cssPathGeneralDescription)
	if err != nil {
		pp.Println("🚀 ~ file: general_description.go ~ line 13 ~ funcFetchGeneralDescription ~ err : ", err)
		return ""
	}
	text, err := Des.Text()
	if err != nil {
		pp.Println("🚀 ~ file: general_description.go ~ line 18 ~ funcFetchGeneralDescription ~ err : ", err)
		return ""
	}
	return text

}
