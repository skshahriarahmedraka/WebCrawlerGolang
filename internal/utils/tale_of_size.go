package utils

import (
	"github.com/k0kubun/pp/v3"
	"github.com/tebeka/selenium"
)

func FetchTaleOfSize(webDriver selenium.WebDriver) map[string]map[string]string {
	sizeChart := make(map[string]map[string]string)

	// Find all table rows in the size chart
	rows, err := webDriver.FindElements(selenium.ByCSSSelector, ".sizeChartTRow")
	if err != nil {
		pp.Printf("Failed to find size chart rows: %v", err)
		return sizeChart
	}

	// Extract column headers
	columnHeaders, err := webDriver.FindElements(selenium.ByCSSSelector, ".sizeChartTHeaderCell")
	if err != nil {
		pp.Printf("Failed to find size chart column headers: %v", err)
		return sizeChart
	}

	var headers []string
	for _, header := range columnHeaders {
		text, err := header.Text()
		if err != nil {
			pp.Printf("Failed to get column header text: %v", err)
			continue
		}
		headers = append(headers, text)
	}

	// Iterate over each row
	for _, row := range rows {
		// Find all table cells in the row
		cells, err := row.FindElements(selenium.ByCSSSelector, ".sizeChartTCell")
		if err != nil {
			pp.Printf("Failed to find size chart cells: %v", err)
			continue
		}

		var measurements []string

		// Iterate over each cell
		for _, cell := range cells {
			// Get the text content of the cell
			text, err := cell.Text()
			if err != nil {
				pp.Printf("Failed to get size chart cell text: %v", err)
				continue
			}

			// Append the text content to the measurements slice
			measurements = append(measurements, text)
		}

		// Check if the row contains measurements
		if len(measurements) > 0 && len(measurements) == len(headers) {
			for i, header := range headers {
				if _, exists := sizeChart[header]; !exists {
					sizeChart[header] = make(map[string]string)
				}
				// Assuming the first column in each row is the size name
				sizeChart[header][measurements[0]] = measurements[i]
			}
		}
	}

	return sizeChart

}
