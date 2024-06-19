package services

import (
	"strings"
	"time"

	"github.com/k0kubun/pp/v3"
	"github.com/skshahriarahmedraka/WebCrawlerGolang/internal/utils"
	"github.com/skshahriarahmedraka/WebCrawlerGolang/models"
	"github.com/tebeka/selenium"
)

func FetchProductData(webDriver selenium.WebDriver, baseUrl, id string) {
	var product models.Product

	url := baseUrl + "/" + id
	pp.Println("🚀 ~ file: product_service.go ~ line 10 ~ funcFetchProductData ~ url : ", url)

	if err := webDriver.Get(url); err != nil {
		pp.Println("🚀 ~ file: product_service.go ~ line 13 ~ funcFetchProductData ~ err", err)
	}

	if err := utils.ScrollToBottom(webDriver); err != nil {
		pp.Println("🚀 ~ file: product_service.go ~ line 36 ~ iferr:=utils.ScrollToBottom ~ err : ", err)
	}

	_, err := webDriver.PageSource()
	if err != nil {
		pp.Println("🚀 ~ file: product_service.go ~ line 46 ~ funcFetchProductData ~ err : ", err)
	}

	err = webDriver.WaitWithTimeoutAndInterval(func(wd selenium.WebDriver) (bool, error) {
		_, err := wd.FindElement(selenium.ByCSSSelector, "span.BVRRNumber.BVRRBuyAgainTotal")
		if err == nil {
			return true, nil
		}
		return false, nil
	}, 10*time.Second, 500*time.Millisecond)
	if err != nil {
		pp.Println("🚀 ~ file: product_service.go ~ line 24 ~ err:=webDriver.WaitWithTimeoutAndInterval ~ err : ", err)
	}

	

	// _, err = webDriver.ExecuteScript("window.scrollTo(0, document.body.scrollHeight);", nil)
	// if err != nil {
	// 	pp.Println("🚀 ~ file: product_service.go ~ line 41 ~ funcFetchProductData ~ err : ", err)
	// }

	// err = utils.Scroll(webDriver, "div.siteMap")
	// if err != nil {
	// 	pp.Println("🚀 ~ file: product_service.go ~ line 36 ~ funcFetchProductData ~ err : ", err)
	// 	_ = utils.Scroll(webDriver, ".js-articlePromotion")
	// }

	

	// fmt.Println("🚀 ~ file: product_service.go ~ line 39 ~ funcFetchProductData ~ htmlSource : ", htmlSource)

	// err = utils.Scroll(webDriver, ".coordinateItems .carouselListitem")
	// if err != nil {
	// 	pp.Println("🚀 ~ file: product_service.go ~ line 36 ~ funcFetchProductData ~ err : ", err)
	// 	_ = utils.Scroll(webDriver, ".js-articlePromotion")
	// }

	breadCrumb := utils.FetchBreadcrumbs(webDriver)
	product.BreadCrumb = strings.Join(breadCrumb, " > ")
	product.ImageUrls = utils.FetchImages(webDriver)
	product.Catagory = utils.FetchCatagory(webDriver)
	product.ProductName = utils.FetchTitle(webDriver)
	product.Pricing = utils.FetchPrice(webDriver)
	product.AvailableSize = utils.FetchAvailableSizes(webDriver)
	product.SenseOfTheSize = utils.FetchSenseOfTheSize(webDriver)
	product.TitleOfDescription = utils.FetchTitleOfDescription(webDriver)
	product.GeneralDescription = utils.FetchGeneralDescription(webDriver)
	product.GeneralDescriptionItemized = utils.FetchGeneralDescriptionItemized(webDriver)
	product.TaleOfSize = utils.FetchTaleOfSize(webDriver)
	product.SpecialFunction = utils.FetchSpecialFunction(webDriver)
	product.Rating = utils.FetchRating(webDriver)
	product.NumberOfReviews = utils.FetchNumberOfReviews(webDriver)
	product.RecommendedRate = utils.FetchRecommendedRate(webDriver)
	product.UserReviews = utils.FetchUserReviews(webDriver)
	product.ReviewRatings = utils.FetchReviewRatings(webDriver)
	product.KWs = utils.FetchKWs(webDriver)
	pp.Println("🚀 ~ file: product_service.go ~ line 31 ~ funcFetchProductData ~ product : ", product)
}
