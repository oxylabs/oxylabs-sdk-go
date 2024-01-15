package serp

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mslmio/oxylabs-sdk-go/internal"
	"github.com/mslmio/oxylabs-sdk-go/oxylabs"
)

// ScrapeGoogleSearch scrapes google with async polling runtime via Oxylabs SERP API
// and google_search as source.
func (c *SerpClientAsync) ScrapeGoogleSearch(
	query string,
	opts ...*GoogleSearchOpts,
) (chan *internal.Resp, error) {
	ctx, cancel := context.WithTimeout(context.Background(), internal.DefaultTimeout)
	defer cancel()

	return c.ScrapeGoogleSearchCtx(ctx, query, opts...)
}

// ScrapeGoogleSearchCtx scrapes google with async polling runtime via Oxylabs SERP API
// and google_search as source.
// The provided context allows customization of the HTTP request, including setting timeouts.
func (c *SerpClientAsync) ScrapeGoogleSearchCtx(
	ctx context.Context,
	query string,
	opts ...*GoogleSearchOpts,
) (chan *internal.Resp, error) {
	respChan := make(chan *internal.Resp)
	errChan := make(chan error)

	// Prepare options.
	opt := &GoogleSearchOpts{}
	if len(opts) > 0 && opts[len(opts)-1] != nil {
		opt = opts[len(opts)-1]
	}

	// Initialize the context map and apply each provided context modifier function.
	context := make(oxylabs.ContextOption)
	for _, modifier := range opt.Context {
		modifier(context)
	}

	// Check if limit_per_page context parameter is used together with limit, start_page or pages parameters.
	if (opt.Limit != 0 || opt.StartPage != 0 || opt.Pages != 0) && context["limit_per_page"] != nil {
		return nil, fmt.Errorf(
			"limit, start_page and pages parameters cannot be used together with limit_per_page context parameter",
		)
	}

	// Set defaults.
	internal.SetDefaultDomain(&opt.Domain)
	internal.SetDefaultStartPage(&opt.StartPage)
	internal.SetDefaultLimit(&opt.Limit, internal.DefaultLimit_SERP)
	internal.SetDefaultPages(&opt.Pages)
	internal.SetDefaultUserAgent(&opt.UserAgent)

	// Check validity of parameters.
	err := opt.checkParameterValidity(context)
	if err != nil {
		return nil, err
	}

	// Prepare payload with common parameters.
	payload := map[string]interface{}{
		"source":          oxylabs.GoogleSearch,
		"domain":          opt.Domain,
		"query":           query,
		"locale":          opt.Locale,
		"geo_location":    opt.GeoLocation,
		"user_agent_type": opt.UserAgent,
		"parse":           opt.Parse,
		"render":          opt.Render,
		"callback_url":    opt.CallbackUrl,
		"context": []map[string]interface{}{
			{
				"key":   "results_language",
				"value": context["results_language"],
			},
			{
				"key":   "filter",
				"value": context["filter"],
			},
			{
				"key":   "nfpr",
				"value": context["nfpr"],
			},
			{
				"key":   "safe_search",
				"value": context["safe_search"],
			},
			{
				"key":   "fpstate",
				"value": context["fpstate"],
			},
			{
				"key":   "tbm",
				"value": context["tbm"],
			},
			{
				"key":   "tbs",
				"value": context["tbs"],
			},
		},
	}

	// If user sends limit_per_page context parameter, use it instead of limit, start_page, and pages parameters.
	if context["limit_per_page"] != nil {
		payload["limit_per_page"] = context["limit_per_page"]
	} else {
		payload["start_page"] = opt.StartPage
		payload["pages"] = opt.Pages
		payload["limit"] = opt.Limit
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
		respChan,
		errChan,
	)

	err = <-errChan
	if err != nil {
		return nil, err
	}

	return respChan, nil
}

// ScrapeGoogleUrl scrapes google with async polling runtime via Oxylabs SERP API
// and google as source.
func (c *SerpClientAsync) ScrapeGoogleUrl(
	url string,
	opts ...*GoogleUrlOpts,
) (chan *internal.Resp, error) {
	ctx, cancel := context.WithTimeout(context.Background(), internal.DefaultTimeout)
	defer cancel()

	return c.ScrapeGoogleUrlCtx(ctx, url, opts...)
}

