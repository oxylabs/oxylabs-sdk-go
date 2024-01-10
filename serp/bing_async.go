package serp

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mslmio/oxylabs-sdk-go/oxylabs"
)

// ScrapeBingSearch scrapes bing with async polling runtime via Oxylabs SERP API
// and bing_search as source.
func (c *SerpClientAsync) ScrapeBingSearch(
	query string,
	opts ...*BingSearchOpts,
) (chan *Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), oxylabs.DefaultTimeout)
	defer cancel()

	return c.ScrapeBingSearchCtx(ctx, query, opts...)
}

// ScrapeBingSearchCtx scrapes bing with async polling runtime via Oxylabs SERP API
// and bing_search as source.
// The provided context allows customization of the HTTP request, including setting timeouts.
func (c *SerpClientAsync) ScrapeBingSearchCtx(
	ctx context.Context,
	query string,
	opts ...*BingSearchOpts,
) (chan *Response, error) {
	responseChan := make(chan *Response)
	errChan := make(chan error)

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

	// Get job ID.
	jobID, err := c.GetJobID(jsonPayload)
	if err != nil {
		return nil, err
	}

	// Poll job status.
	go c.PollJobStatus(ctx, jobID, opt.Parse, customParserFlag, responseChan, errChan)

	err = <-errChan
	if err != nil {
		return nil, err
	}

	return responseChan, nil
}

// ScrapeBingUrl scrapes bing with async polling runtime via Oxylabs SERP API
// and bing as source.
func (c *SerpClientAsync) ScrapeBingUrl(
	url string,
	opts ...*BingUrlOpts,
) (chan *Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), oxylabs.DefaultTimeout)
	defer cancel()

	return c.ScrapeBingUrlCtx(ctx, url, opts...)
}

// ScrapeBingUrlCtx scrapes bing with async polling runtime via Oxylabs SERP API
// and bing as source.
// The provided context allows customization of the HTTP request, including setting timeouts.
func (c *SerpClientAsync) ScrapeBingUrlCtx(
	ctx context.Context,
	url string,
	opts ...*BingUrlOpts,
) (chan *Response, error) {
	responseChan := make(chan *Response)
	errChan := make(chan error)

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

	// Get job ID.
	jobID, err := c.GetJobID(jsonPayload)
	if err != nil {
		return nil, err
	}

	// Poll job status.
	go c.PollJobStatus(ctx, jobID, opt.Parse, customParserFlag, responseChan, errChan)

	err = <-errChan
	if err != nil {
		return nil, err
	}

	return responseChan, nil
}
