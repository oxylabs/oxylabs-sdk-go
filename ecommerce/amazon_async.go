package ecommerce

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/oxylabs/oxylabs-sdk-go/internal"
	"github.com/oxylabs/oxylabs-sdk-go/oxylabs"
)

// ScrapeAmazonUrl scrapes amazon via Oxylabs E-Commerce API with amazon as source.
func (c *EcommerceClientAsync) ScrapeAmazonUrl(
	url string,
	opts ...*AmazonUrlOpts,
) (chan *Resp, error) {
	ctx, cancel := context.WithTimeout(context.Background(), internal.DefaultTimeout)
	defer cancel()

	return c.ScrapeAmazonUrlCtx(ctx, url, opts...)
}

// ScrapeAmazonUrlCtx scrapes amazon via Oxylabs E-Commerce API with amazon as source.
// The provided context allows customization of the HTTP req, including setting timeouts.
func (c *EcommerceClientAsync) ScrapeAmazonUrlCtx(
	ctx context.Context,
	url string,
	opts ...*AmazonUrlOpts,
) (chan *Resp, error) {
	errChan := make(chan error)
	httpRespChan := make(chan *http.Response)
	respChan := make(chan *Resp)

	/// Check validity of url.
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

	// Get job ID.
	jobID, err := c.C.GetJobID(jsonPayload)
	if err != nil {
		return nil, err
	}

	// Poll job status.
	go c.C.PollJobStatus(
		ctx,
		jobID,
		opt.PollInterval,
		httpRespChan,
		errChan,
	)

	// Handle error.
	err = <-errChan
	if err != nil {
		return nil, err
	}

	// Unmarshal the http Response and get the response.
	httpResp := <-httpRespChan
	resp, err := GetResp(httpResp, opt.Parse, customParserFlag)
	if err != nil {
		return nil, err
	}

	// Retrieve internal resp and forward it to the
	// resp channel.
	go func() {
		respChan <- resp
	}()

	return respChan, nil
}

// ScrapeAmazonSearch scrapes amazon via Oxylabs E-Commerce API with amazon_search as source.
func (c *EcommerceClientAsync) ScrapeAmazonSearch(
	query string,
	opts ...*AmazonSearchOpts,
) (chan *Resp, error) {
	ctx, cancel := context.WithTimeout(context.Background(), internal.DefaultTimeout)
	defer cancel()

	return c.ScrapeAmazonSearchCtx(ctx, query, opts...)
}

// ScrapeAmazonSearchCtx scrapes amazon via Oxylabs E-Commerce API with amazon_search as source.
// The provided context allows customization of the HTTP req, including setting timeouts.
func (c *EcommerceClientAsync) ScrapeAmazonSearchCtx(
	ctx context.Context,
	query string,
	opts ...*AmazonSearchOpts,
) (chan *Resp, error) {
	errChan := make(chan error)
	httpRespChan := make(chan *http.Response)
	respChan := make(chan *Resp)

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

	// Get job ID.
	jobID, err := c.C.GetJobID(jsonPayload)
	if err != nil {
		return nil, err
	}

	// Poll job status.
	go c.C.PollJobStatus(
		ctx,
		jobID,
		opt.PollInterval,
		httpRespChan,
		errChan,
	)

	// Handle error.
	err = <-errChan
	if err != nil {
		return nil, err
	}

	// Unmarshal the http Response and get the response.
	httpResp := <-httpRespChan
	resp, err := GetResp(httpResp, opt.Parse, customParserFlag)
	if err != nil {
		return nil, err
	}

	// Retrieve internal resp and forward it to the
	// resp channel.
	go func() {
		respChan <- resp
	}()

	return respChan, nil
}

// ScrapeAmazonProduct scrapes amazon via Oxylabs E-Commerce API with amazon_product as source.
func (c *EcommerceClientAsync) ScrapeAmazonProduct(
	query string,
	opts ...*AmazonProductOpts,
) (chan *Resp, error) {
	ctx, cancel := context.WithTimeout(context.Background(), internal.DefaultTimeout)
	defer cancel()

	return c.ScrapeAmazonProductCtx(ctx, query, opts...)
}

