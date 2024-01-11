package serp

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/mslmio/oxylabs-sdk-go/oxylabs"
)

// Accepted parameters for bing.
var BingSearchAcceptedDomainParameters = []oxylabs.Domain{
	oxylabs.DOMAIN_COM,
	oxylabs.DOMAIN_RU,
	oxylabs.DOMAIN_UA,
	oxylabs.DOMAIN_BY,
	oxylabs.DOMAIN_KZ,
	oxylabs.DOMAIN_TR,
}

// checkParameterValidity checks validity of ScrapeBingSearch parameters.
func (opt *BingSearchOpts) checkParameterValidity() error {
	if opt.Domain != "" && !oxylabs.InList(opt.Domain, BingSearchAcceptedDomainParameters) {
		return fmt.Errorf("invalid domain parameter: %s", opt.Domain)
	}

	if !oxylabs.IsUserAgentValid(opt.UserAgent) {
		return fmt.Errorf("invalid user agent parameter: %v", opt.UserAgent)
	}

	if opt.Render != "" && !oxylabs.IsRenderValid(opt.Render) {
		return fmt.Errorf("invalid render parameter: %v", opt.Render)
	}

	if opt.Limit <= 0 || opt.Pages <= 0 || opt.StartPage <= 0 {
		return fmt.Errorf("limit, pages and start_page parameters must be greater than 0")
	}

	return nil
}

// checkParameterValidity checks validity of ScrapeBingUrl parameters.
func (opt *BingUrlOpts) checkParameterValidity() error {
	if !oxylabs.IsUserAgentValid(opt.UserAgent) {
		return fmt.Errorf("invalid user agent parameter: %v", opt.UserAgent)
	}

	if opt.Render != "" && !oxylabs.IsRenderValid(opt.Render) {
		return fmt.Errorf("invalid render parameter: %v", opt.Render)
	}

	return nil
}

// BingSearchOpts contains all the query parameters available for bing_search.
type BingSearchOpts struct {
	Domain            oxylabs.Domain
	StartPage         int
	Pages             int
	Limit             int
	Locale            oxylabs.Locale
	GeoLocation       string
	UserAgent         oxylabs.UserAgent
	CallbackUrl       string
	Render            oxylabs.Render
	Parse             bool
	ParseInstructions *map[string]interface{}
	PollInterval      time.Duration
}

// ScrapeBingSearch scrapes bing via Oxylabs SERP API with bing_search as source.
func (c *SerpClient) ScrapeBingSearch(
	query string,
	opts ...*BingSearchOpts,
) (*Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), oxylabs.DefaultTimeout)
	defer cancel()

	return c.ScrapeBingSearchCtx(ctx, query, opts...)
}

// ScrapeBingSearchCtx scrapes bing via Oxylabs SERP API with bing_search as source.
// The provided context allows customization of the HTTP request, including setting timeouts.
func (c *SerpClient) ScrapeBingSearchCtx(
	ctx context.Context,
	query string,
	opts ...*BingSearchOpts,
) (*Response, error) {
	// Prepare options.
	opt := &BingSearchOpts{}
	if len(opts) > 0 && opts[len(opts)-1] != nil {
		opt = opts[len(opts)-1]
	}

	// Set defaults.
	SetDefaultDomain(&opt.Domain)
	SetDefaultStartPage(&opt.StartPage)
	SetDefaultLimit(&opt.Limit)
	SetDefaultPages(&opt.Pages)
	SetDefaultUserAgent(&opt.UserAgent)

	// Check validity of parameters.
	err := opt.checkParameterValidity()
	if err != nil {
		return nil, err
	}

	// Prepare payload.
	payload := map[string]interface{}{
		"source":          oxylabs.BingSearch,
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
		"parse":           opt.Parse,
	}

	// Add custom parsing instructions to the payload if provided.
	customParserFlag := false
	if opt.ParseInstructions != nil {
		payload["parsing_instructions"] = &opt.ParseInstructions
		customParserFlag = true
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("error marshalling payload: %v", err)
	}

	// Request.
	res, err := c.Req(ctx, jsonPayload, opt.Parse, customParserFlag, "POST")
	if err != nil {
		return nil, err
	}

	return res, nil
}

// BingUrlOpts contains all the query parameters available for bing.
type BingUrlOpts struct {
	UserAgent         oxylabs.UserAgent
	GeoLocation       string
	Render            oxylabs.Render
	CallbackUrl       string
	Parse             bool
	ParseInstructions *map[string]interface{}
	PollInterval      time.Duration
}

// ScrapeBingUrl scrapes bing via Oxylabs SERP API with bing as source.
func (c *SerpClient) ScrapeBingUrl(
	url string,
	opts ...*BingUrlOpts,
) (*Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), oxylabs.DefaultTimeout)
	defer cancel()

	return c.ScrapeBingUrlCtx(ctx, url, opts...)
}

// ScrapeBingUrlCtx scrapes bing via Oxylabs SERP API with bing as source.
// The provided context allows customization of the HTTP request, including setting timeouts.
func (c *SerpClient) ScrapeBingUrlCtx(
	ctx context.Context,
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

	// Set defaults.
	SetDefaultUserAgent(&opt.UserAgent)

	// Check validity of parameters.
	err = opt.checkParameterValidity()
	if err != nil {
		return nil, err
	}

	// Prepare payload.
	payload := map[string]interface{}{
		"source":          oxylabs.BingUrl,
		"url":             url,
		"user_agent_type": opt.UserAgent,
		"geo_location":    opt.GeoLocation,
		"render":          opt.Render,
		"callback_url":    opt.CallbackUrl,
		"parse":           opt.Parse,
	}

	// Add custom parsing instructions to the payload if provided.
	customParserFlag := false
	if opt.ParseInstructions != nil {
		payload["parsing_instructions"] = &opt.ParseInstructions
		customParserFlag = true
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("error marshalling payload: %v", err)
	}

	// Request.
	res, err := c.Req(ctx, jsonPayload, opt.Parse, customParserFlag, "POST")
	if err != nil {
		return nil, err
	}

	return res, nil
}
