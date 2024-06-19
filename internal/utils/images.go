package utils

import (
	"github.com/k0kubun/pp/v3"
	"github.com/tebeka/selenium"
)

func FetchImages(webDriver selenium.WebDriver) []string {
	cssPathImages := "html body div div div div main div div div div div.contentWrapper div.pdp_article_image div.article_image_wrapper div.article_image img"

	images := []string{}
	imagesElements, err := webDriver.FindElements(selenium.ByCSSSelector, cssPathImages)
	if err != nil {
		pp.Println("🚀 ~ file: images.go ~ line 18 ~ funcFetchImages ~ err : ", err)
		return images
	}
	for _, imageElement := range imagesElements {
		src, err := imageElement.GetAttribute("src")
		if err != nil {
			pp.Println("🚀 ~ file: images.go ~ line 25 ~ funcFetchImages ~ err : ", err)
			continue
		}
		images = append(images, src)
	}		

	return images 

}
