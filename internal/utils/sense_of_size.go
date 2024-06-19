package utils

import (
	"fmt"
	"io/ioutil"
	// "log"
	"net/http"
	// "strconv"
	"strings"
	// "sync"

	"github.com/tebeka/selenium"
)

func FetchSenseOfTheSize(webDriver selenium.WebDriver) struct {
	Small string
	Large string
} {
	var sense struct {
		Small string
		Large string
	}

	return sense

}

func fetchCSS(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func extractLeftValue(cssContent, selector string) (string, error) {
	lines := strings.Split(cssContent, "\n")
	for _, line := range lines {
		if strings.Contains(line, selector) && strings.Contains(line, "left:") {
			parts := strings.Split(line, "left:")
			if len(parts) > 1 {
				return strings.TrimSpace(parts[1]), nil
			}
		}
	}
	return "", fmt.Errorf("left value not found in CSS content")
}
