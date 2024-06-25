package serp

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/oxylabs/oxylabs-sdk-go/internal"
	"github.com/oxylabs/oxylabs-sdk-go/oxylabs"
)

// Accepted parameters for baidu.
var BaiduSearchAcceptedDomainParameters = []oxylabs.Domain{
	oxylabs.DOMAIN_COM,
	oxylabs.DOMAIN_CN,
}

// checkParameterValidity checks validity of ScrapeBaiduSearch parameters.
func (opt *BaiduSearchOpts) checkParameterValidity() error {
	if !internal.InList(opt.Domain, BaiduSearchAcceptedDomainParameters) {
		return fmt.Errorf("invalid domain parameter: %s", opt.Domain)
	}

	if !oxylabs.IsUserAgentValid(opt.UserAgent) {
		return fmt.Errorf("invalid user agent parameter: %v", opt.UserAgent)
	}

	if opt.Limit <= 0 || opt.Pages <= 0 || opt.StartPage <= 0 {
		return fmt.Errorf("limit, pages and start_page parameters must be greater than 0")
	}

	if opt.ParseInstructions != nil {
		if err := oxylabs.ValidateParseInstructions(opt.ParseInstructions); err != nil {
			return fmt.Errorf("invalid parse instructions: %w", err)
		}
	}

	return nil
}

// checkParameterValidity checks validity of ScrapeBaiduUrl parameters.
func (opt *BaiduUrlOpts) checkParameterValidity() error {
	if !oxylabs.IsUserAgentValid(opt.UserAgent) {
		return fmt.Errorf("invalid user agent parameter: %v", opt.UserAgent)
	}

	if opt.ParseInstructions != nil {
		if err := oxylabs.ValidateParseInstructions(opt.ParseInstructions); err != nil {
			return fmt.Errorf("invalid parse instructions: %w", err)
		}
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
	PollInterval      time.Duration
}

// ScrapeBaiduSearch scrapes baidu via Oxylabs SERP API with baidu_search as source.
func (c *SerpClient) ScrapeBaiduSearch(
	query string,
	opts ...*BaiduSearchOpts,
) (*Resp, error) {
	ctx, cancel := context.WithTimeout(context.Background(), internal.DefaultTimeout)
	defer cancel()

	return c.ScrapeBaiduSearchCtx(ctx, query, opts...)
}

// ScrapeBaiduSearchCtx scrapes baidu via Oxylabs SERP API with baidu_search as source.
// The provided context allows customization of the HTTP req, including setting timeouts.
func (c *SerpClient) ScrapeBaiduSearchCtx(
	ctx context.Context,
	query string,
	opts ...*BaiduSearchOpts,
) (*Resp, error) {
	// Prepare options.
	opt := &BaiduSearchOpts{}
	if len(opts) > 0 && opts[len(opts)-1] != nil {
		opt = opts[len(opts)-1]
	}

	// Set defaults.
	internal.SetDefaultDomain(&opt.Domain)
	internal.SetDefaultStartPage(&opt.StartPage)
	internal.SetDefaultLimit(&opt.Limit, internal.DefaultLimit_SERP)
	internal.SetDefaultPages(&opt.Pages)
	internal.SetDefaultUserAgent(&opt.UserAgent)

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

	// Marshal.
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("error marshalling payload: %v", err)
	}

	// Req.
	httpResp, err := c.C.Req(ctx, jsonPayload, "POST")
	if err != nil {
		return nil, err
	}

	// Unmarshal the http Response and get the response.
	resp, err := GetResp(httpResp, customParserFlag, customParserFlag)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// BaiduUrlOpts contains all the query parameters available for baidu.
type BaiduUrlOpts struct {
	UserAgent         oxylabs.UserAgent
	CallbackUrl       string
	ParseInstructions *map[string]interface{}
	PollInterval      time.Duration
}

// ScrapeBaiduUrl scrapes baidu via Oxylabs SERP API with baidu as source.
func (c *SerpClient) ScrapeBaiduUrl(
	url string,
	opts ...*BaiduUrlOpts,
) (*Resp, error) {
	ctx, cancel := context.WithTimeout(context.Background(), internal.DefaultTimeout)
	defer cancel()

	return c.ScrapeBaiduUrlCtx(ctx, url, opts...)
}

// ScrapeBaiduUrlCtx scrapes baidu via Oxylabs SERP API with baidu as source.
// The provided context allows customization of the HTTP req, including setting timeouts.
func (c *SerpClient) ScrapeBaiduUrlCtx(
	ctx context.Context,
	url string,
	opts ...*BaiduUrlOpts,
) (*Resp, error) {
	// Check validity of URL.
	err := internal.ValidateUrl(url, "baidu")
	if err != nil {
		return nil, err
	}

	// Prepare options.
	opt := &BaiduUrlOpts{}
	if len(opts) > 0 && opts[len(opts)-1] != nil {
		opt = opts[len(opts)-1]
	}

	// Set defaults.
	internal.SetDefaultUserAgent(&opt.UserAgent)

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

	// Marshal.
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("error marshalling payload: %v", err)
	}

	// Req.
	httpResp, err := c.C.Req(ctx, jsonPayload, "POST")
	if err != nil {
		return nil, err
	}

	// Unmarshal the http Response and get the response.
	resp, err := GetResp(httpResp, customParserFlag, customParserFlag)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
