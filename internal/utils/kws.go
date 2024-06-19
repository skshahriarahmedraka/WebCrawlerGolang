package utils

import (
	"strings"

	"github.com/tebeka/selenium"
)

func FetchKWs(webDriver selenium.WebDriver) string {
	// cssPathKWs := "html body div#__next div div.page.css-1umyepy div.contentsWrapper main div div.pdpContainer.css-1fgulvy div.itemTagsPosition div.test-category_link.null.css-vxqsdw div.inner a.css-1ka7r5v"
	var tags []string

	// Extract item tags
	itemTagElems, _ := webDriver.FindElements(selenium.ByCSSSelector, ".itemTagsPosition .test-category_link .inner a")
	for _, itemTagElem := range itemTagElems {
		tag, _ := itemTagElem.Text()

		tags = append(tags, tag)
	}

	return strings.Join(tags, ",")

}
