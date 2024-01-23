package serp

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mslmio/oxylabs-sdk-go/internal"
	"github.com/mslmio/oxylabs-sdk-go/oxylabs"
)

// ScrapeBaiduSearch scrapes baidu with async polling runtime via Oxylabs SERP API
// and baidu_search as source.
func (c *SerpClientAsync) ScrapeBaiduSearch(
	query string,
	opts ...*BaiduSearchOpts,
) (chan *SerpResp, error) {
	ctx, cancel := context.WithTimeout(context.Background(), internal.DefaultTimeout)
	defer cancel()

	return c.ScrapeBaiduSearchCtx(ctx, query, opts...)
}

// ScrapeBaiduSearchCtx scrapes baidu with async polling runtime via Oxylabs SERP API
// and baidu_search as source.
// The provided context allows customization of the HTTP req, including setting timeouts.
func (c *SerpClientAsync) ScrapeBaiduSearchCtx(
	ctx context.Context,
	query string,
	opts ...*BaiduSearchOpts,
) (chan *SerpResp, error) {
	httpRespChan := make(chan *http.Response)
	serpRespChan := make(chan *SerpResp)
	errChan := make(chan error)

	// Prepare options.
	opt := &BaiduSearchOpts{}
	if len(opts) > 0 && opts[len(opts)-1] != nil {
		opt = opts[len(opts)-1]
	}

	// Set defaults.
	internal.SetDefaultDomain(&opt.Domain)
	internal.SetDefaultStartPage(&opt.StartPage)
	internal.SetDefaultLimit(&opt.Limit, internal.DefaultLimit_SERP)
	internal.SetDefaultPages(&opt.Pages)
	internal.SetDefaultUserAgent(&opt.UserAgent)

	// Check validity of parameters.
	err := opt.checkParameterValidity()
	if err != nil {
		return nil, err
	}

	// Prepare payload.
	payload := map[string]interface{}{
		"source":          oxylabs.BaiduSearch,
		"domain":          opt.Domain,
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

	// Unmarshal the http Response and get the internal Response.
	httpResp := <-httpRespChan
	internalResp, err := internal.GetResp(httpResp, customParserFlag)
	if err != nil {
		return nil, err
	}

	// Retrieve internal resp and forward it to the
	// serp resp channel.
	go func() {
		serpRespChan <- &SerpResp{*internalResp}
	}()

	return serpRespChan, nil
}

// ScrapeBaiduUrl scrapes baidu with async polling runtime via Oxylabs SERP API
// and baidu as source.
func (c *SerpClientAsync) ScrapeBaiduUrl(
	query string,
	opts ...*BaiduUrlOpts,
) (chan *SerpResp, error) {
	ctx, cancel := context.WithTimeout(context.Background(), internal.DefaultTimeout)
	defer cancel()

	return c.ScrapeBaiduUrlCtx(ctx, query, opts...)
}

// ScrapeBaiduUrlCtx scrapes baidu with async polling runtime via Oxylabs SERP API
// and baidu as source.
// The provided context allows customization of the HTTP req, including setting timeouts.
func (c *SerpClientAsync) ScrapeBaiduUrlCtx(
	ctx context.Context,
	url string,
	opts ...*BaiduUrlOpts,
) (chan *SerpResp, error) {
	httpRespChan := make(chan *http.Response)
	serpRespChan := make(chan *SerpResp)
	errChan := make(chan error)

	// Check validity of URL.
	err := internal.ValidateUrl(url, "baidu")
	if err != nil {
		return nil, err
	}

	// Prepare options.
	opt := &BaiduUrlOpts{}
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
		"source":          oxylabs.BaiduUrl,
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

	// Unmarshal the http Response and get the internal Response.
	httpResp := <-httpRespChan
	internalResp, err := internal.GetResp(httpResp, customParserFlag)
	if err != nil {
		return nil, err
	}

	// Retrieve internal resp and forward it to the
	// serp resp channel.
	go func() {
		serpRespChan <- &SerpResp{*internalResp}
	}()

	return serpRespChan, nil
}
