package serp

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/mslmio/oxylabs-sdk-go/oxylabs"
)

// Accepted parameters for baidu.
var BaiduSearchAcceptedDomainParameters = []oxylabs.Domain{
	oxylabs.DOMAIN_COM,
	oxylabs.DOMAIN_CN,
}

// checkParameterValidity checks validity of ScrapeBaiduSearch parameters.
func (opt *BaiduSearchOpts) checkParameterValidity() error {
	if !oxylabs.InList(opt.Domain, BaiduSearchAcceptedDomainParameters) {
		return fmt.Errorf("invalid domain parameter: %s", opt.Domain)
	}

	if !oxylabs.IsUserAgentValid(opt.UserAgent) {
		return fmt.Errorf("invalid user agent parameter: %v", opt.UserAgent)
	}

	if opt.Limit <= 0 || opt.Pages <= 0 || opt.StartPage <= 0 {
		return fmt.Errorf("limit, pages and start_page parameters must be greater than 0")
	}

	return nil
}

// checkParameterValidity checks validity of ScrapeBaiduUrl parameters.
func (opt *BaiduUrlOpts) checkParameterValidity() error {
	if !oxylabs.IsUserAgentValid(opt.UserAgent) {
		return fmt.Errorf("invalid user agent parameter: %v", opt.UserAgent)
	}

	return nil
}

// BaiduSearchOpts contains all the query parameters available for baidu_search.
type BaiduSearchOpts struct {
	Domain            oxylabs.Domain
	StartPage         int
	Pages             int
	Limit             int
	UserAgent         oxylabs.UserAgent
	CallbackUrl       string
	ParseInstructions *map[string]interface{}
	WaitTime          time.Duration
}

// ScrapeBaiduSearch scrapes baidu via Oxylabs SERP API with baidu_search as source.
func (c *SerpClient) ScrapeBaiduSearch(
	query string,
	opts ...*BaiduSearchOpts,
) (*Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), oxylabs.DefaultTimeout)
	defer cancel()

	return c.ScrapeBaiduSearchCtx(ctx, query, opts...)
}

// ScrapeBaiduSearchCtx scrapes baidu via Oxylabs SERP API with baidu_search as source.
// The provided context allows customization of the HTTP request, including setting timeouts.
func (c *SerpClient) ScrapeBaiduSearchCtx(
	ctx context.Context,
	query string,
	opts ...*BaiduSearchOpts,
) (*Response, error) {
	// Prepare options.
	opt := &BaiduSearchOpts{}
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
		"source":          oxylabs.BaiduSearch,
		"domain":          opt.Domain,
		"query":           query,
		"start_page":      opt.StartPage,
		"pages":           opt.Pages,
		"limit":           opt.Limit,
		"user_agent_type": opt.UserAgent,
		"callback_url":    opt.CallbackUrl,
	}

	// Add custom parsing instructions to the payload if provided.
	customParserFlag := false
	if opt.ParseInstructions != nil {
		payload["parse"] = true
		payload["parsing_instructions"] = &opt.ParseInstructions
		customParserFlag = true
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("error marshalling payload: %v", err)
	}

	// Request.
	res, err := c.Req(ctx, jsonPayload, customParserFlag, customParserFlag, "POST")
	if err != nil {
		return nil, err
	}

	return res, nil
}

// BaiduUrlOpts contains all the query parameters available for baidu.
type BaiduUrlOpts struct {
	UserAgent         oxylabs.UserAgent
	CallbackUrl       string
	ParseInstructions *map[string]interface{}
	WaitTime          time.Duration
}

// ScrapeBaiduUrl scrapes baidu via Oxylabs SERP API with baidu as source.
func (c *SerpClient) ScrapeBaiduUrl(
	url string,
	opts ...*BaiduUrlOpts,
) (*Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), oxylabs.DefaultTimeout)
	defer cancel()

	return c.ScrapeBaiduUrlCtx(ctx, url, opts...)
}

// ScrapeBaiduUrlCtx scrapes baidu via Oxylabs SERP API with baidu as source.
// The provided context allows customization of the HTTP request, including setting timeouts.
func (c *SerpClient) ScrapeBaiduUrlCtx(
	ctx context.Context,
	url string,
	opts ...*BaiduUrlOpts,
) (*Response, error) {
	// Check validity of url.
	err := oxylabs.ValidateURL(url, "baidu")
	if err != nil {
		return nil, err
	}

	// Prepare options.
	opt := &BaiduUrlOpts{}
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
		"source":          oxylabs.BaiduUrl,
		"url":             url,
		"user_agent_type": opt.UserAgent,
		"callback_url":    opt.CallbackUrl,
	}

	// Add custom parsing instructions to the payload if provided.
	customParserFlag := false
	if opt.ParseInstructions != nil {
		payload["parse"] = true
		payload["parsing_instructions"] = &opt.ParseInstructions
		customParserFlag = true
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("error marshalling payload: %v", err)
	}

	// Request.
	res, err := c.Req(ctx, jsonPayload, customParserFlag, customParserFlag, "POST")
	if err != nil {
		return nil, err
	}

	return res, nil
}
