package codeswholesalesdkgolangv2

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
)

const liveEndpointURL = "https://api.codeswholesale.com"
const sandboxEndpointURL = "https://sandbox.codeswholesale.com"

// SDK is the main api object
type SDK struct {
	clientID        string
	clientSecret    string
	clientSignature string
	live            bool
	authorization   *Authorization
}

// NewSDK return a new sdk.
func NewSDK(clientID string, clientSecret string, clientSignature string, live bool) *SDK {
	return &SDK{
		clientID:        clientID,
		clientSecret:    clientSecret,
		clientSignature: clientSignature,
		live:            live,
	}
}

func (sdk *SDK) apiGET(paths ...string) ([]byte, error) {
	// Is authorized
	if !sdk.isAuthorized() {
		err := sdk.authorize()
		if err != nil {
			fmt.Println(err)
		}
	}

	// Endpoint setup.
	var base *url.URL
	u, _ := url.Parse(path.Join(paths...))
	if sdk.live {
		base, _ = url.Parse(liveEndpointURL)
	} else {
		base, _ = url.Parse(sandboxEndpointURL)
	}
	endpoint := base.ResolveReference(u)

	// Build the request.
	client := &http.Client{}
	r, err := http.NewRequest("GET", endpoint.String(), nil)
	if err != nil {
		return nil, err
	}
	r.Header.Add("accept", "application/json")
	r.Header.Add("Authorization", sdk.authorization.TokenType+" "+sdk.authorization.AccessToken)

	// Do the request.
	resp, err := client.Do(r)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}

	// Get the request body.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
