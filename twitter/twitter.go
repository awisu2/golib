package twitter

import (
	"errors"
	"github.com/ChimeraCoder/anaconda"
	"net/url"
)

// consumer keys
type ConsumerKeys struct {
	ConsumerKey    string
	ConsumerSecret string
}

// access token
type AccessTokens struct {
	AccessToken       string
	AccessTokenSecret string
}

var initSetting bool

// first Setting
func InitSetting(keys *ConsumerKeys) error {
	if keys == nil || keys.ConsumerKey == "" || keys.ConsumerSecret == "" {
		return errors.New("Error: no ConsumerKeys.")
	}

	anaconda.SetConsumerKey(keys.ConsumerKey)
	anaconda.SetConsumerSecret(keys.ConsumerSecret)
	initSetting = true
	return nil
}

// post tweet
func PostTweet(status string, v url.Values, access *AccessTokens) error {
	if !initSetting {
		return errors.New("Error: no initSetting")
	}

	api := anaconda.NewTwitterApi(access.AccessToken, access.AccessTokenSecret)
	_, err := api.PostTweet(status, v)
	if err != nil {
		return err
	}
	return nil
}