// ScrapeGoogleUrlCtx scrapes google with async polling runtime via Oxylabs SERP API
// and google as source.
// The provided context allows customization of the HTTP request, including setting timeouts.
func (c *SerpClientAsync) ScrapeGoogleUrlCtx(
	ctx context.Context,
	url string,
	opts ...*GoogleUrlOpts,
) (chan *internal.Resp, error) {
	respChan := make(chan *internal.Resp)
	errChan := make(chan error)

	// Check validity of URL.
	err := internal.ValidateUrl(url, "google")
	if err != nil {
		return nil, err
	}

	// Prepare options.
	opt := &GoogleUrlOpts{}
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
		"source":          oxylabs.GoogleUrl,
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
		respChan,
		errChan,
	)

	err = <-errChan
	if err != nil {
		return nil, err
	}

	return respChan, nil
}

// ScrapeGoogleAds scrapes google with async polling runtime via Oxylabs SERP API
// and google_ads as source.
func (c *SerpClientAsync) ScrapeGoogleAds(
	query string,
	opts ...*GoogleAdsOpts,
) (chan *internal.Resp, error) {
	ctx, cancel := context.WithTimeout(context.Background(), internal.DefaultTimeout)
	defer cancel()

	return c.ScrapeGoogleAdsCtx(ctx, query, opts...)
}

