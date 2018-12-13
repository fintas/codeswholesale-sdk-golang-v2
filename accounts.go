package codeswholesalesdkgolangv2

import (
	"encoding/json"
)

const getAccount = "/v2/accounts/current"

// Account provides you whole information about your account from name to actual balance with credit included.
type Account struct {
	FullName       string  `json:"fullName"`
	Email          string  `json:"email"`
	CurrentBalance float64 `json:"currentBalance"`
	CurrentCredit  int     `json:"currentCredit"`
	TotalToUse     float64 `json:"totalToUse"`
	Links          []struct {
		Rel         string      `json:"rel"`
		Href        string      `json:"href"`
		Hreflang    interface{} `json:"hreflang"`
		Media       interface{} `json:"media"`
		Title       interface{} `json:"title"`
		Type        interface{} `json:"type"`
		Deprecation interface{} `json:"deprecation"`
	} `json:"links"`
}

// GetAccount provides you whole information about your account from name to actual balance with credit included.
func (sdk *SDK) GetAccount() (*Account, error) {
	body, err := sdk.apiGET(getAccount)
	if err != nil {
		return nil, err
	}

	acc := &Account{}
	err = json.Unmarshal(body, acc)
	if err != nil {
		return nil, err
	}

	return acc, nil
}
