package utils

import (
	"github.com/k0kubun/pp/v3"
	"github.com/tebeka/selenium"
)

func FetchCatagory(webDriver selenium.WebDriver) string {
// 	cssPathCatagory := "html body div div div div main div div div div.articlePurchaseBox div.articleInformation div.articleNameHeader a.groupName"

	catagory, err := webDriver.FindElement(selenium.ByCSSSelector, ".categoryName")
	if err != nil {
		pp.Println("🚀 ~ file: catagory.go ~ line 9 ~ funcFetchCatagory ~ err : ", err)
		return ""
	}
	text, err := catagory.Text()
	if err != nil {
		pp.Println("🚀 ~ file: catagory.go ~ line 17 ~ funcFetchCatagory ~ err : ", err)
		return ""
	}
	return text

}
