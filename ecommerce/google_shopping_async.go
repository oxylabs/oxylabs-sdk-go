package ecommerce

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mslmio/oxylabs-sdk-go/internal"
	"github.com/mslmio/oxylabs-sdk-go/oxylabs"
)

// ScrapeGoogleShoppingUrl scrapes google shopping with async polling runtime
// via Oxylabs E-Commerce API and google_shopping as source.
func (c *EcommerceClientAsync) ScrapeGoogleShoppingUrl(
	url string,
	opts ...*GoogleShoppingUrlOpts,
) (chan *EcommerceResp, error) {
	ctx, cancel := context.WithTimeout(context.Background(), internal.DefaultTimeout)
	defer cancel()

	return c.ScrapeGoogleShoppingUrlCtx(ctx, url, opts...)
}

// ScrapeGoogleShoppingUrlCtx scrapes google shopping with async polling runtime
// via Oxylabs E-Commerce API and google_shopping as source.
// The provided context allows customization of the HTTP req, including setting timeouts.
func (c *EcommerceClientAsync) ScrapeGoogleShoppingUrlCtx(
	ctx context.Context,
	url string,
	opts ...*GoogleShoppingUrlOpts,
) (chan *EcommerceResp, error) {
	errChan := make(chan error)
	respChan := make(chan *EcommerceResp)
	internalRespChan := make(chan *internal.Resp)

	// Check validity of url.
	err := internal.ValidateUrl(url, "shopping.google")
	if err != nil {
		return nil, err
	}

	// Prepare options.
	opt := &GoogleShoppingUrlOpts{}
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
		"source":          oxylabs.GoogleShoppingUrl,
		"url":             url,
		"user_agent_type": opt.UserAgent,
		"render":          opt.Render,
		"callback_url":    opt.CallbackUrl,
		"geo_location":    opt.GeoLocation,
		"parse":           opt.Parse,
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

	// Get job ID.
	jobID, err := c.C.GetJobID(jsonPayload)
	if err != nil {
		return nil, err
	}

	// Poll job status.
	go c.C.PollJobStatus(
		ctx,
		jobID,
		opt.Parse,
		customParserFlag,
		opt.PollInterval,
		internalRespChan,
		errChan,
	)

	// Handle error.
	err = <-errChan
	if err != nil {
		return nil, err
	}

	// Retrieve internal resp and forward it to the
	// external resp channel.
	internalResp := <-internalRespChan
	go func() {
		respChan <- &EcommerceResp{
			Resp: *internalResp,
		}
	}()

	return respChan, nil
}

// ScrapeGoogleShoppingSearch scrapes google shopping with async polling runtime
// via Oxylabs E-Commerce API and google_shopping_search as source.
func (c *EcommerceClientAsync) ScrapeGoogleShoppingSearch(
	query string,
	opts ...*GoogleShoppingSearchOpts,
) (chan *EcommerceResp, error) {
	ctx, cancel := context.WithTimeout(context.Background(), internal.DefaultTimeout)
	defer cancel()

	return c.ScrapeGoogleShoppingSearchCtx(ctx, query, opts...)
}

