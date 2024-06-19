package models



type Product struct {
	BreadCrumb     string   //* done
	ImageUrls      []string //* done
	Catagory       string   //* done
	ProductName    string   //* done
	Pricing        string   //* done
	AvailableSize  []string //* done
	SenseOfTheSize struct {
		Small string
		Large string
	} //* done

	TitleOfDescription         string   //* done
	GeneralDescription         string   //* done
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