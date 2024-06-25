package ecommerce

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/oxylabs/oxylabs-sdk-go/internal"
	"github.com/oxylabs/oxylabs-sdk-go/oxylabs"
)

// ScrapeWayfairSearch scrapes wayfair with async polling runtime via Oxylabs E-Commerce API
// and wayfair_search as source.
func (c *EcommerceClientAsync) ScrapeWayfairSearch(
	query string,
	opts ...*WayfairSearchOpts,
) (chan *Resp, error) {
	ctx, cancel := context.WithTimeout(context.Background(), internal.DefaultTimeout)
	defer cancel()

	return c.ScrapeWayfairSearchCtx(ctx, query, opts...)
}

// ScrapeWayfairSearchCtx scrapes wayfair with async polling runtime via Oxylabs E-Commerce API
// and wayfair_search as source.
// The provided context allows customization of the HTTP req, including setting timeouts.
func (c *EcommerceClientAsync) ScrapeWayfairSearchCtx(
	ctx context.Context,
	query string,
	opts ...*WayfairSearchOpts,
) (chan *Resp, error) {
	errChan := make(chan error)
	httpRespChan := make(chan *http.Response)
	respChan := make(chan *Resp)

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
		payload["parse"] = true
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
	resp, err := GetResp(httpResp, customParserFlag, customParserFlag)
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

// ScrapeWayfairUrl scrapes wayfair with async polling runtime via Oxylabs E-Commerce API
// and wayfair as source.
func (c *EcommerceClientAsync) ScrapeWayfairUrl(
	url string,
	opts ...*WayfairUrlOpts,
) (chan *Resp, error) {
	ctx, cancel := context.WithTimeout(context.Background(), internal.DefaultTimeout)
	defer cancel()

	return c.ScrapeWayfairUrlCtx(ctx, url, opts...)
}

// ScrapeWayfairUrlCtx scrapes wayfair with async polling runtime via Oxylabs E-Commerce API
// and wayfair as source.
// The provided context allows customization of the HTTP req, including setting timeouts.
func (c *EcommerceClientAsync) ScrapeWayfairUrlCtx(
	ctx context.Context,
	url string,
	opts ...*WayfairUrlOpts,
) (chan *Resp, error) {
	errChan := make(chan error)
	httpRespChan := make(chan *http.Response)
	respChan := make(chan *Resp)

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
		payload["parse"] = true
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
	resp, err := GetResp(httpResp, customParserFlag, customParserFlag)
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
