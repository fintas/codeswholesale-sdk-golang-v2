package codeswholesalesdkgolangv2

import (
	"encoding/json"
	"fmt"
)

const getLanguages = "/v2/languages"

// Language endpoint to work on languages.
type Language struct {
	Name string `json:"name"`
}

// GetLanguages fetches languages.
func (sdk *SDK) GetLanguages() ([]Language, error) {
	body, err := sdk.apiGET(getLanguages)
	fmt.Println(string(body))
	if err != nil {
		return nil, err
	}

	languages := &[]Language{}
	err = json.Unmarshal(body, languages)
	if err != nil {
		return nil, err
	}

	return *languages, nil
}
