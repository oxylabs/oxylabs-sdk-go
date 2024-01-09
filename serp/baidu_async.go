package serp

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mslmio/oxylabs-sdk-go/oxylabs"
)

// ScrapeBaiduSearch scrapes baidu with async polling runtime via Oxylabs SERP API
// and baidu_search as source.
func (c *SerpClientAsync) ScrapeBaiduSearch(
	query string,
	opts ...*BaiduSearchOpts,
) (chan *Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), oxylabs.DefaultTimeout)
	defer cancel()

	return c.ScrapeBaiduSearchCtx(ctx, query, opts...)
}

// ScrapeBaiduSearchCtx scrapes baidu with async polling runtime via Oxylabs SERP API
// and baidu_search as source.
// The provided context allows customization of the HTTP request, including setting timeouts.
func (c *SerpClientAsync) ScrapeBaiduSearchCtx(
	ctx context.Context,
	query string,
	opts ...*BaiduSearchOpts,
) (chan *Response, error) {
	responseChan := make(chan *Response)
	errChan := make(chan error)

	// Prepare options.
	opt := &BaiduSearchOpts{}
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
		"source":          "baidu_search",
		"domain":          opt.Domain,
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

// ScrapeBaiduUrl scrapes baidu with async polling runtime via Oxylabs SERP API
// and baidu as source.
func (c *SerpClientAsync) ScrapeBaiduUrl(
	query string,
	opts ...*BaiduUrlOpts,
) (chan *Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), oxylabs.DefaultTimeout)
	defer cancel()

	return c.ScrapeBaiduUrlCtx(ctx, query, opts...)
}

// ScrapeBaiduUrlCtx scrapes baidu with async polling runtime via Oxylabs SERP API
// and baidu as source.
// The provided context allows customization of the HTTP request, including setting timeouts.
func (c *SerpClientAsync) ScrapeBaiduUrlCtx(
	ctx context.Context,
	url string,
	opts ...*BaiduUrlOpts,
) (chan *Response, error) {
	responseChan := make(chan *Response)
	errChan := make(chan error)

	// Check validity of url.
	err := oxylabs.ValidateURL(url, "baidu")
	if err != nil {
		return nil, err
	}

	// Prepare options.
	opt := &BaiduUrlOpts{}
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
		"source":          "baidu",
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
