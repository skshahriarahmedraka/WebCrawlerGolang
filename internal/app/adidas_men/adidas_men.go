package adidasmen

import (
	// "fmt"

	"fmt"
	"log"

	"github.com/skshahriarahmedraka/WebCrawlerGolang/internal/services"
	webdriver "github.com/skshahriarahmedraka/WebCrawlerGolang/pkg/web_driver"
	"github.com/tebeka/selenium"
)

var (
	TestProductListUrl string = "https://shop.adidas.jp/item/?gender=mens&order=%d"
	TestUrl            string = "https://shop.adidas.jp/products/F36640"
	TestCount          int    = 1
)

func Main(productListUrl string, productID string, count int) error {

	service, err := selenium.NewChromeDriverService("config/driver/chromedriver", 4444)

	if err != nil {
		log.Fatal("Error:", err)
	}

	defer service.Stop()

	// configure the browser options
	webdriver := webdriver.SeleniumDriver()
	defer webdriver.Quit()

	// baseUrl := "https://shop.adidas.jp/products"
	// productID := "F36640"
	product := services.FetchProductData(webdriver, productListUrl, productID)

	err = services.SaveProductToExcel(product, "product.xlsx")
	fmt.Println("🚀 ~ file: adidas_men.go ~ line 41 ~ funcMain ~ err : ", err)

	// visit the target page
	// err = webdriver.Get("https://scrapingclub.com/exercise/list_infinite_scroll/")
	// if err != nil {
	//  log.Fatal("Error:", err)
	// }

	// // retrieve the page raw HTML as a string
	// // and logging it

	// html, err := webdriver.PageSource()
	// if err != nil {
	//  log.Fatal("Error:", err)
	// }
	// fmt.Println(html)

	return nil
}
