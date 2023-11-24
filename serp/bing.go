package serp

import (
	"encoding/json"
	"fmt"

	"github.com/mslmio/oxylabs-sdk-go/oxylabs"
)

type BingSearchOpts struct {
	Domain      oxylabs.Domain
	StartPage   int
	Pages       int
	Limit       int
	Locale      string
	GeoLocation string
	UserAgent   oxylabs.UserAgent
	CallbackUrl string
	Render      oxylabs.Render
}

// Scrapes Bing via its search engine.
func (c *SerpClient) ScrapeBingSearch(
	query string,
	opts ...*BingSearchOpts,
) (*Response, error) {
	// Prepare options
	opt := &BingSearchOpts{}
	if len(opts) > 0 && opts[len(opts)-1] != nil {
		opt = opts[len(opts)-1]
	}
	SetDefaults(opt)

	// Check validity of parameters.
	err := opt.checkParameterValidity()
	if err != nil {
		return nil, err
	}

	// Prepare payload.
	payload := map[string]interface{}{
		"source":          "bing_search",
		"domain":          opt.Domain,
		"query":           query,
		"start_page":      opt.StartPage,
		"pages":           opt.Pages,
		"limit":           opt.Limit,
		"locale":          opt.Locale,
		"geo_location":    opt.GeoLocation,
		"user_agent_type": opt.UserAgent,
		"callback_url":    opt.CallbackUrl,
		"render":          opt.Render,
	}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("error marshalling payload: %v", err)
	}

	res, err := c.Req(jsonPayload)
	if err != nil {
		return nil, err
	} else {
		return res, nil
	}
}

type BingUrlOpts struct {
	UserAgent   oxylabs.UserAgent
	Render      oxylabs.Render
	CallbackUrl string
}

// Scrapes Bing via provided url.
func (c *SerpClient) ScrapeBingUrl(
	url string,
	opts ...*BingUrlOpts,
) (*Response, error) {
	// Check validity of url.
	err := oxylabs.ValidateURL(url, "bing")
	if err != nil {
		return nil, err
	}

	// Prepare options.
	opt := &BingUrlOpts{}
	if len(opts) > 0 && opts[len(opts)-1] != nil {
		opt = opts[len(opts)-1]
	}
	SetDefaults(opt)

	// Check validity of parameters.
	err = opt.checkParameterValidity()
	if err != nil {
		return nil, err
	}

	// Prepare payload.
	payload := map[string]interface{}{
		"source":          "bing",
		"url":             url,
		"user_agent_type": opt.UserAgent,
		"render":          opt.Render,
		"callback_url":    opt.CallbackUrl,
	}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("error marshalling payload: %v", err)
	}

	res, err := c.Req(jsonPayload)
	if err != nil {
		return nil, err
	} else {
		return res, nil
	}
}
