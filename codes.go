package codeswholesalesdkgolangv2

import (
	"encoding/json"
	"fmt"
)

const getCode = "/v2/codes/"

// Code is a ordered code
type Code struct {
	Code     string  `json:"code"`
	CodeID   string  `json:"codeId"`
	CodeType string  `json:"codeType"`
	Filename string  `json:"filename"`
	Links    []Links `json:"links"`
}

// GetCode gets an ordered code by id.
func (sdk *SDK) GetCode(codeID string) (*Code, error) {
	body, err := sdk.apiGET(getCode, codeID)
	fmt.Println(string(body))
	if err != nil {
		return nil, err
	}

	code := &Code{}
	err = json.Unmarshal(body, code)
	if err != nil {
		return nil, err
	}

	return code, nil
}