// ScrapeAmazonProductCtx scrapes amazon via Oxylabs E-Commerce API with amazon_product as source.
// The provided context allows customization of the HTTP req, including setting timeouts.
func (c *EcommerceClientAsync) ScrapeAmazonProductCtx(
	ctx context.Context,
	query string,
	opts ...*AmazonProductOpts,
) (chan *Resp, error) {
	errChan := make(chan error)
	httpRespChan := make(chan *http.Response)
	respChan := make(chan *Resp)

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

	// Get job ID.
	jobID, err := c.C.GetJobID(jsonPayload)
	if err != nil {
		return nil, err
	}

	// Poll job status.
	go c.C.PollJobStatus(
		ctx,
		jobID,
		opt.PollInterval,
		httpRespChan,
		errChan,
	)

	// Handle error.
	err = <-errChan
	if err != nil {
		return nil, err
	}

	// Unmarshal the http Response and get the response.
	httpResp := <-httpRespChan
	resp, err := GetResp(httpResp, opt.Parse, customParserFlag)
	if err != nil {
		return nil, err
	}

	// Retrieve internal resp and forward it to the
	// resp channel.
	go func() {
		respChan <- resp
	}()

	return respChan, nil
}

// ScrapeAmazonPricing scrapes amazon via Oxylabs E-Commerce API with amazon_pricing as source.
func (c *EcommerceClientAsync) ScrapeAmazonPricing(
	query string,
	opts ...*AmazonPricingOpts,
) (chan *Resp, error) {
	ctx, cancel := context.WithTimeout(context.Background(), internal.DefaultTimeout)
	defer cancel()

	return c.ScrapeAmazonPricingCtx(ctx, query, opts...)
}

// ScrapeAmazonPricingCtx scrapes amazon via Oxylabs E-Commerce API with amazon_pricing as source.
// The provided context allows customization of the HTTP req, including setting timeouts.
func (c *EcommerceClientAsync) ScrapeAmazonPricingCtx(
	ctx context.Context,
	query string,
	opts ...*AmazonPricingOpts,
) (chan *Resp, error) {
	errChan := make(chan error)
	httpRespChan := make(chan *http.Response)
	respChan := make(chan *Resp)

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

	// Get job ID.
	jobID, err := c.C.GetJobID(jsonPayload)
	if err != nil {
		return nil, err
	}

	// Poll job status.
	go c.C.PollJobStatus(
		ctx,
		jobID,
		opt.PollInterval,
		httpRespChan,
		errChan,
	)

	// Handle error.
	err = <-errChan
	if err != nil {
		return nil, err
	}

	// Unmarshal the http Response and get the response.
	httpResp := <-httpRespChan
	resp, err := GetResp(httpResp, opt.Parse, customParserFlag)
	if err != nil {
		return nil, err
	}

	// Retrieve internal resp and forward it to the
	// resp channel.
	go func() {
		respChan <- resp
	}()

	return respChan, nil
}

// ScrapeAmazonReviews scrapes amazon via Oxylabs E-Commerce API with amazon_reviews as source.
func (c *EcommerceClientAsync) ScrapeAmazonReviews(
	query string,
	opts ...*AmazonReviewsOpts,
) (chan *Resp, error) {
	ctx, cancel := context.WithTimeout(context.Background(), internal.DefaultTimeout)
	defer cancel()

	return c.ScrapeAmazonReviewsCtx(ctx, query, opts...)
}

// ScrapeAmazonReviewsCtx scrapes amazon via Oxylabs E-Commerce API with amazon_reviews as source.
// The provided context allows customization of the HTTP req, including setting timeouts.
func (c *EcommerceClientAsync) ScrapeAmazonReviewsCtx(
	ctx context.Context,
	query string,
	opts ...*AmazonReviewsOpts,
) (chan *Resp, error) {
	errChan := make(chan error)
	httpRespChan := make(chan *http.Response)
	respChan := make(chan *Resp)

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

	// Get job ID.
	jobID, err := c.C.GetJobID(jsonPayload)
	if err != nil {
		return nil, err
	}

	// Poll job status.
	go c.C.PollJobStatus(
		ctx,
		jobID,
		opt.PollInterval,
		httpRespChan,
		errChan,
	)

	// Handle error.
	err = <-errChan
	if err != nil {
		return nil, err
	}

	// Unmarshal the http Response and get the response.
	httpResp := <-httpRespChan
	resp, err := GetResp(httpResp, opt.Parse, customParserFlag)
	if err != nil {
		return nil, err
	}

	// Retrieve internal resp and forward it to the
	// resp channel.
	go func() {
		respChan <- resp
	}()

	return respChan, nil
}

// ScrapeAmazonQuestions scrapes amazon via Oxylabs E-Commerce API with amazon_questions as source.
func (c *EcommerceClientAsync) ScrapeAmazonQuestions(
	query string,
	opts ...*AmazonQuestionsOpts,
) (chan *Resp, error) {
	ctx, cancel := context.WithTimeout(context.Background(), internal.DefaultTimeout)
	defer cancel()

	return c.ScrapeAmazonQuestionsCtx(ctx, query, opts...)
}

