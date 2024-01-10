package ecommerce

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mslmio/oxylabs-sdk-go/oxylabs"
)

// ScrapeWayfairSearch scrapes wayfair with async polling runtime via Oxylabs E-Commerce API
// and wayfair_search as source.
func (c *EcommerceClientAsync) ScrapeWayfairSearch(
	query string,
	opts ...*WayfairSearchOpts,
) (chan *Resp, error) {
	ctx, cancel := context.WithTimeout(context.Background(), oxylabs.DefaultTimeout)
	defer cancel()

	return c.ScrapeWayfairSearchCtx(ctx, query, opts...)
}

// ScrapeWayfairSearchCtx scrapes wayfair with async polling runtime via Oxylabs E-Commerce API
// and wayfair_search as source.
// The provided context allows customization of the HTTP request, including setting timeouts.
func (c *EcommerceClientAsync) ScrapeWayfairSearchCtx(
	ctx context.Context,
	query string,
	opts ...*WayfairSearchOpts,
) (chan *Resp, error) {
	responseChan := make(chan *Resp)
	errChan := make(chan error)

	// Prepare options.
	opt := &WayfairSearchOpts{}
	if len(opts) > 0 && opts[len(opts)-1] != nil {
		opt = opts[len(opts)-1]
	}

	// Set defaults.
	SetDefaultLimit(&opt.Limit)
	SetDefaultPages(&opt.Pages)
	SetDefaultStartPage(&opt.StartPage)
	SetDefaultUserAgent(&opt.UserAgent)

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
	go c.PollJobStatus(ctx, jobID, false, responseChan, errChan)

	err = <-errChan
	if err != nil {
		return nil, err
	}

	return responseChan, nil
}

// ScrapeWayfairUrl scrapes wayfair with async polling runtime via Oxylabs E-Commerce API
// and wayfair as source.
func (c *EcommerceClientAsync) ScrapeWayfairUrl(
	url string,
	opts ...*WayfairUrlOpts,
) (chan *Resp, error) {
	ctx, cancel := context.WithTimeout(context.Background(), oxylabs.DefaultTimeout)
	defer cancel()

	return c.ScrapeWayfairUrlCtx(ctx, url, opts...)
}

// ScrapeWayfairUrlCtx scrapes wayfair with async polling runtime via Oxylabs E-Commerce API
// and wayfair as source.
// The provided context allows customization of the HTTP request, including setting timeouts.
func (c *EcommerceClientAsync) ScrapeWayfairUrlCtx(
	ctx context.Context,
	url string,
	opts ...*WayfairUrlOpts,
) (chan *Resp, error) {
	responseChan := make(chan *Resp)
	errChan := make(chan error)

	// Check validity of url.
	err := oxylabs.ValidateURL(url, "wayfair")
	if err != nil {
		return nil, err
	}

	// Prepare options.
	opt := &WayfairUrlOpts{}
	if len(opts) > 0 && opts[len(opts)-1] != nil {
		opt = opts[len(opts)-1]
	}

	// Set defaults.
	SetDefaultUserAgent(&opt.UserAgent)

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
	go c.PollJobStatus(ctx, jobID, false, responseChan, errChan)

	err = <-errChan
	if err != nil {
		return nil, err
	}

	return responseChan, nil
}
