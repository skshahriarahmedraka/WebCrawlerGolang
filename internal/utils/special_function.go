package utils

import (
	"strings"

	"github.com/k0kubun/pp/v3"
	"github.com/tebeka/selenium"
)

func FetchSpecialFunction(webDriver selenium.WebDriver) map[string]string {
	// 	cssPathSpecialFunction := "html body div div div div main div div div div div.content div.item_part.details"
	// 	product.SpecialFunction = make(map[string]string)
	// 	c.OnHTML(cssPathSpecialFunction, func(e *colly.HTMLElement) {
	// 		// product.SpecialFunction[e.ChildText("a")] = e.ChildText("p")
	// 		title := e.ChildText("a")

	// 		// Get the full text of the element
	// 		fullText := e.DOM.Find("*").Remove().End().Text()

	// 		// Clean up the full text
	// 		fullText = strings.TrimSpace(fullText)

	// 		// Save the title and corresponding text after <br> into the map
	// 		product.SpecialFunction[title] = fullText
	// 	})

	specialFunction, err := webDriver.FindElements(selenium.ByCSSSelector, "div.item_part.details")
	if err != nil {
		pp.Println("🚀 ~ file: special_function.go ~ line 23 ~ funcFetchSpecialFunction ~ err : ", err)
		return nil
	}

	specialFunctionMap := make(map[string]string)
	for _, element := range specialFunction {
		title, err := element.FindElement(selenium.ByCSSSelector, "a")
		if err != nil {
			pp.Println("🚀 ~ file: special_function.go ~ line 31 ~ funcFetchSpecialFunction ~ err : ", err)
			continue
		}
		titleText, err := title.Text()
		if err != nil {
			pp.Println("🚀 ~ file: special_function.go ~ line 36 ~ funcFetchSpecialFunction ~ err : ", err)
			continue
		}

		fullText, err := element.Text()
		if err != nil {
			pp.Println("🚀 ~ file: special_function.go ~ line 42 ~ funcFetchSpecialFunction ~ err : ", err)
			continue
		}
		fullText = strings.TrimSpace(fullText)

		specialFunctionMap[titleText] = fullText
	}

	return specialFunctionMap
}
