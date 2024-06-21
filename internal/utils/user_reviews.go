package utils

import (

	// "github.com/k0kubun/pp/v3"
	"strings"

	"github.com/skshahriarahmedraka/WebCrawlerGolang/models"
	"github.com/tebeka/selenium"
)

func FetchUserReviews(webDriver selenium.WebDriver) []models.UserReview {
	userReviews := []models.UserReview{}

	// var userReviews []Review
	reviewElems, _ := webDriver.FindElements(selenium.ByCSSSelector, ".BVRRContentReview")
	for _, reviewElem := range reviewElems {
		review := models.UserReview{}

		// Extract review date
		dateElem, _ := reviewElem.FindElement(selenium.ByCSSSelector, ".BVRRValue.BVRRReviewDate")
		date, _ := dateElem.Text()
		review.Date = date

		// Extract review title
		titleElem, _ := reviewElem.FindElement(selenium.ByCSSSelector, ".BVRRValue.BVRRReviewTitle")
		title, _ := titleElem.Text()
		review.ReviewTitle = title

		// Extract review description
		descElem, _ := reviewElem.FindElement(selenium.ByCSSSelector, ".BVRRReviewTextContainer")
		desc, _ := descElem.Text()
		review.ReviewDescription = desc

		// Extract review rating
		ratingElem, _ := reviewElem.FindElement(selenium.ByCSSSelector, ".BVRRNumber.BVRRRatingNumber")
		ratingText, _ := ratingElem.Text()
		review.Rating = ratingText
		idAttr, _ := reviewElem.GetAttribute("id")

		// Extract the reviewer ID
		reviewerID := parseReviewerIDFromId(idAttr)

		// Set the reviewer ID
		review.ReviewerID = reviewerID

		// Append review to the slice
		userReviews = append(userReviews, review)
	}
	// productMeta.UserReviews = userReviews

	return userReviews

}

func parseReviewerIDFromId(id string) string {
	parts := strings.Split(id, "_")
	// The reviewer ID should be the last part of the ID string
	reviewerID := parts[len(parts)-1]
	return reviewerID
}
