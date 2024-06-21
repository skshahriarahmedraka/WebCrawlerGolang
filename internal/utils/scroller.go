package utils

import (
	"fmt"
	"time"

	"github.com/tebeka/selenium"
)

func Scroll(wd selenium.WebDriver, selector string) error {
	el, err := wd.FindElement(selenium.ByCSSSelector, selector)
	if err != nil {
		return fmt.Errorf("dom scrolling error: %v", err)
	}
	if _, err := wd.ExecuteScript("arguments[0].scrollIntoView(true);", []interface{}{el}); err != nil {
		return fmt.Errorf("failed to scroll element into view: %v", err)
	}
	time.Sleep(2 * time.Second)
	return nil
}

// func ScrollToBottom(wd selenium.WebDriver) error {
// 	for {
// 		// Get the current scroll height
// 		oldHeight, err := wd.ExecuteScript("return document.body.scrollHeight;", nil)
// 		if err != nil {
// 			return err
// 		}

// 		// Scroll down by a large amount (enough to load new content)
// 		_, err = wd.ExecuteScript("window.scrollTo(0, document.body.scrollHeight);", nil)
// 		if err != nil {
// 			return err
// 		}

// 		// Wait for new content to load
// 		time.Sleep(2 * time.Second)

// 		// Get the new scroll height
// 		newHeight, err := wd.ExecuteScript("return document.body.scrollHeight;", nil)
// 		if err != nil {
// 			return err
// 		}

// 		// Check if the scroll height has increased
// 		if oldHeight == newHeight {
// 			break
// 		}
// 	}

// 	return nil
// }

func WaitForPageLoad(webDriver selenium.WebDriver) error {
	for {
		readyState, err := webDriver.ExecuteScript("return document.readyState;", nil)
		fmt.Println("🚀 ~ file: scroller.go ~ line 58 ~ funcWaitForPageLoad ~ readyState : ", readyState)
		fmt.Printf("🚀 ~ file: scroller.go ~ line 58 ~ funcWaitForPageLoad ~ readyState : %T", readyState)
		if err != nil {
			return err
		}
		if readyState == "complete" {
			break
		}
		time.Sleep(500 * time.Millisecond)
	}
	return nil
}

func GetPageHeight(wd selenium.WebDriver) (int, error) {
	height, err := wd.ExecuteScript("return document.body.scrollHeight;", nil)
	if err != nil {
		return 0, fmt.Errorf("error getting page height: %v", err)
	}
	return int(height.(float64)), nil
}

// scrollToBottom scrolls to the bottom of the page until no more new content is loaded.
func ScrollToBottom(wd selenium.WebDriver) error {
	for {
		initialHeight, err := GetPageHeight(wd)
		if err != nil {
			return fmt.Errorf("error getting initial page height: %v", err)
		}

		if _, err := wd.ExecuteScript("window.scrollTo(0, document.body.scrollHeight);", nil); err != nil {
			return fmt.Errorf("error scrolling down: %v", err)
		}

		// Wait for new content to load
		time.Sleep(3 * time.Second)

		newHeight, err := GetPageHeight(wd)
		if err != nil {
			return fmt.Errorf("error getting new page height: %v", err)
		}

		if newHeight == initialHeight {
			// If the height hasn't changed, assume we've reached the bottom
			break
		}
	}
	return nil
}