// ScrapeGoogleAdsCtx scrapes google with async polling runtime via Oxylabs SERP API
// and google_ads as source.
// The provided context allows customization of the HTTP request, including setting timeouts.
func (c *SerpClientAsync) ScrapeGoogleAdsCtx(
	ctx context.Context,
	query string,
	opts ...*GoogleAdsOpts,
) (chan *internal.Resp, error) {
	respChan := make(chan *internal.Resp)
	errChan := make(chan error)

	// Prepare options.
	opt := &GoogleAdsOpts{}
	if len(opts) > 0 && opts[len(opts)-1] != nil {
		opt = opts[len(opts)-1]
	}

	// Initialize the context map apply each provided context modifier function.
	context := make(oxylabs.ContextOption)
	for _, modifier := range opt.Context {
		modifier(context)
	}

	// Set defaults.
	internal.SetDefaultDomain(&opt.Domain)
	internal.SetDefaultStartPage(&opt.StartPage)
	internal.SetDefaultPages(&opt.Pages)
	internal.SetDefaultUserAgent(&opt.UserAgent)

	// Check validity of parameters.
	err := opt.checkParameterValidity(context)
	if err != nil {
		return nil, err
	}

	payload := map[string]interface{}{
		"source":          oxylabs.GoogleAds,
		"domain":          opt.Domain,
		"query":           query,
		"start_page":      opt.StartPage,
		"pages":           opt.Pages,
		"locale":          opt.Locale,
		"geo_location":    opt.GeoLocation,
		"user_agent_type": opt.UserAgent,
		"parse":           opt.Parse,
		"render":          opt.Render,
		"callback_url":    opt.CallbackUrl,
		"context": []map[string]interface{}{
			{
				"key":   "results_language",
				"value": context["results_language"],
			},
			{
				"key":   "nfpr",
				"value": context["nfpr"],
			},
			{
				"key":   "tbm",
				"value": context["tbm"],
			},
			{
				"key":   "tbs",
				"value": context["tbs"],
			},
		},
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
	jobID, err := c.C.GetJobID(jsonPayload)
	if err != nil {
		return nil, err
	}

	// Poll job status.
	go c.C.PollJobStatus(ctx,
		jobID,
		opt.Parse,
		customParserFlag,
		opt.PollInterval,
		respChan,
		errChan,
	)

	err = <-errChan
	if err != nil {
		return nil, err
	}

	return respChan, nil
}

// ScrapeGoogleSuggestions scrapes google with async polling runtime via Oxylabs SERP API
// and google_suggestions as source.
func (c *SerpClientAsync) ScrapeGoogleSuggestions(
	query string,
	opts ...*GoogleSuggestionsOpts,
) (chan *internal.Resp, error) {
	ctx, cancel := context.WithTimeout(context.Background(), internal.DefaultTimeout)
	defer cancel()

	return c.ScrapeGoogleSuggestionsCtx(ctx, query, opts...)
}

// ScrapeGoogleSuggestionsCtx scrapes google with async polling runtime via Oxylabs SERP API
// and google_suggestions as source.
// The provided context allows customization of the HTTP request, including setting timeouts.
func (c *SerpClientAsync) ScrapeGoogleSuggestionsCtx(
	ctx context.Context,
	query string,
	opts ...*GoogleSuggestionsOpts,
) (chan *internal.Resp, error) {
	respChan := make(chan *internal.Resp)
	errChan := make(chan error)

	// Prepare options.
	opt := &GoogleSuggestionsOpts{}
	if len(opts) > 0 && opts[len(opts)-1] != nil {
		opt = opts[len(opts)-1]
	}

	// Set defaults.
	internal.SetDefaultUserAgent(&opt.UserAgent)

	// Check validity of parameters.
	err := opt.checkParameterValidity()
	if err != nil {
		return nil, err
	}

	// Prepare payload.
	payload := map[string]interface{}{
		"source":          oxylabs.GoogleSuggestions,
		"query":           query,
		"locale":          opt.Locale,
		"geo_location":    opt.GeoLocation,
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

	// Get job ID.
	jobID, err := c.C.GetJobID(jsonPayload)
	if err != nil {
		return nil, err
	}

	// Poll job status.
	go c.C.PollJobStatus(
		ctx,
		jobID,
		customParserFlag,
		customParserFlag,
		opt.PollInterval,
		respChan,
		errChan,
	)

	err = <-errChan
	if err != nil {
		return nil, err
	}

	return respChan, nil
}

// ScrapeGoogleHotels scrapes google with async polling runtime via Oxylabs SERP API
// and google_hotels as source.
func (c *SerpClientAsync) ScrapeGoogleHotels(
	query string,
	opts ...*GoogleHotelsOpts,
) (chan *internal.Resp, error) {
	ctx, cancel := context.WithTimeout(context.Background(), internal.DefaultTimeout)
	defer cancel()

	return c.ScrapeGoogleHotelsCtx(ctx, query, opts...)
}

// ScrapeGoogleHotelsCtx scrapes google with async polling runtime via Oxylabs SERP API
// and google_hotels as source.
// The provided context allows customization of the HTTP request, including setting timeouts.
func (c *SerpClientAsync) ScrapeGoogleHotelsCtx(
	ctx context.Context,
	query string,
	opts ...*GoogleHotelsOpts,
) (chan *internal.Resp, error) {
	respChan := make(chan *internal.Resp)
	errChan := make(chan error)

	// Prepare options.
	opt := &GoogleHotelsOpts{}
	if len(opts) > 0 && opts[len(opts)-1] != nil {
		opt = opts[len(opts)-1]
	}

	// Initialize the context map apply each provided context modifier function.
	context := make(oxylabs.ContextOption)
	for _, modifier := range opt.Context {
		modifier(context)
	}

	// Set defaults.
	internal.SetDefaultDomain(&opt.Domain)
	internal.SetDefaultStartPage(&opt.StartPage)
	internal.SetDefaultLimit(&opt.Limit, internal.DefaultLimit_SERP)
	internal.SetDefaultPages(&opt.Pages)
	internal.SetDefaultUserAgent(&opt.UserAgent)

	// Check validity of parameters.
	err := opt.checkParameterValidity(context)
	if err != nil {
		return nil, err
	}

	// Prepare payload.
	payload := map[string]interface{}{
		"source":          oxylabs.GoogleHotels,
		"domain":          opt.Domain,
		"query":           query,
		"start_page":      opt.StartPage,
		"pages":           opt.Pages,
		"limit":           opt.Limit,
		"locale":          opt.Locale,
		"geo_location":    opt.GeoLocation,
		"user_agent_type": opt.UserAgent,
		"render":          opt.Render,
		"callback_url":    opt.CallbackUrl,
		"context": []map[string]interface{}{
			{
				"key":   "results_language",
				"value": context["results_language"],
			},
			{
				"key":   "nfpr",
				"value": context["nfpr"],
			},
			{
				"key":   "hotel_occupancy",
				"value": context["hotel_occupancy"],
			},
			{
				"key":   "hotel_dates",
				"value": context["hotel_dates"],
			},
		},
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

	// Get job ID.
	jobID, err := c.C.GetJobID(jsonPayload)
	if err != nil {
		return nil, err
	}

	// Poll job status.
	go c.C.PollJobStatus(
		ctx,
		jobID,
		customParserFlag,
		customParserFlag,
		opt.PollInterval,
		respChan,
		errChan,
	)

	err = <-errChan
	if err != nil {
		return nil, err
	}

	return respChan, nil
}

// ScrapeGoogleTravelHotels scrapes google with async polling runtime via Oxylabs SERP API
// and google_travel_hotels as source.
func (c *SerpClientAsync) ScrapeGoogleTravelHotels(
	query string,
	opts ...*GoogleTravelHotelsOpts,
) (chan *internal.Resp, error) {
	ctx, cancel := context.WithTimeout(context.Background(), internal.DefaultTimeout)
	defer cancel()

	return c.ScrapeGoogleTravelHotelsCtx(ctx, query, opts...)
}

// ScrapeGoogleTravelHotelsCtx scrapes google with async polling runtime via Oxylabs SERP API
// and google_travel_hotels as source.
// The provided context allows customization of the HTTP request, including setting timeouts.
func (c *SerpClientAsync) ScrapeGoogleTravelHotelsCtx(
	ctx context.Context,
	query string,
	opts ...*GoogleTravelHotelsOpts,
) (chan *internal.Resp, error) {
	respChan := make(chan *internal.Resp)
	errChan := make(chan error)

	// Prepare options.
	opt := &GoogleTravelHotelsOpts{}
	if len(opts) > 0 && opts[len(opts)-1] != nil {
		opt = opts[len(opts)-1]
	}

	// Initialize the context map apply each provided context modifier function.
	context := make(oxylabs.ContextOption)
	for _, modifier := range opt.Context {
		modifier(context)
	}

	// Set defaults.
	internal.SetDefaultDomain(&opt.Domain)
	internal.SetDefaultStartPage(&opt.StartPage)

	// Check validity of parameters.
	err := opt.checkParameterValidity(context)
	if err != nil {
		return nil, err
	}

	// Prepare payload.
	payload := map[string]interface{}{
		"source":          oxylabs.GoogleTravelHotels,
		"domain":          opt.Domain,
		"query":           query,
		"start_page":      opt.StartPage,
		"locale":          opt.Locale,
		"geo_location":    opt.GeoLocation,
		"user_agent_type": opt.UserAgent,
		"render":          opt.Render,
		"callback_url":    opt.CallbackUrl,
		"context": []map[string]interface{}{
			{
				"key":   "hotel_occupancy",
				"value": context["hotel_occupancy"],
			},
			{
				"key":   "hotel_classes",
				"value": context["hotel_classes"],
			},
			{
				"key":   "hotel_dates",
				"value": context["hotel_dates"],
			},
		},
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

	// Get job ID.
	jobID, err := c.C.GetJobID(jsonPayload)
	if err != nil {
		return nil, err
	}

	// Poll job status.
	go c.C.PollJobStatus(
		ctx,
		jobID,
		customParserFlag,
		customParserFlag,
		opt.PollInterval,
		respChan,
		errChan,
	)

	err = <-errChan
	if err != nil {
		return nil, err
	}

	return respChan, nil
}

// ScrapeGoogleImages scrapes google with async polling runtime via Oxylabs SERP API
// and google_images as source.
func (c *SerpClientAsync) ScrapeGoogleImages(
	url string,
	opts ...*GoogleImagesOpts,
) (chan *internal.Resp, error) {
	ctx, cancel := context.WithTimeout(context.Background(), internal.DefaultTimeout)
	defer cancel()

	return c.ScrapeGoogleImagesCtx(ctx, url, opts...)
}

// ScrapeGoogleImagesCtx scrapes google with async polling runtime via Oxylabs SERP API
// and google_images as source.
// The provided context allows customization of the HTTP request, including setting timeouts.
func (c *SerpClientAsync) ScrapeGoogleImagesCtx(
	ctx context.Context,
	url string,
	opts ...*GoogleImagesOpts,
) (chan *internal.Resp, error) {
	respChan := make(chan *internal.Resp)
	errChan := make(chan error)

	// Check validity of URL.
	err := internal.ValidateUrl(url, "google")
	if err != nil {
		return nil, err
	}

	// Prepare options.
	opt := &GoogleImagesOpts{}
	if len(opts) > 0 && opts[len(opts)-1] != nil {
		opt = opts[len(opts)-1]
	}

	// Initialize the context map apply each provided context modifier function.
	context := make(oxylabs.ContextOption)
	for _, modifier := range opt.Context {
		modifier(context)
	}

	// Set defaults.
	internal.SetDefaultDomain(&opt.Domain)
	internal.SetDefaultStartPage(&opt.StartPage)
	internal.SetDefaultPages(&opt.Pages)

	// Check validity of parameters.
	err = opt.checkParameterValidity()
	if err != nil {
		return nil, err
	}

	// Prepare payload.
	payload := map[string]interface{}{
		"source":          oxylabs.GoogleImages,
		"domain":          opt.Domain,
		"query":           url,
		"start_page":      opt.StartPage,
		"pages":           opt.Pages,
		"locale":          opt.Locale,
		"geo_location":    opt.GeoLocation,
		"user_agent_type": opt.UserAgent,
		"render":          opt.Render,
		"callback_url":    opt.CallbackUrl,
		"context": []map[string]interface{}{
			{
				"key":   "nfpr",
				"value": context["nfpr"],
			},
			{
				"key":   "results_language",
				"value": context["results_language"],
			},
		},
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

	// Get job ID.
	jobID, err := c.C.GetJobID(jsonPayload)
	if err != nil {
		return nil, err
	}

	// Poll job status.
	go c.C.PollJobStatus(
		ctx,
		jobID,
		customParserFlag,
		customParserFlag,
		opt.PollInterval,
		respChan,
		errChan,
	)

	err = <-errChan
	if err != nil {
		return nil, err
	}

	return respChan, nil
}

// ScrapeGoogleTrendsExplore scrapes google with async polling runtime via Oxylabs SERP API
// and google_trends_explore as source.
func (c *SerpClientAsync) ScrapeGoogleTrendsExplore(
	query string,
	opts ...*GoogleTrendsExploreOpts,
) (chan *internal.Resp, error) {
	ctx, cancel := context.WithTimeout(context.Background(), internal.DefaultTimeout)
	defer cancel()

	return c.ScrapeGoogleTrendsExploreCtx(ctx, query, opts...)
}

// ScrapeGoogleTrendsExploreCtx scrapes google with async polling runtime via Oxylabs SERP API
// and google_trends_explore as source.
// The provided context allows customization of the HTTP request, including setting timeouts.
func (c *SerpClientAsync) ScrapeGoogleTrendsExploreCtx(
	ctx context.Context,
	query string,
	opts ...*GoogleTrendsExploreOpts,
) (chan *internal.Resp, error) {
	respChan := make(chan *internal.Resp)
	errChan := make(chan error)

	// Prepare options.
	opt := &GoogleTrendsExploreOpts{}
	if len(opts) > 0 && opts[len(opts)-1] != nil {
		opt = opts[len(opts)-1]
	}

	// Initialize the context map apply each provided context modifier function.
	context := make(oxylabs.ContextOption)
	for _, modifier := range opt.Context {
		modifier(context)
	}

	// Set defaults.
	internal.SetDefaultUserAgent(&opt.UserAgent)

	// Check validity of parameters.
	err := opt.checkParameterValidity(context)
	if err != nil {
		return nil, err
	}

	// Prepare payload.
	payload := map[string]interface{}{
		"source":       oxylabs.GoogleTrendsExplore,
		"query":        query,
		"geo_location": opt.GeoLocation,
		"context": []map[string]interface{}{
			{
				"key":   "search_type",
				"value": context["search_type"],
			},
			{
				"key":   "date_from",
				"value": context["date_from"],
			},
			{
				"key":   "date_to",
				"value": context["date_to"],
			},
			{
				"key":   "category_id",
				"value": context["category_id"],
			},
		},
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

	// Get job ID.
	jobID, err := c.C.GetJobID(jsonPayload)
	if err != nil {
		return nil, err
	}

	// Poll job status.
	go c.C.PollJobStatus(
		ctx,
		jobID,
		customParserFlag,
		customParserFlag,
		opt.PollInterval,
		respChan,
		errChan,
	)

	err = <-errChan
	if err != nil {
		return nil, err
	}

	return respChan, nil
}
