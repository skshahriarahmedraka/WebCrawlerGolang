package adidasmen

import (

	"log"

	"github.com/k0kubun/pp/v3"
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

	
	product := services.FetchProductData(webdriver, productListUrl, productID)

	err = services.SaveProductToExcel(product, "product.xlsx")
	pp.Println("🚀 ~ file: adidas_men.go ~ line 41 ~ funcMain ~ err : ", err)

	

	return nil
}
