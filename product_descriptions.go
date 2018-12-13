package codeswholesalesdkgolangv2

import (
	"encoding/json"
	"fmt"
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

// GetProductDescription gets a product description for a specific product.
func (sdk *SDK) GetProductDescription(productID string) (*ProductDescription, error) {
	body, err := sdk.apiGET(fmt.Sprintf(getDescription, productID))
	if err != nil {
		return nil, err
	}

	description := &ProductDescription{}
	err = json.Unmarshal(body, description)
	if err != nil {
		return nil, err
	}

	return description, nil
}
