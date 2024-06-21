package services

import (
	"encoding/json"
	"fmt"

	"github.com/skshahriarahmedraka/WebCrawlerGolang/models"
	"github.com/xuri/excelize/v2"
)

func SaveProductToExcel(product models.Product, filePath string) error {
	f := excelize.NewFile()

	// Create a new sheet.
	index, err := f.NewSheet("Product")
	fmt.Println("🚀 ~ file: save_data_service.go ~ line 16 ~ funcSaveProductToExcel ~ err : ", err)

	// Helper function to convert data to JSON strings
	toJSON := func(v interface{}) string {
		data, err := json.Marshal(v)
		if err != nil {
			return ""
		}
		return string(data)
	}

	// Set value of cells.
	f.SetCellValue("Product", "A1", "BreadCrumb")
	f.SetCellValue("Product", "B1", product.BreadCrumb)

	f.SetCellValue("Product", "A2", "ImageUrls")
	f.SetCellValue("Product", "B2", toJSON(product.ImageUrls))

	f.SetCellValue("Product", "A3", "Catagory")
	f.SetCellValue("Product", "B3", product.Catagory)

	f.SetCellValue("Product", "A4", "ProductName")
	f.SetCellValue("Product", "B4", product.ProductName)

	f.SetCellValue("Product", "A5", "Pricing")
	f.SetCellValue("Product", "B5", product.Pricing)

	f.SetCellValue("Product", "A6", "AvailableSize")
	f.SetCellValue("Product", "B6", toJSON(product.AvailableSize))

	f.SetCellValue("Product", "A7", "SenseOfTheSize Small")
	f.SetCellValue("Product", "B7", product.SenseOfTheSize.Small)

	f.SetCellValue("Product", "A8", "SenseOfTheSize Large")
	f.SetCellValue("Product", "B8", product.SenseOfTheSize.Large)

	f.SetCellValue("Product", "A9", "TitleOfDescription")
	f.SetCellValue("Product", "B9", product.TitleOfDescription)

	f.SetCellValue("Product", "A10", "GeneralDescription")
	f.SetCellValue("Product", "B10", product.GeneralDescription)

	f.SetCellValue("Product", "A11", "GeneralDescriptionItemized")
	f.SetCellValue("Product", "B11", toJSON(product.GeneralDescriptionItemized))

	f.SetCellValue("Product", "A12", "TaleOfSize")
	f.SetCellValue("Product", "B12", toJSON(product.TaleOfSize))

	f.SetCellValue("Product", "A13", "SpecialFunction")
	f.SetCellValue("Product", "B13", toJSON(product.SpecialFunction))

	f.SetCellValue("Product", "A14", "Rating")
	f.SetCellValue("Product", "B14", product.Rating)

	f.SetCellValue("Product", "A15", "NumberOfReviews")
	f.SetCellValue("Product", "B15", product.NumberOfReviews)

	f.SetCellValue("Product", "A16", "RecommendedRate")
	f.SetCellValue("Product", "B16", product.RecommendedRate)

	f.SetCellValue("Product", "A17", "UserReviews")
	f.SetCellValue("Product", "B17", toJSON(product.UserReviews))

	f.SetCellValue("Product", "A18", "ReviewRatings")
	f.SetCellValue("Product", "B18", toJSON(product.ReviewRatings))

	f.SetCellValue("Product", "A19", "KWs")
	f.SetCellValue("Product", "B19", product.KWs)

	// Set active sheet of the workbook.
	f.SetActiveSheet(index)

	// Save spreadsheet by the given path.
	if err := f.SaveAs(filePath); err != nil {
		return err
	}

	return nil
}
