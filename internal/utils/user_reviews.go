package utils

import (

	"strings"

	"github.com/skshahriarahmedraka/WebCrawlerGolang/models"
	"github.com/tebeka/selenium"
)

func FetchUserReviews(webDriver selenium.WebDriver) []models.UserReview {
	userReviews := []models.UserReview{}

	reviews, _ := webDriver.FindElements(selenium.ByCSSSelector, ".BVRRContentReview")
	for _, r := range reviews {
		review := models.UserReview{}

		d, _ := r.FindElement(selenium.ByCSSSelector, ".BVRRValue.BVRRReviewDate")
		date, _ := d.Text()
		review.Date = date

		t, _ := r.FindElement(selenium.ByCSSSelector, ".BVRRValue.BVRRReviewTitle")
		title, _ := t.Text()
		review.ReviewTitle = title

		des, _ := r.FindElement(selenium.ByCSSSelector, ".BVRRReviewTextContainer")
		desc, _ := des.Text()
		review.ReviewDescription = desc

		ra, _ := r.FindElement(selenium.ByCSSSelector, ".BVRRNumber.BVRRRatingNumber")
		ratingText, _ := ra.Text()
		review.Rating = ratingText
		idAttr, _ := r.GetAttribute("id")

		reviewerID := parseReviewerIDFromId(idAttr)

		review.ReviewerID = reviewerID

		userReviews = append(userReviews, review)
	}

	return userReviews

}

func parseReviewerIDFromId(id string) string {
	parts := strings.Split(id, "_")
	// The reviewer ID should be the last part of the ID string
	reviewerID := parts[len(parts)-1]
	return reviewerID
}
