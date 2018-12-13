package codeswholesalesdkgolangv2

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const authorization = "/oauth/token"

// Authorization holds the oauth token
type Authorization struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	expires     *time.Time
}

func (sdk *SDK) authorize() error {
	// Endpoint setup.
	var endpoint string
	if sdk.live {
		endpoint = fmt.Sprintf("%s%s", liveEndpointURL, authorization)
	} else {
		endpoint = fmt.Sprintf("%s%s", sandboxEndpointURL, authorization)
	}

	// Body setup
	data := fmt.Sprintf("grant_type=client_credentials&client_id=%s&client_secret=%s", sdk.clientID, sdk.clientSecret)

	// Build the request.
	client := &http.Client{}
	r, err := http.NewRequest("POST", endpoint, strings.NewReader(data))
	if err != nil {
		return err
	}
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Do the request.
	resp, err := client.Do(r)
	if err != nil {
		return err
	}

	// Get the request body.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Parse auth
	auth := &Authorization{}
	err = json.Unmarshal(body, auth)
	fmt.Println(string(body))
	if err != nil {
		return err
	}

	// Set expire timestamp
	now := time.Now()
	now.Add(time.Duration(auth.ExpiresIn) * time.Second)
	auth.expires = &now

	sdk.authorization = auth

	return nil
}

func (sdk *SDK) isAuthorized() bool {
	if sdk.authorization != nil && time.Since(*sdk.authorization.expires).Seconds() < 0 {
		return true
	}

	return false
}
