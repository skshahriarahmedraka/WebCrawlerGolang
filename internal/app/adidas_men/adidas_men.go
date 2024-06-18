package adidasmen

import (
	"fmt"
	"io"
	"log"
	"regexp"
	"strconv"

	// "strconv"
	"net/http"
	"strings"
	"sync"

	// "github.com/chromedp/chromedp"
	"github.com/gocolly/colly/v2"
	"github.com/k0kubun/pp/v3"
)

var (
	TestProductListUrl string = "https://shop.adidas.jp/item/?gender=mens&order=%d"
	TestUrl            string = "https://shop.adidas.jp/products/F36640/"
	TestCount          int    = 1
)

type Product struct {
	BreadCrumb     string //* done
	ImageUrls      []string //* done
	Catagory       string //* done
	ProductName    string //* done
	Pricing        string //* done
	AvailableSize  []string //* done
	SenseOfTheSize struct { 
		Small string
		Large string
	}   //* done

	TitleOfDescription         string //* done
	GeneralDescription         string //* done
	GeneralDescriptionItemized []string //* done
	TaleOfSize                 map[string]map[string]string
	SpecialFunction            map[string]string //* done

	Rating          string 
	NumberOfReviews string
	RecommendedRate string
	UserReviews     []UserReview
	ReviewRatings   map[string]string

	KWs string //* comma separated keywords
}

type CoordinatedProduct struct {
	Name           string
	Pricing        string
	ProductNumber  string
	ImageUrl       string
	ProductPageUrl string
}

type UserReview struct {
	Date              string
	Rating            string
	ReviewTitle       string
	ReviewDescription string
	ReviewerID        string
}

