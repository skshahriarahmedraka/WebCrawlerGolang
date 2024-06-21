package adidasmen

// import (
// 	"log"
// 	"strings"
// 	"sync"

// 	"github.com/gocolly/colly/v2"
// 	"github.com/k0kubun/pp/v3"
// )

// var (
// 	TestProductListUrl string = "https://shop.adidas.jp/item/?gender=mens&order=%d"
// 	TestUrl            string = "https://shop.adidas.jp/products/IV8087/"
// 	TestCount          int    = 1
// )

// type Product struct {
// 	BreadCrumb     string
// 	ImageUrls      []string
// 	Catagory       string
// 	ProductName    string
// 	Pricing        string
// 	AvailableSize  string
// 	SenseOfTheSize string
// }

// func Main(productListUrl string, url string, count int) error {

// 	var product Product
// 	cssPathBreadCrumb := "ul.breadcrumbList.clearfix.test-breadcrumb.css-2lfxqg li.breadcrumbListItem a"
// 	cssPathCatagory := "html body div#__next div div.page.css-1umyepy div.contentsWrapper main div div.pdpContainer.css-1fgulvy div.articleOverview.test-articleOverview div.articlePurchaseBox.css-gxzada div.articleInformation.css-itvqo3 div.articleNameHeader.css-t1z1wj a.groupName"
// 	cssPathProductName := "html body div#__next div div.page.css-1umyepy div.contentsWrapper main div div.pdpContainer.css-1fgulvy div.articleOverview.test-articleOverview div.articlePurchaseBox.css-gxzada div.articleInformation.inMarquee.css-itvqo3 div.articleNameHeader.css-t1z1wj h1.itemTitle.test-itemTitle"
// 	cssPathPricing := "html body div#__next div div.page.css-1umyepy div.contentsWrapper main div div.pdpContainer.css-1fgulvy div.articleOverview.test-articleOverview div.articlePurchaseBox.css-gxzada div.articleInformation.inMarquee.css-itvqo3 div.articlePrice.test-articlePrice.css-1apqb46 p.price-text.test-price-text.mod-flat span.price-value.test-price-value"
// 	cssPathSizes := "html body div#__next div div.page.css-1umyepy div.contentsWrapper main div div.pdpContainer.css-1fgulvy div.css-zn7duo div.sizeDescription.section.css-w0j2zy div.sizeChart.test-sizeChart.css-l7ym9o table.sizeChartTable tbody"
// 	cssPathImages := " html body div#__next div div.page.css-1umyepy div.contentsWrapper main div div.pdpContainer.css-1fgulvy div.articleOverview.test-articleOverview div.articleImageWrapper.clearfix.css-cdlca7 div.contentWrapper.clearfix div.pdp_article_image div.article_image_wrapper div.clearfix,.article_image.css-10iexnx img.test-img.image.test-image"
// 	breadCrumb := []string{}
// 	availableSizes := []string{}
// 	var wg sync.WaitGroup
// 	wg.Add(1)

// 	c := colly.NewCollector(
// 		colly.AllowedDomains("shop.adidas.jp"),
// 	)
// 	c.OnResponse(func(r *colly.Response) {
// 		pp.Println("Visited", r.Request.URL)
// 	})

// 	//* Get the breadCrumb
// 	c.OnHTML(cssPathBreadCrumb, func(e *colly.HTMLElement) {
// 		breadCrumb = append(breadCrumb, e.Text)
// 	})
// 	//* Get the catagory
// 	c.OnHTML(cssPathCatagory, func(e *colly.HTMLElement) {
// 		span1Text := e.ChildText("span:nth-child(1)")
// 		span2Text := e.ChildText("span:nth-child(2)")

// 		product.Catagory = span1Text + " " + span2Text
// 	})
// 	//* Get the product name
// 	c.OnHTML(cssPathProductName, func(e *colly.HTMLElement) {
// 		product.ProductName = e.Text
// 	})
// 	//* Get the pricing
// 	c.OnHTML(cssPathPricing, func(e *colly.HTMLElement) {
// 		product.Pricing = e.Text
// 	})

// 	//* Get the available sizes

// 	//* Get the available sizes
// 	// Get the available sizes
// 	c.OnHTML(cssPathSizes, func(e *colly.HTMLElement) {
// 		// TODO: 
// 	})
// 	//* Get the images
// 	c.OnHTML(cssPathImages, func(e *colly.HTMLElement) {
// 		product.ImageUrls = append(product.ImageUrls, e.Attr("src"))
// 	})


// 	c.OnScraped(func(r *colly.Response) {
// 		wg.Done()
// 	})

// 	// downloading the target HTML page
// 	if err := c.Visit("https://shop.adidas.jp/products/IV8087"); err != nil {
// 		log.Fatalln(err)
// 	}
// 	wg.Wait()
// 	product.BreadCrumb = strings.Join(breadCrumb[1:], " > ")
// 	product.AvailableSize = strings.Join(availableSizes, ", ")
// 	pp.Println("🚀 ~ file: adidas_men.go ~ line 57 ~ funcMain ~ product : ", product)
// 	return nil
// }
