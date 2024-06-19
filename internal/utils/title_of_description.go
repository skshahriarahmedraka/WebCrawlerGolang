package utils

import (
	"github.com/k0kubun/pp/v3"
	"github.com/tebeka/selenium"
)

func FetchTitleOfDescription(webDriver selenium.WebDriver) string {
	cssPathTitleOfDescription := "html body div div div div main div div div div h2.heading.itemName.test-commentItem-topHeading"

	titleOfDes, err := webDriver.FindElement(selenium.ByCSSSelector, cssPathTitleOfDescription)
	if err != nil {
		pp.Println("🚀 ~ file: title_of_description.go ~ line 13 ~ funcFetchTitleOfDescription ~ err : ", err)
		return ""
	}
	text, err := titleOfDes.Text()
	if err != nil {
		pp.Println("🚀 ~ file: title_of_description.go ~ line 18 ~ funcFetchTitleOfDescription ~ err : ", err)
		return ""
	}
	return text

}
