package services

import (
	"strings"
	"time"

	"github.com/k0kubun/pp/v3"
	"github.com/skshahriarahmedraka/WebCrawlerGolang/internal/utils"
	"github.com/skshahriarahmedraka/WebCrawlerGolang/models"
	"github.com/tebeka/selenium"
)

func FetchProductData(webDriver selenium.WebDriver, baseUrl, id string) models.Product {
	var product models.Product

	url := baseUrl + "/" + id
	pp.Println("🚀 ~ file: product_service.go ~ line 10 ~ funcFetchProductData ~ url : ", url)

	if err := webDriver.Get(url); err != nil {
		pp.Println("🚀 ~ file: product_service.go ~ line 13 ~ funcFetchProductData ~ err", err)
	}
	

	breadCrumb := utils.FetchBreadcrumbs(webDriver)
	product.BreadCrumb = strings.Join(breadCrumb, " > ")

	product.Catagory = utils.FetchCatagory(webDriver)
	product.ProductName = utils.FetchTitle(webDriver)
	product.Pricing = utils.FetchPrice(webDriver)
	product.AvailableSize = utils.FetchAvailableSizes(webDriver)
	product.SenseOfTheSize = utils.FetchSenseOfTheSize(webDriver)


	err := utils.Scroll(webDriver, "h2.heading.itemName.test-commentItem-topHeading")
	if err != nil {
		pp.Println("🚀 ~ file: product_service.go ~ line 36 ~ funcFetchProductData ~ err : ", err)
		_ = utils.Scroll(webDriver, ".js-articlePromotion")
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



	product.ImageUrls = utils.FetchImages(webDriver)

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

	return product
}
