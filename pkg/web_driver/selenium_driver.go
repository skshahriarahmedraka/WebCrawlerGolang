package webdriver

import (
	"github.com/k0kubun/pp/v3"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

func SeleniumDriver() selenium.WebDriver {

	caps := selenium.Capabilities{
		"browserName": "chrome",
	}
	chromeCaps := chrome.Capabilities{
		Path: "",
		Args: []string{
			"--window-size=1920,1080",
			"--headless",
			"--no-sandbox",
			"--log-level=3",
			"--disable-gpu",
			"--disable-dev-shm-usage",
			"--disable-web-security",
		},
	}
	caps.AddChrome(chromeCaps)

	// create a new remote client with the specified options
	driver, err := selenium.NewRemote(caps, "http://127.0.0.1:4444/wd/hub")
	if err != nil {
		pp.Println("🚀 ~ file: selenium_driver.go ~ line 32 ~ funcSeleniumDriver ~ err : ", err)
	}
	return driver
}