// ScrapeAmazonQuestionsCtx scrapes amazon via Oxylabs E-Commerce API with amazon_questions as source.
// The provided context allows customization of the HTTP req, including setting timeouts.
func (c *EcommerceClientAsync) ScrapeAmazonQuestionsCtx(
	ctx context.Context,
	query string,
	opts ...*AmazonQuestionsOpts,
) (chan *Resp, error) {
	errChan := make(chan error)
	httpRespChan := make(chan *http.Response)
	respChan := make(chan *Resp)

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

	// Get job ID.
	jobID, err := c.C.GetJobID(jsonPayload)
	if err != nil {
		return nil, err
	}

	// Poll job status.
	go c.C.PollJobStatus(
		ctx,
		jobID,
		opt.PollInterval,
		httpRespChan,
		errChan,
	)

	// Handle error.
	err = <-errChan
	if err != nil {
		return nil, err
	}

	// Unmarshal the http Response and get the response.
	httpResp := <-httpRespChan
	resp, err := GetResp(httpResp, opt.Parse, customParserFlag)
	if err != nil {
		return nil, err
	}

	// Retrieve internal resp and forward it to the
	// resp channel.
	go func() {
		respChan <- resp
	}()

	return respChan, nil
}

// ScrapeAmazonBestSellers scrapes amazon via Oxylabs E-Commerce API with amazon_bestsellers as source.
func (c *EcommerceClientAsync) ScrapeAmazonBestsellers(
	query string,
	opts ...*AmazonBestsellersOpts,
) (chan *Resp, error) {
	ctx, cancel := context.WithTimeout(context.Background(), internal.DefaultTimeout)
	defer cancel()

	return c.ScrapeAmazonBestsellersCtx(ctx, query, opts...)
}

// ScrapeAmazonBestsellersCtx scrapes amazon via Oxylabs E-Commerce API with amazon_bestsellers as source.
// The provided context allows customization of the HTTP req, including setting timeouts.
func (c *EcommerceClientAsync) ScrapeAmazonBestsellersCtx(
	ctx context.Context,
	query string,
	opts ...*AmazonBestsellersOpts,
) (chan *Resp, error) {
	errChan := make(chan error)
	httpRespChan := make(chan *http.Response)
	respChan := make(chan *Resp)

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

	// Get job ID.
	jobID, err := c.C.GetJobID(jsonPayload)
	if err != nil {
		return nil, err
	}

	// Poll job status.
	go c.C.PollJobStatus(
		ctx,
		jobID,
		opt.PollInterval,
		httpRespChan,
		errChan,
	)

	// Handle error.
	err = <-errChan
	if err != nil {
		return nil, err
	}

	// Unmarshal the http Response and get the response.
	httpResp := <-httpRespChan
	resp, err := GetResp(httpResp, opt.Parse, customParserFlag)
	if err != nil {
		return nil, err
	}

	// Retrieve internal resp and forward it to the
	// resp channel.
	go func() {
		respChan <- resp
	}()

	return respChan, nil
}

// ScrapeAmazonSellers scrapes amazon via Oxylabs E-Commerce API with amazon_sellers as source.
func (c *EcommerceClientAsync) ScrapeAmazonSellers(
	query string,
	opts ...*AmazonSellersOpts,
) (chan *Resp, error) {
	ctx, cancel := context.WithTimeout(context.Background(), internal.DefaultTimeout)
	defer cancel()

	return c.ScrapeAmazonSellersCtx(ctx, query, opts...)
}

// ScrapeAmazonSellersCtx scrapes amazon via Oxylabs E-Commerce API with amazon_sellers as source.
// The provided context allows customization of the HTTP req, including setting timeouts.
func (c *EcommerceClientAsync) ScrapeAmazonSellersCtx(
	ctx context.Context,
	query string,
	opts ...*AmazonSellersOpts,
) (chan *Resp, error) {
	errChan := make(chan error)
	httpRespChan := make(chan *http.Response)
	respChan := make(chan *Resp)

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

	// Get job ID.
	jobID, err := c.C.GetJobID(jsonPayload)
	if err != nil {
		return nil, err
	}

	// Poll job status.
	go c.C.PollJobStatus(
		ctx,
		jobID,
		opt.PollInterval,
		httpRespChan,
		errChan,
	)

	// Handle error.
	err = <-errChan
	if err != nil {
		return nil, err
	}

	// Unmarshal the http Response and get the response.
	httpResp := <-httpRespChan
	resp, err := GetResp(httpResp, opt.Parse, customParserFlag)
	if err != nil {
		return nil, err
	}

	// Retrieve internal resp and forward it to the
	// resp channel.
	go func() {
		respChan <- resp
	}()

	return respChan, nil
}
