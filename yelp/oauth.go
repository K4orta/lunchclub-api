package yelp

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"

	"github.com/garyburd/go-oauth/oauth"
)

type client struct {
	client oauth.Client
	token  oauth.Credentials
}

var oauthClient client

func init() {
	oauthClient.client.Credentials.Token = os.Getenv("LCAPI_YELP_KEY")
	oauthClient.client.Credentials.Secret = os.Getenv("LCAPI_YELP_SECRET")
	oauthClient.token.Token = os.Getenv("LCAPI_YELP_TOKEN")
	oauthClient.token.Secret = os.Getenv("LCAPI_YELP_TOKEN_SECRET")
}

func (c *client) get(url string, params url.Values, v interface{}) error {
	resp, err := c.client.Get(nil, &c.token, url, params)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return fmt.Errorf("yelp status %d", resp.StatusCode)
	}
	return json.NewDecoder(resp.Body).Decode(v)
}
