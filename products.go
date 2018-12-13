package codeswholesalesdkgolangv2

import (
	"encoding/json"
)

const getProduct = "/v2/products/"
const getProducts = "/v2/products"

// ProductList struct.
type ProductList struct {
	Items []Product `json:"items"`
}

// Product struct.
type Product struct {
	Identifier string   `json:"identifier"`
	Images     []Image  `json:"images"`
	Languages  []string `json:"languages"`
	Links      []Link   `json:"links"`
	Name       string   `json:"name"`
	Platform   string   `json:"platform"`
	Prices     []Price  `json:"prices"`
	ProductID  string   `json:"productId"`
	Quantity   int      `json:"quantity"`
	Regions    []string `json:"regions"`
}

// Image struct.
type Image struct {
	Format string `json:"format"`
	Image  string `json:"image"`
}

// Link struct.
type Link struct {
	Deprecation string `json:"deprecation"`
	Href        string `json:"href"`
	Hreflang    string `json:"hreflang"`
	Media       string `json:"media"`
	Rel         string `json:"rel"`
	Templated   bool   `json:"templated"`
	Title       string `json:"title"`
	Type        string `json:"type"`
}

// Price struct.
type Price struct {
	From  float32 `json:"from"`
	To    float32 `json:"to"`
	Value float32 `json:"value"`
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
