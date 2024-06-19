package utils

import (
	"github.com/k0kubun/pp/v3"
	"github.com/tebeka/selenium"
)

func FetchBreadcrumbs(webDriver selenium.WebDriver) []string {
	var breadcrumbs []string
	// cssPathBreadCrumb := "html body div#__next div div.page.css-1umyepy div.contentsWrapper main div div.pdpContainer.css-1fgulvy div.breadcrumb_wrap ul.breadcrumbList.clearfix.test-breadcrumb.css-2lfxqg li.breadcrumbListItem.breadcrumbLink.test-breadcrumbLink a.breadcrumbListItemLink"

	// breadcrumbList, err := webDriver.FindElements(selenium.ByCSSSelector, ".breadcrumbListItemLink")
	breadcrumbList, err := webDriver.FindElements(selenium.ByCSSSelector, "li.breadcrumbListItem.breadcrumbLink.test-breadcrumbLink a.breadcrumbListItemLink")
	if err != nil {
		pp.Println("🚀 ~ file: breadcrumb.go ~ line 14 ~ funcFetchBreadcrumbs ~ err : ", err)
		return []string{}
	}
	for _, item := range breadcrumbList {
		text, err := item.Text()
		if err != nil {
			pp.Println("🚀 ~ file: breadcrumb.go ~ line 21 ~ funcFetchBreadcrumbs ~ err : ", err)
			continue
		}
		breadcrumbs = append(breadcrumbs, text)
	}

	return breadcrumbs
}
