package ecommerce

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mslmio/oxylabs-sdk-go/internal"
	"github.com/mslmio/oxylabs-sdk-go/oxylabs"
)

// ScrapeUniversalUrl scrapes all urls with async polling runtime via Oxylabs E-Commerce API
// and universal as source.
func (c *EcommerceClientAsync) ScrapeUniversalUrl(
	url string,
	opts ...*UniversalUrlOpts,
) (chan *EcommerceResp, error) {
	ctx, cancel := context.WithTimeout(context.Background(), internal.DefaultTimeout)
	defer cancel()

	return c.ScrapeUniversalUrlCtx(ctx, url, opts...)
}

// ScrapeUniversalUrlCtx scrapes all urls with async polling runtime via Oxylabs E-Commerce API
// and universal as source.
// The provided context allows customization of the HTTP req, including setting timeouts.
func (c *EcommerceClientAsync) ScrapeUniversalUrlCtx(
	ctx context.Context,
	url string,
	opts ...*UniversalUrlOpts,
) (chan *EcommerceResp, error) {
	errChan := make(chan error)
	internalRespChan := make(chan *internal.Resp)
	ecommerceRespChan := make(chan *EcommerceResp)

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
	// ecommerce resp channel.
	go func() {
		internalResp := <-internalRespChan
		ecommerceRespChan <- &EcommerceResp{*internalResp}
	}()

	return ecommerceRespChan, nil
}
