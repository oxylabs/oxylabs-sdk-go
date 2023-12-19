package serp

import (
	"encoding/json"
	"fmt"

	"github.com/mslmio/oxylabs-sdk-go/oxylabs"
)

// ScrapeGoogleSearch scrapes google with google_search as source with async polling runtime.
func (c *SerpClientAsync) ScrapeGoogleSearch(
	query string,
	opts ...*GoogleSearchOpts,
) (chan *Response, error) {
	responseChan := make(chan *Response)
	errChan := make(chan error)

	// Prepare options.
	opt := &GoogleSearchOpts{}
	if len(opts) > 0 && opts[len(opts)-1] != nil {
		opt = opts[len(opts)-1]
	}

	// Initialize the context map and apply each provided context modifier function.
	context := make(ContextOption)
	for _, modifier := range opt.Context {
		modifier(context)
	}

	// Check if limit_per_page context parameter is used together with limit, start_page or pages parameters.
	if (opt.Limit != 0 || opt.StartPage != 0 || opt.Pages != 0) && context["limit_per_page"] != nil {
		return nil, fmt.Errorf("limit, start_page and pages parameters cannot be used together with limit_per_page context parameter")
	}

	// Set defaults.
	SetDefaultDomain(&opt.Domain)
	SetDefaultStartPage(&opt.StartPage)
	SetDefaultLimit(&opt.Limit)
	SetDefaultPages(&opt.Pages)
	SetDefaultUserAgent(&opt.UserAgent)

	// Check validity of parameters.
	err := opt.checkParameterValidity(context)
	if err != nil {
		return nil, err
	}

	// Prepare payload with common parameters.
	payload := map[string]interface{}{
		"source":          "google_search",
		"domain":          opt.Domain,
		"query":           query,
		"geo_location":    &opt.GeoLocation,
		"user_agent_type": opt.UserAgent,
		"parse":           opt.Parse,
		"render":          opt.Render,
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
	go c.PollJobStatus(jobID, opt.Parse, responseChan, errChan)

	err = <-errChan
	if err != nil {
		return nil, err
	}

	return responseChan, nil
}

// ScrapeGoogleUrl scrapes google with google as source with async polling runtime.
func (c *SerpClientAsync) ScrapeGoogleUrl(
	url string,
	opts ...*GoogleUrlOpts,
) (chan *Response, error) {
	responseChan := make(chan *Response)
	errChan := make(chan error)

	// Check validity of url.
	err := oxylabs.ValidateURL(url, "google")
	if err != nil {
		return nil, err
	}

	// Prepare options.
	opt := &GoogleUrlOpts{}
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
		"source":          "google",
		"url":             url,
		"user_agent_type": opt.UserAgent,
		"render":          opt.Render,
		"callback_url":    opt.CallbackUrl,
		"geo_location":    &opt.GeoLocation,
		"parse":           opt.Parse,
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
	go c.PollJobStatus(jobID, opt.Parse, responseChan, errChan)

	err = <-errChan
	if err != nil {
		return nil, err
	}

	return responseChan, nil
}

// ScrapeGoogleAds scrapes google with google_ads as source with async polling runtime.
func (c *SerpClientAsync) ScrapeGoogleAds(
	query string,
	opts ...*GoogleAdsOpts,
) (chan *Response, error) {
	responseChan := make(chan *Response)
	errChan := make(chan error)

	// Prepare options.
	opt := &GoogleAdsOpts{}
	if len(opts) > 0 && opts[len(opts)-1] != nil {
		opt = opts[len(opts)-1]
	}

	// Initialize the context map apply each provided context modifier function.
	context := make(ContextOption)
	for _, modifier := range opt.Context {
		modifier(context)
	}

	// Set defaults.
	SetDefaultDomain(&opt.Domain)
	SetDefaultStartPage(&opt.StartPage)
	SetDefaultLimit(&opt.Limit)
	SetDefaultPages(&opt.Pages)
	SetDefaultUserAgent(&opt.UserAgent)

	// Check validity of parameters.
	err := opt.checkParameterValidity(context)
	if err != nil {
		return nil, err
	}

	payload := map[string]interface{}{
		"source":          "google_search",
		"domain":          opt.Domain,
		"query":           query,
		"geo_location":    &opt.GeoLocation,
		"user_agent_type": opt.UserAgent,
		"parse":           opt.Parse,
		"render":          opt.Render,
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
	go c.PollJobStatus(jobID, opt.Parse, responseChan, errChan)

	err = <-errChan
	if err != nil {
		return nil, err
	}

	return responseChan, nil
}

// ScrapeGoogleSuggestions scrapes google with google_suggestions as source with async polling runtime.
func (c *SerpClientAsync) ScrapeGoogleSuggestions(
	query string,
	opts ...*GoogleSuggestionsOpts,
) (chan *Response, error) {
	responseChan := make(chan *Response)
	errChan := make(chan error)

	// Prepare options.
	opt := &GoogleSuggestionsOpts{}
	if len(opts) > 0 && opts[len(opts)-1] != nil {
		opt = opts[len(opts)-1]
	}

	// Set defaults.
	SetDefaultUserAgent(&opt.UserAgent)

	// Check validity of parameters.
	err := opt.checkParameterValidity()
	if err != nil {
		return nil, err
	}

	// Prepare payload.
	payload := map[string]interface{}{
		"source":          "google_suggestions",
		"query":           query,
		"geo_location":    &opt.GeoLocation,
		"user_agent_type": opt.UserAgent,
		"render":          opt.Render,
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
	go c.PollJobStatus(jobID, false, responseChan, errChan)

	err = <-errChan
	if err != nil {
		return nil, err
	}

	return responseChan, nil
}

// ScrapeGoogleTravelHotels scrapes google with google_hotels as source with async polling runtime.
func (c *SerpClientAsync) ScrapeGoogleHotels(
	query string,
	opts ...*GoogleHotelsOpts,
) (chan *Response, error) {
	responseChan := make(chan *Response)
	errChan := make(chan error)

	// Prepare options.
	opt := &GoogleHotelsOpts{}
	if len(opts) > 0 && opts[len(opts)-1] != nil {
		opt = opts[len(opts)-1]
	}

	// Initialize the context map apply each provided context modifier function.
	context := make(ContextOption)
	for _, modifier := range opt.Context {
		modifier(context)
	}

	// Set defaults.
	SetDefaultDomain(&opt.Domain)
	SetDefaultStartPage(&opt.StartPage)
	SetDefaultLimit(&opt.Limit)
	SetDefaultPages(&opt.Pages)
	SetDefaultUserAgent(&opt.UserAgent)

	// Check validity of parameters.
	err := opt.checkParameterValidity(context)
	if err != nil {
		return nil, err
	}

	// Prepare payload.
	payload := map[string]interface{}{
		"source":           "google_hotels",
		"domain":           opt.Domain,
		"query":            query,
		"start_page":       opt.StartPage,
		"pages":            opt.Pages,
		"limit":            opt.Limit,
		"locale":           opt.Locale,
		"results_language": opt.ResultsLanguage,
		"geo_location":     &opt.GeoLocation,
		"user_agent_type":  opt.UserAgent,
		"render":           opt.Render,
		"callback_url":     opt.CallbackURL,
		"context": []map[string]interface{}{
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
	go c.PollJobStatus(jobID, false, responseChan, errChan)

	err = <-errChan
	if err != nil {
		return nil, err
	}

	return responseChan, nil
}

// ScrapeGoogleTravelHotels scrapes google with google_travel_hotels as source with async polling runtime.
func (c *SerpClientAsync) ScrapeGoogleTravelHotels(
	query string,
	opts ...*GoogleTravelHotelsOpts,
) (chan *Response, error) {
	responseChan := make(chan *Response)
	errChan := make(chan error)

	// Prepare options.
	opt := &GoogleTravelHotelsOpts{}
	if len(opts) > 0 && opts[len(opts)-1] != nil {
		opt = opts[len(opts)-1]
	}

	// Initialize the context map apply each provided context modifier function.
	context := make(ContextOption)
	for _, modifier := range opt.Context {
		modifier(context)
	}

	// Set defaults.
	SetDefaultDomain(&opt.Domain)
	SetDefaultStartPage(&opt.StartPage)
	SetDefaultLimit(&opt.Limit)
	SetDefaultPages(&opt.Pages)

	// Check validity of parameters.
	err := opt.checkParameterValidity(context)
	if err != nil {
		return nil, err
	}

	// Prepare payload.
	payload := map[string]interface{}{
		"source":          "google_travel_hotels",
		"domain":          opt.Domain,
		"query":           query,
		"start_page":      opt.StartPage,
		"pages":           opt.Pages,
		"limit":           opt.Limit,
		"locale":          opt.Locale,
		"geo_location":    &opt.GeoLocation,
		"user_agent_type": opt.UserAgent,
		"render":          opt.Render,
		"callback_url":    opt.CallbackURL,
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
	go c.PollJobStatus(jobID, false, responseChan, errChan)

	err = <-errChan
	if err != nil {
		return nil, err
	}

	return responseChan, nil
}

// ScrapeGoogleImages scrapes google with google_images as source with async polling runtime.
func (c *SerpClientAsync) ScrapeGoogleImages(
	url string,
	opts ...*GoogleImagesOpts,
) (chan *Response, error) {
	responseChan := make(chan *Response)
	errChan := make(chan error)

	// Check validity of url.
	err := oxylabs.ValidateURL(url, "google")
	if err != nil {
		return nil, err
	}

	// Prepare options.
	opt := &GoogleImagesOpts{}
	if len(opts) > 0 && opts[len(opts)-1] != nil {
		opt = opts[len(opts)-1]
	}

	// Initialize the context map apply each provided context modifier function.
	context := make(ContextOption)
	for _, modifier := range opt.Context {
		modifier(context)
	}

	// Set defaults.
	SetDefaultDomain(&opt.Domain)
	SetDefaultStartPage(&opt.StartPage)
	SetDefaultPages(&opt.Pages)

	// Check validity of parameters.
	err = opt.checkParameterValidity(context)
	if err != nil {
		return nil, err
	}

	// Prepare payload.
	payload := map[string]interface{}{
		"source":          "google_images",
		"domain":          opt.Domain,
		"query":           url,
		"start_page":      opt.StartPage,
		"pages":           opt.Pages,
		"locale":          opt.Locale,
		"geo_location":    &opt.GeoLocation,
		"user_agent_type": opt.UserAgent,
		"render":          opt.Render,
		"callback_url":    opt.CallbackURL,
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
	go c.PollJobStatus(jobID, false, responseChan, errChan)

	err = <-errChan
	if err != nil {
		return nil, err
	}

	return responseChan, nil
}

// ScrapeGoogleTrendsExplore scrapes google with google_trends_explore as source with async polling runtime.
func (c *SerpClientAsync) ScrapeGoogleTrendsExplore(
	query string,
	opts ...*GoogleTrendsExploreOpts,
) (chan *Response, error) {
	responseChan := make(chan *Response)
	errChan := make(chan error)

	// Prepare options.
	opt := &GoogleTrendsExploreOpts{}
	if len(opts) > 0 && opts[len(opts)-1] != nil {
		opt = opts[len(opts)-1]
	}

	// Initialize the context map apply each provided context modifier function.
	context := make(ContextOption)
	for _, modifier := range opt.Context {
		modifier(context)
	}

	// Set defaults.
	SetDefaultUserAgent(&opt.UserAgent)

	// Check validity of parameters.
	err := opt.checkParameterValidity(context)
	if err != nil {
		return nil, err
	}

	// Prepare payload.
	payload := map[string]interface{}{
		"source":       "google_trends_explore",
		"query":        query,
		"geo_location": &opt.GeoLocation,
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
		"callback_url":    opt.CallbackURL,
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
	go c.PollJobStatus(jobID, false, responseChan, errChan)

	err = <-errChan
	if err != nil {
		return nil, err
	}

	return responseChan, nil
}
