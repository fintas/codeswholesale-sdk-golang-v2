package codeswholesalesdkgolangv2

import (
	"encoding/json"
)

const getDescription = "/v2/products/%s/description"

// ProductDescription struct.
type ProductDescription struct {
	Category                string            `json:"category"`
	DeveloperHomepage       string            `json:"developerHomepage"`
	DeveloperName           string            `json:"developerName"`
	Eans                    []string          `json:"eans"`
	Editions                []string          `json:"editions"`
	ExtensionPacks          []string          `json:"extensionPacks"`
	FactSheets              []FactSheets      `json:"factSheets"`
	InTheGameLanguages      []string          `json:"inTheGameLanguages"`
	Keywords                string            `json:"keywords"`
	LocalizedTitles         []LocalizedTitles `json:"localizedTitles"`
	MinimumRequirements     string            `json:"minimumRequirements"`
	OfficialTitle           string            `json:"officialTitle"`
	Pegirating              string            `json:"pegirating"`
	Photos                  []Photos          `json:"photos"`
	Platform                string            `json:"platform"`
	ProductID               string            `json:"productId"`
	RecommendedRequirements string            `json:"recommendedRequirements"`
	Releases                []Releases        `json:"releases"`
	Videos                  []Videos          `json:"videos"`
}

// FactSheets struct.
type FactSheets struct {
	Description string `json:"description"`
	Territory   string `json:"territory"`
}

// LocalizedTitles struct.
type LocalizedTitles struct {
	Territory string `json:"territory"`
	Title     string `json:"title"`
}

// Photos struct.
type Photos struct {
	Territory string `json:"territory"`
	Type      string `json:"type"`
	URL       string `json:"url"`
}

// Releases struct.
type Releases struct {
	ReleaseDate   string `json:"releaseDate"`
	ReleaseStatus string `json:"releaseStatus"`
	Territory     string `json:"territory"`
}

// Videos struct.
type Videos struct {
	AgeWarning      bool   `json:"ageWarning"`
	PreviewFrameURL string `json:"previewFrameURL"`
	Title           string `json:"title"`
	URL             string `json:"url"`
}

// GetProducts gets all products from price list.
func (sdk *SDK) GetProducts() ([]Product, error) {
	body, err := sdk.apiGET(getProducts)
	if err != nil {
		return nil, err
	}

	list := &ProductList{}
	err = json.Unmarshal(body, list)
	if err != nil {
		return nil, err
	}

	return list.Items, nil
}

// GetProduct gets a specific product from price list.
func (sdk *SDK) GetProduct(productID string) (*Product, error) {
	body, err := sdk.apiGET(getProduct, productID)
	if err != nil {
		return nil, err
	}

	product := &Product{}
	err = json.Unmarshal(body, product)
	if err != nil {
		return nil, err
	}

	return product, nil
}
