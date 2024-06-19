package utils

import (
	"fmt"
	"log"
	"time"

	"github.com/tebeka/selenium"
)

func Scroll(wd selenium.WebDriver, selector string) error {
	el, err := wd.FindElement(selenium.ByCSSSelector, selector)
	if err != nil {
		return fmt.Errorf("dom scrolling error: %v", err)
	}
	//fmt.Println("autoScroll...")
	if _, err := wd.ExecuteScript("arguments[0].scrollIntoView(true);", []interface{}{el}); err != nil {
		log.Printf("Failed to scroll element into view: %v", err)
	}
	time.Sleep(2 * time.Second)

	return nil
}


func ScrollToBottom(wd selenium.WebDriver) error {
	for {
		// Get the current scroll height
		oldHeight, err := wd.ExecuteScript("return document.body.scrollHeight;", nil)
		if err != nil {
			return err
		}

		// Scroll down by a large amount (enough to load new content)
		_, err = wd.ExecuteScript("window.scrollTo(0, document.body.scrollHeight);", nil)
		if err != nil {
			return err
		}

		// Wait for new content to load
		time.Sleep(2 * time.Second)

		// Get the new scroll height
		newHeight, err := wd.ExecuteScript("return document.body.scrollHeight;", nil)
		if err != nil {
			return err
		}

		// Check if the scroll height has increased
		if oldHeight == newHeight {
			break
		}
	}

	return nil
}