func Main(productListUrl string, url string, count int) error {

	var product Product
	// cssPathBreadCrumb := "ul.breadcrumbList.clearfix.test-breadcrumb.css-2lfxqg li.breadcrumbListItem a"
	cssPathBreadCrumb := "html body div div div div main div div div.breadcrumb_wrap ul.breadcrumbList.clearfix.test-breadcrumb li.breadcrumbListItem a"
	cssPathCatagory := "html body div div div div main div div div div.articlePurchaseBox div.articleInformation div.articleNameHeader a.groupName"
	cssPathProductName := "html body div div div div main div div div div div div h1.itemTitle.test-itemTitle"
	cssPathPricing := "html body div div div div main div div div div div div.articlePrice p.price-text span.price-value.test-price-value"
	cssPathSizes := "html body div div div div main div div div div div div div ul.sizeSelectorList li.sizeSelectorListItem button.sizeSelectorListItemButton"
	cssPathImages := "html body div div div div main div div div div div.contentWrapper div.pdp_article_image div.article_image_wrapper div.article_image img"
	cssPathSenseOfTheSize := "html body div div div div main div div div div div div div div span"
	cssPathTitleOfDescription := "html body div div div div main div div div div h2.heading.itemName.test-commentItem-topHeading"
	cssPathGeneralDescription := "html body div div div div main div div div div div div.description_part.details.test-itemComment-descriptionPart"
	cssPathGeneralDescriptionItemized := "html body div div div div main div div div div div ul li.articleFeaturesItem"
	cssPathSpecialFunction := "html body div div div div main div div div div div.content div.item_part.details"
	cssPathRating := "#BVRRRatingOverall_ > span:nth-child(1)"
	cssPathNumberOfReviews := "html body div#__next div div.page.css-1umyepy div.contentsWrapper main div div.pdpContainer.css-1fgulvy div.set-pdp-bv_container.css-0 div#BVRRContainer.bv_container.js-bv_container.js-componentsTabTarget.js-review.BVBrowserFF div#BVRRWidgetID.BVRRRootElement.BVRRWidget div#BVRRContentContainerID.BVRRContainer div#BVRRQuickTakeContentContainerID.BVDI_QT.BVDI_QTDashboard.BVDI_QTDisplayStyle2 div.BVDIInside.BVDI_QTInside div.BVDIBody.BVDI_QTBody div.BVDI_QTSummaryBox div.BVDIBody.BVDI_QTBodySummaryBox div#BVRRQuickTakeSummaryID.BVRRQuickTakeSummary.BVRRQuickTakeSummaryOneCloud div.BVRRRatingSummary.BVRRSecondaryRatingSummary div.BVRRRatingSummary.BVRRPrimaryRatingSummary div.BVRRRatingSummaryStyle2 div.BVRRQuickTakeCustomWrapper div.BVRRBuyAgainContainer span.BVRRValue.BVRRBuyAgain span.BVRRNumber.BVRRBuyAgainTotal"
	cssPathReviewRating := "div.BVRRCustomRatingEntryWrapper:nth-child(1) > div:nth-child(1) > div:nth-child(2)"
	breadCrumb := []string{}
	images := []string{}

	// availableSizes := []string{}
	var wg sync.WaitGroup
	wg.Add(1)

	c := colly.NewCollector(
		colly.AllowedDomains("shop.adidas.jp"),
	)
	c.OnResponse(func(r *colly.Response) {
		pp.Println("Visited", r.Request.URL)
	})

	//* Get the breadCrumb
	c.OnHTML(cssPathBreadCrumb, func(e *colly.HTMLElement) {
		breadCrumb = append(breadCrumb, e.Text)
	})
	//* Get the catagory
	c.OnHTML(cssPathCatagory, func(e *colly.HTMLElement) {
		span1Text := e.ChildText("span:nth-child(1)")
		span2Text := e.ChildText("span:nth-child(2)")

		product.Catagory = span1Text + " " + span2Text
	})
	//* Get the product name
	c.OnHTML(cssPathProductName, func(e *colly.HTMLElement) {
		product.ProductName = e.Text
	})
	//* Get the pricing
	c.OnHTML(cssPathPricing, func(e *colly.HTMLElement) {
		product.Pricing = e.Text
	})

	//* Get the available sizes
	// Get the available sizes
	c.OnHTML(cssPathSizes, func(e *colly.HTMLElement) {
		if !strings.Contains(e.Attr("class"), "disable") {
			product.AvailableSize = append(product.AvailableSize, e.Text)
		}
	})
	//* Get the images
	c.OnHTML(cssPathImages, func(e *colly.HTMLElement) {
		src := e.Attr("src")
		if src != "" {
			images = append(images, src)
		}
	})

	//* Get the sense of the size
	cssSelector := ".css-zrdet1 .bar .marker.mod-marker_2_5"
	var once sync.Once
	c.OnHTML(cssPathSenseOfTheSize, func(e *colly.HTMLElement) {
		once.Do(func() {

			cssURL := e.Request.AbsoluteURL(e.Attr("href"))

			// Fetch the CSS file content
			cssContent, err := fetchCSS(cssURL)
			if err != nil {
				fmt.Println("Error fetching CSS:", err)
				return
			}

			// Extract the left value from the CSS content
			leftValue, err := extractLeftValue(cssContent, cssSelector)
			if err != nil {
				fmt.Println("Error extracting left value:", err)
				return
			}
			leftValue = strings.Trim(strings.Replace(leftValue, "%", "",1), " ")
			leftNum, _ := strconv.ParseFloat(leftValue,64)
			
			rightValue := 100 - leftNum
			product.SenseOfTheSize.Small = strconv.FormatFloat(leftNum, 'f', -1, 64)
			product.SenseOfTheSize.Large = strconv.FormatFloat(rightValue, 'f', -1, 64)
		})

	})

	//* Get the title of description
	c.OnHTML(cssPathTitleOfDescription, func(e *colly.HTMLElement) {
		product.TitleOfDescription = e.Text
	})

	//* Get the general description
	c.OnHTML(cssPathGeneralDescription, func(e *colly.HTMLElement) {
		product.GeneralDescription = e.Text
	})

	//* Get the general description itemized
	c.OnHTML(cssPathGeneralDescriptionItemized, func(e *colly.HTMLElement) {
		product.GeneralDescriptionItemized = append(product.GeneralDescriptionItemized, e.Text)
	})

	//* Get the special function
	product.SpecialFunction = make(map[string]string)
	c.OnHTML(cssPathSpecialFunction, func(e *colly.HTMLElement) {
		// product.SpecialFunction[e.ChildText("a")] = e.ChildText("p")
		title := e.ChildText("a")

        // Get the full text of the element
        fullText := e.DOM.Find("*").Remove().End().Text()

        // Clean up the full text
        fullText = strings.TrimSpace(fullText)

        // Save the title and corresponding text after <br> into the map
        product.SpecialFunction[title] = fullText
	})

	//* Get the rating
	c.OnHTML(cssPathRating, func(e *colly.HTMLElement) {
        fmt.Println("🚀 ~ file: adidas_men.go ~ line 201 ~ c.OnHTML ~ e.Text : ", e.Text)
		product.Rating = e.Text
	}) 
	

	//* Get the number of reviews
	// c.OnHTML(cssPathNumberOfReviews, func(e *colly.HTMLElement) {
    //     fmt.Println("🚀 ~ file: adidas_men.go ~ line 207 ~ c.OnHTML ~ e.Text : ", e.Text)
	// 	product.NumberOfReviews = e.Text
	// })

	// OnHTML callback to extract the number of reviews
	c.OnHTML(cssPathNumberOfReviews, func(e *colly.HTMLElement) {
		fmt.Println("🚀 ~ file: adidas_men.go ~ line 207 ~ c.OnHTML ~ e.Text : ", e.Text)
		product.NumberOfReviews = e.Text
	})

	//* Get the review ratings
	product.ReviewRatings = make(map[string]string)
	c.OnHTML(cssPathReviewRating, func(e *colly.HTMLElement) {
		// <div class="BVRRRating BVRRRatingRadio BVRRRatingFit"><div class="BVRRLabel BVRRRatingRadioLabel1">タイトすぎる</div><div class="BVRRRatingRadioImage"><img src="https://adidasjp.ugc.bazaarvoice.com/7896-ja_jp/2_8/5/ratingSlider.gif" class="BVImgOrSprite" alt="2.8 / 5" title="2.8 / 5"></div><div class="BVRRLabel BVRRRatingRadioLabel2">ルーズすぎる</div></div>
		title := e.ChildText("div.BVRRLabel")
		rating := e.ChildAttr("div.BVRRRatingRadioImage img", "title")
		product.ReviewRatings[title] = rating
	})


	c.OnScraped(func(r *colly.Response) {
		wg.Done()
	})

	// downloading the target HTML page
	if err := c.Visit(TestUrl); err != nil {
		log.Fatalln(err)
	}
	wg.Wait()
	product.BreadCrumb = strings.Join(breadCrumb[1:], " > ")
	// product.AvailableSize = strings.Join(availableSizes, ", ")
	product.ImageUrls = images
	pp.Println("🚀 ~ file: adidas_men.go ~ line 57 ~ funcMain ~ product : ", product)
	return nil
}

// Function to fetch the CSS file content
func fetchCSS(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("bad status: %s", resp.Status)
	}

	buf := new(strings.Builder)
	_, err = io.Copy(buf, resp.Body)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

// Function to extract the left value from the CSS content
func extractLeftValue(cssContent, cssPath string) (string, error) {
	// Clean up the CSS path
	cssPath = strings.TrimSpace(cssPath)

	// Convert CSS path to a regular expression
	re, err := regexp.Compile(fmt.Sprintf(`%s\s*{[^}]*left:\s*([^;]+);`, regexp.QuoteMeta(cssPath)))
	if err != nil {
		return "", err
	}

	// Find the left value
	matches := re.FindStringSubmatch(cssContent)
	if len(matches) > 1 {
		return matches[1], nil
	}

	return "", fmt.Errorf("left value not found")
}
