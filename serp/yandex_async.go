package serp

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mslmio/oxylabs-sdk-go/internal"
	"github.com/mslmio/oxylabs-sdk-go/oxylabs"
)

// ScrapeYandexSearch scrapes yandex with async polling runtime via Oxylabs SERP API
// and yandex_search as source.
func (c *SerpClientAsync) ScrapeYandexSearch(
	query string,
	opts ...*YandexSearchOpts,
) (chan *internal.Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), internal.DefaultTimeout)
	defer cancel()

	return c.ScrapeYandexSearchCtx(ctx, query, opts...)
}

// ScrapeYandexSearchCtx scrapes yandex with async polling runtime via Oxylabs SERP API
// and yandex_search as source.
// The provided context allows customization of the HTTP request, including setting timeouts.
func (c *SerpClientAsync) ScrapeYandexSearchCtx(
	ctx context.Context,
	query string,
	opts ...*YandexSearchOpts,
) (chan *internal.Response, error) {
	responseChan := make(chan *internal.Response)
	errChan := make(chan error)

	// Prepare options.
	opt := &YandexSearchOpts{}
	if len(opts) > 0 && opts[len(opts)-1] != nil {
		opt = opts[len(opts)-1]
	}

	// Set defaults.
	internal.SetDefaultDomain(&opt.Domain)
	internal.SetDefaultStartPage(&opt.StartPage)
	internal.SetDefaultLimit(&opt.Limit, internal.DefaultLimit_SERP)
	internal.SetDefaultPages(&opt.Pages)
	internal.SetDefaultUserAgent(&opt.UserAgent)

	// Check the validity of the parameters.
	err := opt.checkParameterValidity()
	if err != nil {
		return nil, err
	}

	// Prepare the payload.
	payload := map[string]interface{}{
		"source":          oxylabs.YandexSearch,
		"domain":          opt.Domain,
		"query":           query,
		"start_page":      opt.StartPage,
		"pages":           opt.Pages,
		"limit":           opt.Limit,
		"locale":          opt.Locale,
		"geo_location":    opt.GeoLocation,
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

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("error marshalling payload: %v", err)
	}

	// Get the job ID.
	jobID, err := c.InternalClient.GetJobID(jsonPayload)
	if err != nil {
		return nil, err
	}

	// Poll job status.
	go c.InternalClient.PollJobStatus(
		ctx,
		jobID,
		customParserFlag,
		customParserFlag,
		opt.PollInterval,
		responseChan,
		errChan,
	)

	err = <-errChan
	if err != nil {
		return nil, err
	}

	return responseChan, nil
}

// ScrapeYandexUrl scrapes yandex with async polling runtime via Oxylabs SERP API
// and yandex as source.
func (c *SerpClientAsync) ScrapeYandexUrl(
	url string,
	opts ...*YandexUrlOpts,
) (chan *internal.Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), internal.DefaultTimeout)
	defer cancel()

	return c.ScrapeYandexUrlCtx(ctx, url, opts...)
}

// ScrapeYandexUrlCtx scrapes yandex with async polling runtime via Oxylabs SERP API
// and yandex as source.
// The provided context allows customization of the HTTP request, including setting timeouts.
func (c *SerpClientAsync) ScrapeYandexUrlCtx(
	ctx context.Context,
	url string,
	opts ...*YandexUrlOpts,
) (chan *internal.Response, error) {
	responseChan := make(chan *internal.Response)
	errChan := make(chan error)

	// Check the validity of the Url.
	err := internal.ValidateUrl(url, "yandex")
	if err != nil {
		return nil, err
	}

	// Prepare options.
	opt := &YandexUrlOpts{}
	if len(opts) > 0 && opts[len(opts)-1] != nil {
		opt = opts[len(opts)-1]
	}

	// Set defaults.
	internal.SetDefaultUserAgent(&opt.UserAgent)

	// Check the validity of parameters.
	err = opt.checkParameterValidity()
	if err != nil {
		return nil, err
	}

	// Prepare the payload.
	payload := map[string]interface{}{
		"source":          oxylabs.YandexUrl,
		"url":             url,
		"user_agent_type": opt.UserAgent,
		"render":          opt.Render,
		"callback_url":    opt.CallbackUrl,
	}

	// Add custom parsing instructions to the payload if provided.
	customParserFlag := false
	if opt.ParseInstructions != nil {
		payload["parse"] = true
		payload["parsing_instructions"] = &opt.ParseInstructions
		customParserFlag = true
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("error marshalling payload: %v", err)
	}

	// Get the job ID.
	jobID, err := c.InternalClient.GetJobID(jsonPayload)
	if err != nil {
		return nil, err
	}

	// Poll job status.
	go c.InternalClient.PollJobStatus(
		ctx,
		jobID,
		customParserFlag,
		customParserFlag,
		opt.PollInterval,
		responseChan,
		errChan,
	)

	err = <-errChan
	if err != nil {
		return nil, err
	}

	return responseChan, nil
}
