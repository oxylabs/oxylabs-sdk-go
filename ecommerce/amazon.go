package ecommerce

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/mslmio/oxylabs-sdk-go/internal"
	"github.com/mslmio/oxylabs-sdk-go/oxylabs"
)

// AmazonUrlOpts contains all the query parameters available for amazon.
type AmazonUrlOpts struct {
	UserAgent         oxylabs.UserAgent
	Render            oxylabs.Render
	CallbackUrl       string
	Parse             bool
	ParseInstructions *map[string]interface{}
	PollInterval      time.Duration
}

// checkParameterValidity checks validity of ScrapeAmazonUrl parameters.
func (opt *AmazonUrlOpts) checkParameterValidity() error {
	if !oxylabs.IsUserAgentValid(opt.UserAgent) {
		return fmt.Errorf("invalid user agent parameter: %v", opt.UserAgent)
	}

	if opt.Render != "" && !oxylabs.IsRenderValid(opt.Render) {
		return fmt.Errorf("invalid render parameter: %v", opt.Render)
	}

	return nil
}

// ScrapeAmazonUrl scrapes amazon via Oxylabs E-Commerce API with amazon as source.
func (c *EcommerceClient) ScrapeAmazonUrl(
	url string,
	opts ...*AmazonUrlOpts,
) (*EcommerceResp, error) {
	ctx, cancel := context.WithTimeout(context.Background(), internal.DefaultTimeout)
	defer cancel()

	return c.ScrapeAmazonUrlCtx(ctx, url, opts...)
}

