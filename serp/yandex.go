package serp

import (
	"encoding/json"
	"fmt"

	"github.com/mslmio/oxylabs-sdk-go/oxylabs"
)

type YandexSearchOpts struct {
	Domain      oxylabs.Domain
	StartPage   int
	Pages       int
	Limit       int
	Locale      string
	GeoLocation string
	UserAgent   oxylabs.UserAgent
	CallbackUrl string
}

// Scrapes Yandex via its search engine.
func (c *SerpClient) ScrapeYandexSearch(
	query string,
	opts ...*YandexSearchOpts,
) (*Response, error) {
	// Prepare options
	opt := &YandexSearchOpts{}
	if len(opts) > 0 && opts[len(opts)-1] != nil {
		opt = opts[len(opts)-1]
	}
	opt.setDefaults()

	// Check validity of parameters.
	err := opt.checkParameterValidity()
	if err != nil {
		return nil, err
	}

	// Prepare payload.
	payload := map[string]interface{}{
		"source":          "yandex_search",
		"domain":          opt.Domain,
		"query":           query,
		"start_page":      opt.StartPage,
		"pages":           opt.Pages,
		"limit":           opt.Limit,
		"locale":          opt.Locale,
		"geo_location":    opt.GeoLocation,
		"user_agent_type": opt.UserAgent,
		"callback_url":    opt.CallbackUrl,
	}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		fmt.Printf("error marshalling payload: %v", err)
	}

	res, err := c.Req(jsonPayload)
	if err != nil {
		return nil, err
	} else {
		return res, nil
	}
}

type YandexUrlOpts struct {
	UserAgent   oxylabs.UserAgent
	Render      oxylabs.Render
	CallbackUrl string
}

// Scrapes Yandex via provided url.
func (c *SerpClient) ScrapeYandexUrl(
	url string,
	opts ...*YandexUrlOpts,
) (*Response, error) {
	// Check validity of url.
	err := oxylabs.ValidateURL(url, "yandex")
	if err != nil {
		return nil, err
	}

	// Prepare options.
	opt := &YandexUrlOpts{}
	if len(opts) > 0 && opts[len(opts)-1] != nil {
		opt = opts[len(opts)-1]
	}
	opt.setDefaults()

	if opt.Render != "" && !oxylabs.IsRenderValid(opt.Render) {
		return nil, fmt.Errorf("invalid render option: %v", opt.Render)
	}

	if !oxylabs.IsUserAgentValid(string(opt.UserAgent)) {
		return nil, fmt.Errorf("invalid user agent option: %v", opt.UserAgent)
	}

	// Prepare payload.
	payload := map[string]interface{}{
		"source":          "yandex",
		"url":             url,
		"user_agent_type": opt.UserAgent,
		"render":          opt.Render,
		"callback_url":    opt.CallbackUrl,
	}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		fmt.Printf("error marshalling payload: %v", err)
	}

	res, err := c.Req(jsonPayload)
	if err != nil {
		return nil, err
	} else {
		return res, nil
	}

}