// ScrapeGoogleShoppingSearchCtx scrapes google shopping with async polling runtime
// via Oxylabs E-Commerce API and google_shopping_search as source.
// The provided context allows customization of the HTTP req, including setting timeouts.
func (c *EcommerceClientAsync) ScrapeGoogleShoppingSearchCtx(
	ctx context.Context,
	query string,
	opts ...*GoogleShoppingSearchOpts,
) (chan *EcommerceResp, error) {
	errChan := make(chan error)
	respChan := make(chan *EcommerceResp)
	internalRespChan := make(chan *internal.Resp)

	// Prepare options.
	opt := &GoogleShoppingSearchOpts{}
	if len(opts) > 0 && opts[len(opts)-1] != nil {
		opt = opts[len(opts)-1]
	}

	// Initialize the context map and apply each provided context modifier function.
	context := make(oxylabs.ContextOption)
	for _, modifier := range opt.Context {
		modifier(context)
	}

	// Set defaults.
	internal.SetDefaultSortBy(context)
	internal.SetDefaultPages(&opt.Pages)
	internal.SetDefaultDomain(&opt.Domain)
	internal.SetDefaultStartPage(&opt.StartPage)
	internal.SetDefaultUserAgent(&opt.UserAgent)

	// Check validity of parameters.
	err := opt.checkParameterValidity(context)
	if err != nil {
		return nil, err
	}

	// Prepare payload with common parameters.
	payload := map[string]interface{}{
		"source":           oxylabs.GoogleShoppingSearch,
		"domain":           opt.Domain,
		"query":            query,
		"start_page":       opt.StartPage,
		"pages":            opt.Pages,
		"locale":           opt.Locale,
		"results_language": opt.ResultsLanguage,
		"geo_location":     opt.GeoLocation,
		"user_agent_type":  opt.UserAgent,
		"render":           opt.Render,
		"callback_url":     opt.CallbackURL,
		"parse":            opt.Parse,
		"context": []map[string]interface{}{
			{
				"key":   "nfpr",
				"value": context["nfpr"],
			},
			{
				"key":   "sort_by",
				"value": context["sort_by"],
			},
			{
				"key":   "min_price",
				"value": context["min_price"],
			},
			{
				"key":   "max_price",
				"value": context["max_price"],
			},
		},
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

	// Get job ID.
	jobID, err := c.C.GetJobID(jsonPayload)
	if err != nil {
		return nil, err
	}

	// Poll job status.
	go c.C.PollJobStatus(
		ctx,
		jobID,
		opt.Parse,
		customParserFlag,
		opt.PollInterval,
		internalRespChan,
		errChan,
	)

	// Handle error.
	err = <-errChan
	if err != nil {
		return nil, err
	}

	// Retrieve internal resp and forward it to the
	// external resp channel.
	internalResp := <-internalRespChan
	go func() {
		respChan <- &EcommerceResp{
			Resp: *internalResp,
		}
	}()

	return respChan, nil
}

// ScrapeGoogleShoppingProduct scrapes google shopping with async polling runtime
// via Oxylabs E-Commerce API with google_shopping_product as source.
func (c *EcommerceClientAsync) ScrapeGoogleShoppingProduct(
	query string,
	opts ...*GoogleShoppingProductOpts,
) (chan *EcommerceResp, error) {
	ctx, cancel := context.WithTimeout(context.Background(), internal.DefaultTimeout)
	defer cancel()

	return c.ScrapeGoogleShoppingProductCtx(ctx, query, opts...)
}

// ScrapeGoogleShoppingProductCtx scrapes google shopping with async polling runtime
// via Oxylabs E-Commerce API and google_shopping_product as source.
// The provided context allows customization of the HTTP req, including setting timeouts.
func (c *EcommerceClientAsync) ScrapeGoogleShoppingProductCtx(
	ctx context.Context,
	query string,
	opts ...*GoogleShoppingProductOpts,
) (chan *EcommerceResp, error) {
	errChan := make(chan error)
	respChan := make(chan *EcommerceResp)
	internalRespChan := make(chan *internal.Resp)

	// Prepare options.
	opt := &GoogleShoppingProductOpts{}
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

	// Prepare payload with common parameters.
	payload := map[string]interface{}{
		"source":           oxylabs.GoogleShoppingProduct,
		"domain":           opt.Domain,
		"query":            query,
		"locale":           opt.Locale,
		"results_language": opt.ResultsLanguage,
		"geo_location":     opt.GeoLocation,
		"user_agent_type":  opt.UserAgent,
		"render":           opt.Render,
		"callback_url":     opt.CallbackURL,
		"parse":            opt.Parse,
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

	// Get job ID.
	jobID, err := c.C.GetJobID(jsonPayload)
	if err != nil {
		return nil, err
	}

	// Poll job status.
	go c.C.PollJobStatus(
		ctx,
		jobID,
		opt.Parse,
		customParserFlag,
		opt.PollInterval,
		internalRespChan,
		errChan,
	)

	// Handle error.
	err = <-errChan
	if err != nil {
		return nil, err
	}

	// Retrieve internal resp and forward it to the
	// external resp channel.
	internalResp := <-internalRespChan
	go func() {
		respChan <- &EcommerceResp{
			Resp: *internalResp,
		}
	}()

	return respChan, nil
}

// ScrapeGoogleShoppingPricing scrapes google shopping with async polling runtime
// via Oxylabs E-Commerce API and google_shopping_pricing as source.
func (c *EcommerceClientAsync) ScrapeGoogleShoppingPricing(
	query string,
	opts ...*GoogleShoppingPricingOpts,
) (chan *EcommerceResp, error) {
	ctx, cancel := context.WithTimeout(context.Background(), internal.DefaultTimeout)
	defer cancel()

	return c.ScrapeGoogleShoppingPricingCtx(ctx, query, opts...)
}

// ScrapeGoogleShoppingPricingCtx scrapes google shopping via Oxylabs E-Commerce API
// with google_shopping_pricing as source.
// The provided context allows customization of the HTTP req, including setting timeouts.
func (c *EcommerceClientAsync) ScrapeGoogleShoppingPricingCtx(
	ctx context.Context,
	query string,
	opts ...*GoogleShoppingPricingOpts,
) (chan *EcommerceResp, error) {
	errChan := make(chan error)
	respChan := make(chan *EcommerceResp)
	internalRespChan := make(chan *internal.Resp)

	// Prepare options.
	opt := &GoogleShoppingPricingOpts{}
	if len(opts) > 0 && opts[len(opts)-1] != nil {
		opt = opts[len(opts)-1]
	}

	// Set defaults.
	internal.SetDefaultPages(&opt.Pages)
	internal.SetDefaultDomain(&opt.Domain)
	internal.SetDefaultStartPage(&opt.StartPage)
	internal.SetDefaultUserAgent(&opt.UserAgent)

	// Check validity of parameters.
	err := opt.checkParameterValidity()
	if err != nil {
		return nil, err
	}

	// Prepare payload with common parameters.
	payload := map[string]interface{}{
		"source":           oxylabs.GoogleShoppingPricing,
		"domain":           opt.Domain,
		"query":            query,
		"start_page":       opt.StartPage,
		"pages":            opt.Pages,
		"locale":           opt.Locale,
		"results_language": opt.ResultsLanguage,
		"geo_location":     opt.GeoLocation,
		"user_agent_type":  opt.UserAgent,
		"render":           opt.Render,
		"callback_url":     opt.CallbackURL,
		"parse":            opt.Parse,
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

	// Get job ID.
	jobID, err := c.C.GetJobID(jsonPayload)
	if err != nil {
		return nil, err
	}

	// Poll job status.
	go c.C.PollJobStatus(
		ctx,
		jobID,
		opt.Parse,
		customParserFlag,
		opt.PollInterval,
		internalRespChan,
		errChan,
	)

	// Handle error.
	err = <-errChan
	if err != nil {
		return nil, err
	}

	// Retrieve internal resp and forward it to the
	// external resp channel.
	internalResp := <-internalRespChan
	go func() {
		respChan <- &EcommerceResp{
			Resp: *internalResp,
		}
	}()

	return respChan, nil
}
