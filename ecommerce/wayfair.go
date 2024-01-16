package ecommerce

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/mslmio/oxylabs-sdk-go/internal"
	"github.com/mslmio/oxylabs-sdk-go/oxylabs"
)

// checkParameterValidity checks validity of ScrapeWayfairSearch parameters.
func (opt *WayfairSearchOpts) checkParametersValidity() error {
	if !oxylabs.IsUserAgentValid(opt.UserAgent) {
		return fmt.Errorf("invalid user agent parameter: %v", opt.UserAgent)
	}

	if opt.Limit <= 0 || opt.Pages <= 0 || opt.StartPage <= 0 {
		return fmt.Errorf("limit, pages and start_page parameters must be greater than 0")
	}

	if opt.Limit != 24 && opt.Limit != 48 && opt.Limit != 96 {
		return fmt.Errorf("invalid limit parameter: %v", opt.Limit)
	}

	return nil
}

// WayfairSearchOpts contains all the query parameters available for wayfair_search.
type WayfairSearchOpts struct {
	StartPage         int
	Pages             int
	Limit             int
	UserAgent         oxylabs.UserAgent
	CallbackUrl       string
	Parse             bool
	ParseInstructions *map[string]interface{}
	PollInterval      time.Duration
}

// ScrapeWayfairSearch scrapes wayfair via Oxylabs E-Commerce API with wayfair_search as source.
func (c *EcommerceClient) ScrapeWayfairSearch(
	query string,
	opts ...*WayfairSearchOpts,
) (*EcommerceResp, error) {
	ctx, cancel := context.WithTimeout(context.Background(), internal.DefaultTimeout)
	defer cancel()

	return c.ScrapeWayfairSearchCtx(ctx, query, opts...)
}

// ScrapeWayfairSearchCtx scrapes wayfair via Oxylabs E-Commerce API with wayfair_search as source.
// The provided context allows customization of the HTTP req, including setting timeouts.
func (c *EcommerceClient) ScrapeWayfairSearchCtx(
	ctx context.Context,
	query string,
	opts ...*WayfairSearchOpts,
) (*EcommerceResp, error) {
	// Prepare options.
	opt := &WayfairSearchOpts{}
	if len(opts) > 0 && opts[len(opts)-1] != nil {
		opt = opts[len(opts)-1]
	}

	// Set defaults.
	internal.SetDefaultPages(&opt.Pages)
	internal.SetDefaultStartPage(&opt.StartPage)
	internal.SetDefaultUserAgent(&opt.UserAgent)
	internal.SetDefaultLimit(&opt.Limit, internal.DefaultLimit_ECOMMERCE)

	// Check validity of parameters.
	err := opt.checkParametersValidity()
	if err != nil {
		return nil, err
	}

	// Prepare payload.
	payload := map[string]interface{}{
		"source":          oxylabs.WayfairSearch,
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
		payload["parsing_instructions"] = &opt.ParseInstructions
		customParserFlag = true
	}

	// Marshal.
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("error marshalling payload: %v", err)
	}

	// Req.
	internalResp, err := c.C.Req(ctx, jsonPayload, opt.Parse, customParserFlag, "POST")
	if err != nil {
		return nil, err
	}

	// Map resp.
	resp := &EcommerceResp{
		Resp: *internalResp,
	}

	return resp, nil
}

// WayfairUrlOpts contains all the query parameters available for wayfair.
type WayfairUrlOpts struct {
	UserAgent         oxylabs.UserAgent
	CallbackUrl       string
	Parse             bool
	ParseInstructions *map[string]interface{}
	PollInterval      time.Duration
}

// checkParameterValidity checks validity of ScrapeWayfairUrl parameters.
func (opt *WayfairUrlOpts) checkParametersValidity() error {
	if !oxylabs.IsUserAgentValid(opt.UserAgent) {
		return fmt.Errorf("invalid user agent parameter: %v", opt.UserAgent)
	}

	return nil
}

// ScrapeWayfairUrl scrapes wayfair via Oxylabs E-Commerce API with wayfair as source.
func (c *EcommerceClient) ScrapeWayfairUrl(
	url string,
	opts ...*WayfairUrlOpts,
) (*EcommerceResp, error) {
	ctx, cancel := context.WithTimeout(context.Background(), internal.DefaultTimeout)
	defer cancel()

	return c.ScrapeWayfairUrlCtx(ctx, url, opts...)
}

// ScrapeWayfairUrlCtx scrapes wayfair via Oxylabs E-Commerce API with wayfair as source.
// The provided context allows customization of the HTTP req, including setting timeouts.
func (c *EcommerceClient) ScrapeWayfairUrlCtx(
	ctx context.Context,
	url string,
	opts ...*WayfairUrlOpts,
) (*EcommerceResp, error) {
	// Check validity of url.
	err := internal.ValidateUrl(url, "wayfair")
	if err != nil {
		return nil, err
	}

	// Prepare options.
	opt := &WayfairUrlOpts{}
	if len(opts) > 0 && opts[len(opts)-1] != nil {
		opt = opts[len(opts)-1]
	}

	// Set defaults.
	internal.SetDefaultUserAgent(&opt.UserAgent)

	// Check validity of parameters.
	err = opt.checkParametersValidity()
	if err != nil {
		return nil, err
	}

	// Prepare payload.
	payload := map[string]interface{}{
		"source":          oxylabs.Wayfair,
		"url":             url,
		"user_agent_type": opt.UserAgent,
		"callback_url":    opt.CallbackUrl,
	}
	// Add custom parsing instructions to the payload if provided.
	customParserFlag := false
	if opt.ParseInstructions != nil {
		payload["parsing_instructions"] = &opt.ParseInstructions
		customParserFlag = true
	}

	// Marshal.
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("error marshalling payload: %v", err)
	}

	// Req.
	internalResp, err := c.C.Req(ctx, jsonPayload, opt.Parse, customParserFlag, "POST")
	if err != nil {
		return nil, err
	}

	// Map resp.
	resp := &EcommerceResp{
		Resp: *internalResp,
	}

	return resp, nil
}
