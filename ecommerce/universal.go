package ecommerce

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/mslmio/oxylabs-sdk-go/internal"
	"github.com/mslmio/oxylabs-sdk-go/oxylabs"
)

// UniversalUrlOpts contains all the query parameters available for universal url scrape.
type UniversalUrlOpts struct {
	UserAgent         oxylabs.UserAgent
	CallbackUrl       string
	GeoLocation       string
	Locale            oxylabs.Locale
	Render            oxylabs.Render
	ContentEncoding   string
	Context           []func(oxylabs.ContextOption)
	CallbackURL       string
	Parse             bool
	ParserType        interface{}
	ParseInstructions *map[string]interface{}
	PollInterval      time.Duration
}

// checkParameterValidity checks validity of UniversalUrlOpts parameters.
func (opt *UniversalUrlOpts) checkParametersValidity(ctx oxylabs.ContextOption) error {
	if !oxylabs.IsUserAgentValid(opt.UserAgent) {
		return fmt.Errorf("invalid user agent parameter: %v", opt.UserAgent)
	}

	if opt.Render != "" && !oxylabs.IsRenderValid(opt.Render) {
		return fmt.Errorf("invalid render parameter: %v", opt.Render)
	}

	if ctx["http_method"] != "post" && ctx["http_method"] != "get" {
		return fmt.Errorf("invalid http method")
	}

	if ctx["content"] != nil && ctx["http_method"] != "post" {
		return fmt.Errorf("content is useful only if http method is post")
	}

	return nil
}

// ScrapeUniversalUrl scrapes all urls via Oxylabs E-Commerce API with universal_ecommerce as source.
func (c *EcommerceClient) ScrapeUniversalUrl(
	url string,
	opts ...*UniversalUrlOpts,
) (*EcommerceResp, error) {
	ctx, cancel := context.WithTimeout(context.Background(), internal.DefaultTimeout)
	defer cancel()

	return c.ScrapeUniversalUrlCtx(ctx, url, opts...)
}

// ScrapeUniversalUrlCtx scrapes all urls via Oxylabs E-Commerce API with universal_ecommerce as source.
// The provided context allows customization of the HTTP req, including setting timeouts.
func (c *EcommerceClient) ScrapeUniversalUrlCtx(
	ctx context.Context,
	url string,
	opts ...*UniversalUrlOpts,
) (*EcommerceResp, error) {
	// Prepare options.
	opt := &UniversalUrlOpts{}
	if len(opts) > 0 && opts[len(opts)-1] != nil {
		opt = opts[len(opts)-1]
	}

	// Initialize the context map and apply each provided context modifier function.
	context := make(oxylabs.ContextOption)
	for _, modifier := range opt.Context {
		modifier(context)
	}

	// Set defaults.
	internal.SetDefaultHttpMethod(context)
	internal.SetDefaultUserAgent(&opt.UserAgent)
	internal.SetDefaultContentEncoding(&opt.ContentEncoding)

	// Check validity of parameters.
	err := opt.checkParametersValidity(context)
	if err != nil {
		return nil, err
	}

	// Prepare payload.
	payload := map[string]interface{}{
		"source":           oxylabs.Universal,
		"url":              url,
		"user_agent_type":  opt.UserAgent,
		"geo_location":     opt.GeoLocation,
		"locale":           opt.Locale,
		"render":           opt.Render,
		"content_encoding": opt.ContentEncoding,
		"context": []map[string]interface{}{
			{
				"key":   "content",
				"value": context["content"],
			},
			{
				"key":   "cookies",
				"value": context["cookies"],
			},
			{
				"key":   "follow_redirects",
				"value": context["follow_redirects"],
			},
			{
				"key":   "headers",
				"value": context["headers"],
			},
			{
				"key":   "http_method",
				"value": context["http_method"],
			},
			{
				"key":   "session_id",
				"value": context["session_id"],
			},
			{
				"key":   "successful_status_codes",
				"value": context["successful_status_codes"],
			},
		},
		"callback_url": opt.CallbackUrl,
		"parse":        opt.Parse,
		"parser_type":  opt.ParserType,
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
