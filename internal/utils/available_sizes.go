package utils

import (
	// "log"
	// "strings"

	// "github.com/k0kubun/pp/v3"
	"github.com/tebeka/selenium"
)

func FetchAvailableSizes(webdriver selenium.WebDriver) []string {
		// cssPathSizes := "html body div div div div main div div div div div div div ul.sizeSelectorList li.sizeSelectorListItem button.sizeSelectorListItemButton"
	// 	c.OnHTML(cssPathSizes, func(e *colly.HTMLElement) {
	// 		if !strings.Contains(e.Attr("class"), "disable") {
	// 			product.AvailableSize = append(product.AvailableSize, e.Text)
	// 		}
	// 	})
	vailableSizes := []string{}
	// sizesElements, err := webdriver.FindElements(selenium.ByCSSSelector, cssPathSizes)
	// if err != nil {
	// 	log.Printf("Failed to find size elements: %v", err)
	// } else {
	// 	for _, sizeElement := range sizesElements {
	// 		// Check if the element is not disabled
	// 		classAttr, err := sizeElement.GetAttribute("class")
	// 		if err != nil {
	// 			log.Printf("Failed to get class attribute: %v", err)
	// 			continue
	// 		}
	// 		if !strings.Contains(classAttr, "disable") {
	// 			size, err := sizeElement.Text()
	// 			if err != nil {
	// 				log.Printf("Failed to get size text: %v", err)
	// 				continue
	// 			}
	// 			vailableSizes = append(vailableSizes, size)
	// 		}
	// 	}
	// 	// Print the available sizes
	// }
	// pp.Println("🚀 ~ file: available_sizes.go ~ line 18 ~ funcFetchAvailableSizes ~ vailableSizes : ", vailableSizes)
	return vailableSizes
}
