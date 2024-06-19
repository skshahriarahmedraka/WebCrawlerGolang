package utils

import (
	"github.com/k0kubun/pp/v3"
	"github.com/tebeka/selenium"
)

func FetchTitle(webDriver selenium.WebDriver) string {
	// 	cssPathProductName := "html body div div div div main div div div div div div h1.itemTitle.test-itemTitle"

	title, err := webDriver.FindElement(selenium.ByCSSSelector, ".itemTitle")
	if err != nil {
		pp.Println("🚀 ~ file: title.go ~ line 12 ~ funcFetchTitle ~ err : ", err)
		return ""
	}
	text, err := title.Text()
	if err != nil {
		pp.Println("🚀 ~ file: title.go ~ line 17 ~ funcFetchTitle ~ err : ", err)
		return ""
	}
	return text

}