// ScrapeAmazonUrlCtx scrapes amazon via Oxylabs E-Commerce API with amazon as source.
// The provided context allows customization of the HTTP req, including setting timeouts.
func (c *EcommerceClient) ScrapeAmazonUrlCtx(
	ctx context.Context,
	url string,
	opts ...*AmazonUrlOpts,
) (*EcommerceResp, error) {
	// Check validity of url.
	err := internal.ValidateUrl(url, "amazon")
	if err != nil {
		return nil, err
	}

	// Prepare options.
	opt := &AmazonUrlOpts{}
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

	//Prepare payload.
	payload := map[string]interface{}{
		"source":          oxylabs.AmazonUrl,
		"url":             url,
		"user_agent_type": opt.UserAgent,
		"render":          opt.Render,
		"callback_url":    opt.CallbackUrl,
		"parse":           opt.Parse,
	}

	// Add custom parsing instructions to the payload if provided.
	customParserFlag := false
	if opt.ParseInstructions != nil {
		payload["parse_instructions"] = opt.ParseInstructions
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

// AmazonSearchOpts contains all the query parameters available for amazon_search.
type AmazonSearchOpts struct {
	Domain            oxylabs.Domain
	StartPage         int
	Pages             int
	GeoLocation       string
	UserAgent         oxylabs.UserAgent
	Render            oxylabs.Render
	CallbackUrl       string
	Parse             bool
	ParseInstructions *map[string]interface{}
	Context           []func(oxylabs.ContextOption)
	PollInterval      time.Duration
}

// checkParameterValidity checks validity of ScrapeAmazonSearch parameters.
func (opt *AmazonSearchOpts) checkParameterValidity() error {
	if !oxylabs.IsUserAgentValid(opt.UserAgent) {
		return fmt.Errorf("invalid user agent parameter: %v", opt.UserAgent)
	}

	if opt.Render != "" && !oxylabs.IsRenderValid(opt.Render) {
		return fmt.Errorf("invalid render parameter: %v", opt.Render)
	}

	if opt.Pages <= 0 || opt.StartPage <= 0 {
		return fmt.Errorf("pages and start_page parameters must be greater than 0")
	}

	return nil
}

// ScrapeAmazonSearch scrapes amazon via Oxylabs E-Commerce API with amazon_search as source.
func (c *EcommerceClient) ScrapeAmazonSearch(
	query string,
	opts ...*AmazonSearchOpts,
) (*EcommerceResp, error) {
	ctx, cancel := context.WithTimeout(context.Background(), internal.DefaultTimeout)
	defer cancel()

	return c.ScrapeAmazonSearchCtx(ctx, query, opts...)
}

// ScrapeAmazonSearchCtx scrapes amazon via Oxylabs E-Commerce API with amazon_search as source.
// The provided context allows customization of the HTTP req, including setting timeouts.
func (c *EcommerceClient) ScrapeAmazonSearchCtx(
	ctx context.Context,
	query string,
	opts ...*AmazonSearchOpts,
) (*EcommerceResp, error) {
	// Prepare options.
	opt := &AmazonSearchOpts{}
	if len(opts) > 0 && opts[len(opts)-1] != nil {
		opt = opts[len(opts)-1]
	}

	// Initialize the context map and apply each provided context modifier function.
	context := make(oxylabs.ContextOption)
	for _, modifier := range opt.Context {
		modifier(context)
	}

	// Set defaults.
	internal.SetDefaultDomain(&opt.Domain)
	internal.SetDefaultUserAgent(&opt.UserAgent)
	internal.SetDefaultStartPage(&opt.StartPage)
	internal.SetDefaultPages(&opt.Pages)

	// Check validity of parameters.
	err := opt.checkParameterValidity()
	if err != nil {
		return nil, err
	}

	// Prepare payload.
	payload := map[string]interface{}{
		"source":          oxylabs.AmazonSearch,
		"domain":          opt.Domain,
		"query":           query,
		"start_page":      opt.StartPage,
		"pages":           opt.Pages,
		"geo_location":    opt.GeoLocation,
		"user_agent_type": opt.UserAgent,
		"render":          opt.Render,
		"callback_url":    opt.CallbackUrl,
		"parse":           opt.Parse,
		"context": []map[string]interface{}{
			{
				"key":   "category_id",
				"value": context["category_id"],
			},
			{
				"key":   "merchant_id",
				"value": context["merchant_id"],
			},
		},
	}

	// Add custom parsing instructions to the payload if provided.
	customParserFlag := false
	if opt.ParseInstructions != nil {
		payload["parse_instructions"] = opt.ParseInstructions
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

// AmazonProductOpts contains all the query parameters available for amazon_product.
type AmazonProductOpts struct {
	Domain            oxylabs.Domain
	GeoLocation       string
	UserAgent         oxylabs.UserAgent
	Render            oxylabs.Render
	CallbackUrl       string
	Parse             bool
	ParseInstructions *map[string]interface{}
	Context           []func(oxylabs.ContextOption)
	PollInterval      time.Duration
}

// checkParameterValidity checks validity of ScrapeAmazonProduct parameters.
func (opt *AmazonProductOpts) checkParameterValidity() error {
	if !oxylabs.IsUserAgentValid(opt.UserAgent) {
		return fmt.Errorf("invalid user agent parameter: %v", opt.UserAgent)
	}

	if opt.Render != "" && !oxylabs.IsRenderValid(opt.Render) {
		return fmt.Errorf("invalid render parameter: %v", opt.Render)
	}

	return nil
}

// ScrapeAmazonProduct scrapes amazon via Oxylabs E-Commerce API with amazon_product as source.
func (c *EcommerceClient) ScrapeAmazonProduct(
	query string,
	opts ...*AmazonProductOpts,
) (*EcommerceResp, error) {
	ctx, cancel := context.WithTimeout(context.Background(), internal.DefaultTimeout)
	defer cancel()

	return c.ScrapeAmazonProductCtx(ctx, query, opts...)
}

// ScrapeAmazonProductCtx scrapes amazon via Oxylabs E-Commerce API with amazon_product as source.
// The provided context allows customization of the HTTP req, including setting timeouts.
func (c *EcommerceClient) ScrapeAmazonProductCtx(
	ctx context.Context,
	query string,
	opts ...*AmazonProductOpts,
) (*EcommerceResp, error) {
	// Prepare options.
	opt := &AmazonProductOpts{}
	if len(opts) > 0 && opts[len(opts)-1] != nil {
		opt = opts[len(opts)-1]
	}

	// Initialize the context map and apply each provided context modifier function.
	context := make(oxylabs.ContextOption)
	for _, modifier := range opt.Context {
		modifier(context)
	}

	// Set defaults.
	internal.SetDefaultDomain(&opt.Domain)
	internal.SetDefaultUserAgent(&opt.UserAgent)

	// Check validity of parameters.
	err := opt.checkParameterValidity()
	if err != nil {
		return nil, err
	}

	// Prepare payload.
	payload := map[string]interface{}{
		"source":          oxylabs.AmazonProduct,
		"domain":          opt.Domain,
		"query":           query,
		"geo_location":    opt.GeoLocation,
		"user_agent_type": opt.UserAgent,
		"render":          opt.Render,
		"callback_url":    opt.CallbackUrl,
		"parse":           opt.Parse,
		"context": []map[string]interface{}{
			{
				"key":   "autoselect_variant",
				"value": context["autoselect_variant"],
			},
		},
	}

	// Add custom parsing instructions to the payload if provided.
	customParserFlag := false
	if opt.ParseInstructions != nil {
		payload["parse_instructions"] = opt.ParseInstructions
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

// AmazonPricingOpts contains all the query parameters available for amazon_pricing.
type AmazonPricingOpts struct {
	Domain            oxylabs.Domain
	StartPage         int
	Pages             int
	GeoLocation       string
	UserAgent         oxylabs.UserAgent
	Render            oxylabs.Render
	CallbackUrl       string
	Parse             bool
	ParseInstructions *map[string]interface{}
	PollInterval      time.Duration
}

// checkParameterValidity checks validity of ScrapeAmazonPricing parameters.
func (opt *AmazonPricingOpts) checkParameterValidity() error {
	if !oxylabs.IsUserAgentValid(opt.UserAgent) {
		return fmt.Errorf("invalid user agent parameter: %v", opt.UserAgent)
	}

	if opt.Render != "" && !oxylabs.IsRenderValid(opt.Render) {
		return fmt.Errorf("invalid render parameter: %v", opt.Render)
	}

	if opt.Pages <= 0 || opt.StartPage <= 0 {
		return fmt.Errorf("pages and start_page parameters must be greater than 0")
	}

	return nil
}

// ScrapeAmazonPricing scrapes amazon via Oxylabs E-Commerce API with amazon_pricing as source.
func (c *EcommerceClient) ScrapeAmazonPricing(
	query string,
	opts ...*AmazonPricingOpts,
) (*EcommerceResp, error) {
	ctx, cancel := context.WithTimeout(context.Background(), internal.DefaultTimeout)
	defer cancel()

	return c.ScrapeAmazonPricingCtx(ctx, query, opts...)
}

// ScrapeAmazonPricingCtx scrapes amazon via Oxylabs E-Commerce API with amazon_pricing as source.
// The provided context allows customization of the HTTP req, including setting timeouts.
func (c *EcommerceClient) ScrapeAmazonPricingCtx(
	ctx context.Context,
	query string,
	opts ...*AmazonPricingOpts,
) (*EcommerceResp, error) {
	// Prepare options.
	opt := &AmazonPricingOpts{}
	if len(opts) > 0 && opts[len(opts)-1] != nil {
		opt = opts[len(opts)-1]
	}

	// Set defaults.
	internal.SetDefaultDomain(&opt.Domain)
	internal.SetDefaultUserAgent(&opt.UserAgent)
	internal.SetDefaultStartPage(&opt.StartPage)
	internal.SetDefaultPages(&opt.Pages)

	// Check validity of parameters.
	err := opt.checkParameterValidity()
	if err != nil {
		return nil, err
	}

	// Prepare payload.
	payload := map[string]interface{}{
		"source":          oxylabs.AmazonPricing,
		"domain":          opt.Domain,
		"query":           query,
		"start_page":      opt.StartPage,
		"pages":           opt.Pages,
		"geo_location":    opt.GeoLocation,
		"user_agent_type": opt.UserAgent,
		"render":          opt.Render,
		"callback_url":    opt.CallbackUrl,
		"parse":           opt.Parse,
	}

	// Add custom parsing instructions to the payload if provided.
	customParserFlag := false
	if opt.ParseInstructions != nil {
		payload["parse_instructions"] = opt.ParseInstructions
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

// AmazonReviewsOpts contains all the query parameters available for amazon_reviews.
type AmazonReviewsOpts struct {
	Domain            oxylabs.Domain
	GeoLocation       string
	UserAgent         oxylabs.UserAgent
	StartPage         int
	Pages             int
	Render            oxylabs.Render
	CallbackUrl       string
	Parse             bool
	ParseInstructions *map[string]interface{}
	PollInterval      time.Duration
}

// checkParameterValidity checks validity of ScrapeAmazonReviews parameters.
func (opt *AmazonReviewsOpts) checkParameterValidity() error {
	if !oxylabs.IsUserAgentValid(opt.UserAgent) {
		return fmt.Errorf("invalid user agent parameter: %v", opt.UserAgent)
	}

	if opt.Render != "" && !oxylabs.IsRenderValid(opt.Render) {
		return fmt.Errorf("invalid render parameter: %v", opt.Render)
	}

	if opt.Pages <= 0 || opt.StartPage <= 0 {
		return fmt.Errorf("pages and start_page parameters must be greater than 0")
	}

	return nil
}

// ScrapeAmazonReviews scrapes amazon via Oxylabs E-Commerce API with amazon_reviews as source.
func (c *EcommerceClient) ScrapeAmazonReviews(
	query string,
	opts ...*AmazonReviewsOpts,
) (*EcommerceResp, error) {
	ctx, cancel := context.WithTimeout(context.Background(), internal.DefaultTimeout)
	defer cancel()

	return c.ScrapeAmazonReviewsCtx(ctx, query, opts...)
}

// ScrapeAmazonReviewsCtx scrapes amazon via Oxylabs E-Commerce API with amazon_reviews as source.
// The provided context allows customization of the HTTP req, including setting timeouts.
func (c *EcommerceClient) ScrapeAmazonReviewsCtx(
	ctx context.Context,
	query string,
	opts ...*AmazonReviewsOpts,
) (*EcommerceResp, error) {
	// Prepare options.
	opt := &AmazonReviewsOpts{}
	if len(opts) > 0 && opts[len(opts)-1] != nil {
		opt = opts[len(opts)-1]
	}

	// Set defaults.
	internal.SetDefaultDomain(&opt.Domain)
	internal.SetDefaultUserAgent(&opt.UserAgent)
	internal.SetDefaultStartPage(&opt.StartPage)
	internal.SetDefaultPages(&opt.Pages)

	// Check validity of parameters.
	err := opt.checkParameterValidity()
	if err != nil {
		return nil, err
	}

	// Prepare payload.
	payload := map[string]interface{}{
		"source":          oxylabs.AmazonReviews,
		"domain":          opt.Domain,
		"query":           query,
		"geo_location":    opt.GeoLocation,
		"user_agent_type": opt.UserAgent,
		"start_page":      opt.StartPage,
		"pages":           opt.Pages,
		"render":          opt.Render,
		"callback_url":    opt.CallbackUrl,
		"parse":           opt.Parse,
	}

	// Add custom parsing instructions to the payload if provided.
	customParserFlag := false
	if opt.ParseInstructions != nil {
		payload["parse_instructions"] = opt.ParseInstructions
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

// AmazonQuestionsOpts contains all the query parameters available for amazon_questions.
type AmazonQuestionsOpts struct {
	Domain            oxylabs.Domain
	GeoLocation       string
	UserAgent         oxylabs.UserAgent
	Render            oxylabs.Render
	CallbackUrl       string
	Parse             bool
	ParseInstructions *map[string]interface{}
	PollInterval      time.Duration
}

// checkParameterValidity checks validity of ScrapeAmazonQuestions parameters.
func (opt *AmazonQuestionsOpts) checkParameterValidity() error {
	if !oxylabs.IsUserAgentValid(opt.UserAgent) {
		return fmt.Errorf("invalid user agent parameter: %v", opt.UserAgent)
	}

	if opt.Render != "" && !oxylabs.IsRenderValid(opt.Render) {
		return fmt.Errorf("invalid render parameter: %v", opt.Render)
	}

	return nil
}

// ScrapeAmazonQuestions scrapes amazon via Oxylabs E-Commerce API with amazon_questions as source.
func (c *EcommerceClient) ScrapeAmazonQuestions(
	query string,
	opts ...*AmazonQuestionsOpts,
) (*EcommerceResp, error) {
	ctx, cancel := context.WithTimeout(context.Background(), internal.DefaultTimeout)
	defer cancel()

	return c.ScrapeAmazonQuestionsCtx(ctx, query, opts...)
}

// ScrapeAmazonQuestionsCtx scrapes amazon via Oxylabs E-Commerce API with amazon_questions as source.
// The provided context allows customization of the HTTP req, including setting timeouts.
func (c *EcommerceClient) ScrapeAmazonQuestionsCtx(
	ctx context.Context,
	query string,
	opts ...*AmazonQuestionsOpts,
) (*EcommerceResp, error) {
	// Prepare options.
	opt := &AmazonQuestionsOpts{}
	if len(opts) > 0 && opts[len(opts)-1] != nil {
		opt = opts[len(opts)-1]
	}

	// Set defaults.
	internal.SetDefaultDomain(&opt.Domain)
	internal.SetDefaultUserAgent(&opt.UserAgent)

	// Check validity of parameters.
	err := opt.checkParameterValidity()
	if err != nil {
		return nil, err
	}

	// Prepare payload.
	payload := map[string]interface{}{
		"source":          oxylabs.AmazonQuestions,
		"domain":          opt.Domain,
		"query":           query,
		"geo_location":    opt.GeoLocation,
		"user_agent_type": opt.UserAgent,
		"render":          opt.Render,
		"callback_url":    opt.CallbackUrl,
		"parse":           opt.Parse,
	}

	// Add custom parsing instructions to the payload if provided.
	customParserFlag := false
	if opt.ParseInstructions != nil {
		payload["parse_instructions"] = opt.ParseInstructions
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

// AmazonBestsellersOpts contains all the query parameters available for amazon_bestsellers.
type AmazonBestsellersOpts struct {
	Domain            oxylabs.Domain
	StartPage         int
	Pages             int
	GeoLocation       string
	UserAgent         oxylabs.UserAgent
	Render            oxylabs.Render
	CallbackUrl       string
	Parse             bool
	ParseInstructions *map[string]interface{}
	PollInterval      time.Duration
}

// checkParameterValidity checks validity of ScrapeAmazonBestsellers parameters.
func (opt *AmazonBestsellersOpts) checkParameterValidity() error {
	if !oxylabs.IsUserAgentValid(opt.UserAgent) {
		return fmt.Errorf("invalid user agent parameter: %v", opt.UserAgent)
	}

	if opt.Render != "" && !oxylabs.IsRenderValid(opt.Render) {
		return fmt.Errorf("invalid render parameter: %v", opt.Render)
	}

	if opt.Pages <= 0 || opt.StartPage <= 0 {
		return fmt.Errorf("pages and start_page parameters must be greater than 0")
	}

	return nil
}

// ScrapeAmazonBestsellers scrapes amazon via Oxylabs E-Commerce API with amazon_bestsellers as source.
func (c *EcommerceClient) ScrapeAmazonBestsellers(
	query string,
	opts ...*AmazonBestsellersOpts,
) (*EcommerceResp, error) {
	ctx, cancel := context.WithTimeout(context.Background(), internal.DefaultTimeout)
	defer cancel()

	return c.ScrapeAmazonBestsellersCtx(ctx, query, opts...)
}

// ScrapeAmazonBestsellersCtx scrapes amazon via Oxylabs E-Commerce API with amazon_bestsellers as source.
// The provided context allows customization of the HTTP req, including setting timeouts.
func (c *EcommerceClient) ScrapeAmazonBestsellersCtx(
	ctx context.Context,
	query string,
	opts ...*AmazonBestsellersOpts,
) (*EcommerceResp, error) {
	// Prepare options.
	opt := &AmazonBestsellersOpts{}
	if len(opts) > 0 && opts[len(opts)-1] != nil {
		opt = opts[len(opts)-1]
	}

	// Set defaults.
	internal.SetDefaultDomain(&opt.Domain)
	internal.SetDefaultUserAgent(&opt.UserAgent)
	internal.SetDefaultStartPage(&opt.StartPage)
	internal.SetDefaultPages(&opt.Pages)

	// Check validity of parameters.
	err := opt.checkParameterValidity()
	if err != nil {
		return nil, err
	}

	// Prepare payload.
	payload := map[string]interface{}{
		"source":          oxylabs.AmazonBestsellers,
		"domain":          opt.Domain,
		"query":           query,
		"start_page":      opt.StartPage,
		"pages":           opt.Pages,
		"geo_location":    opt.GeoLocation,
		"user_agent_type": opt.UserAgent,
		"render":          opt.Render,
		"callback_url":    opt.CallbackUrl,
		"parse":           opt.Parse,
	}

	// Add custom parsing instructions to the payload if provided.
	customParserFlag := false
	if opt.ParseInstructions != nil {
		payload["parse_instructions"] = opt.ParseInstructions
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

// AmazonSellersOpts contains all the query parameters available for amazon_seller.
type AmazonSellersOpts struct {
	Domain            oxylabs.Domain
	GeoLocation       string
	UserAgent         oxylabs.UserAgent
	Render            oxylabs.Render
	CallbackUrl       string
	Parse             bool
	ParseInstructions *map[string]interface{}
	PollInterval      time.Duration
}

// checkParameterValidity checks validity of ScrapeAmazonSeller parameters.
func (opt *AmazonSellersOpts) checkParameterValidity() error {
	if !oxylabs.IsUserAgentValid(opt.UserAgent) {
		return fmt.Errorf("invalid user agent parameter: %v", opt.UserAgent)
	}

	if opt.Render != "" && !oxylabs.IsRenderValid(opt.Render) {
		return fmt.Errorf("invalid render parameter: %v", opt.Render)
	}

	return nil
}

// ScrapeAmazonSellers scrapes amazon via Oxylabs E-Commerce API with amazon_seller as source.
func (c *EcommerceClient) ScrapeAmazonSellers(
	query string,
	opts ...*AmazonSellersOpts,
) (*EcommerceResp, error) {
	ctx, cancel := context.WithTimeout(context.Background(), internal.DefaultTimeout)
	defer cancel()

	return c.ScrapeAmazonSellersCtx(ctx, query, opts...)
}

// ScrapeAmazonSellerCtx scrapes amazon via Oxylabs E-Commerce API with amazon_seller as source.
// The provided context allows customization of the HTTP req, including setting timeouts.
func (c *EcommerceClient) ScrapeAmazonSellersCtx(
	ctx context.Context,
	query string,
	opts ...*AmazonSellersOpts,
) (*EcommerceResp, error) {
	// Prepare options.
	opt := &AmazonSellersOpts{}
	if len(opts) > 0 && opts[len(opts)-1] != nil {
		opt = opts[len(opts)-1]
	}

	// Set defaults.
	internal.SetDefaultDomain(&opt.Domain)
	internal.SetDefaultUserAgent(&opt.UserAgent)

	// Check validity of parameters.
	err := opt.checkParameterValidity()
	if err != nil {
		return nil, err
	}

	// Prepare payload.
	payload := map[string]interface{}{
		"source":          oxylabs.AmazonSellers,
		"domain":          opt.Domain,
		"query":           query,
		"geo_location":    opt.GeoLocation,
		"user_agent_type": opt.UserAgent,
		"render":          opt.Render,
		"callback_url":    opt.CallbackUrl,
		"parse":           opt.Parse,
	}

	// Add custom parsing instructions to the payload if provided.
	customParserFlag := false
	if opt.ParseInstructions != nil {
		payload["parse_instructions"] = opt.ParseInstructions
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
