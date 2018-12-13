package codeswholesalesdkgolangv2

import (
	"encoding/json"
	"fmt"
)

const getProduct = "/v2/products/"
const getProducts = "/v2/products"

type Product struct {
	Identifier string `json:"identifier"`
	Images     []struct {
		Format string `json:"format"`
		Image  string `json:"image"`
	} `json:"images"`
	Languages []string `json:"languages"`
	Links     []struct {
		Deprecation string `json:"deprecation"`
		Href        string `json:"href"`
		Hreflang    string `json:"hreflang"`
		Media       string `json:"media"`
		Rel         string `json:"rel"`
		Templated   bool   `json:"templated"`
		Title       string `json:"title"`
		Type        string `json:"type"`
	} `json:"links"`
	Name     string `json:"name"`
	Platform string `json:"platform"`
	Prices   []struct {
		From  int `json:"from"`
		To    int `json:"to"`
		Value int `json:"value"`
	} `json:"prices"`
	ProductID string   `json:"productId"`
	Quantity  int      `json:"quantity"`
	Regions   []string `json:"regions"`
}
