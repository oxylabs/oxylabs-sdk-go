package serp

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/mslmio/oxylabs-sdk-go/oxylabs"
)

// Accepted Parameters for context options in google.
var AcceptedTbmParameters = []string{
	"app",
	"bks",
	"blg",
	"dsc",
	"isch",
	"nws",
	"pts",
	"plcs",
	"rcp",
	"lcl",
}
var AcceptedSearchTypeParameters = []string{
	"web_search",
	"image_search",
	"google_shopping",
	"youtube_search",
}

// checkParameterValidity checks validity of ScrapeGoogleSearch parameters.
func (opt *GoogleSearchOpts) checkParameterValidity(ctx ContextOption) error {
	if !oxylabs.IsUserAgentValid(opt.UserAgent) {
		return fmt.Errorf("invalid user agent parameter: %v", opt.UserAgent)
	}

	if opt.Render != "" && !oxylabs.IsRenderValid(opt.Render) {
		return fmt.Errorf("invalid render parameter: %v", opt.Render)
	}

	if opt.Limit <= 0 || opt.Pages <= 0 || opt.StartPage <= 0 {
		return fmt.Errorf("limit, pages and start_page parameters must be greater than 0")
	}

	if ctx["tbm"] != nil && !oxylabs.InList(ctx["tbm"].(string), AcceptedTbmParameters) {
		return fmt.Errorf("invalid tbm parameter: %v", ctx["tbm"])
	}

	return nil
}

// checkParameterValidity checks validity of ScrapeGoogleUrl parameters.
func (opt *GoogleUrlOpts) checkParameterValidity() error {
	if !oxylabs.IsUserAgentValid(opt.UserAgent) {
		return fmt.Errorf("invalid user agent parameter: %v", opt.UserAgent)
	}

	if opt.Render != "" && !oxylabs.IsRenderValid(opt.Render) {
		return fmt.Errorf("invalid render parameter: %v", opt.Render)
	}

	return nil
}

// checkParameterValidity checks validity of ScrapeGoogleAds parameters.
func (opt *GoogleAdsOpts) checkParameterValidity(ctx ContextOption) error {
	if !oxylabs.IsUserAgentValid(opt.UserAgent) {
		return fmt.Errorf("invalid user agent parameter: %v", opt.UserAgent)
	}

	if opt.Render != "" && !oxylabs.IsRenderValid(opt.Render) {
		return fmt.Errorf("invalid render parameter: %v", opt.Render)
	}

	if opt.Pages <= 0 || opt.StartPage <= 0 {
		return fmt.Errorf("pages and start_page parameters must be greater than 0")
	}

	if ctx["tbm"] != nil && !oxylabs.InList(ctx["tbm"].(string), AcceptedTbmParameters) {
		return fmt.Errorf("invalid tbm parameter: %v", ctx["tbm"])
	}

	return nil
}

// checkParameterValidity checks validity of ScrapeGoogleSuggestions parameters.
func (opt *GoogleSuggestionsOpts) checkParameterValidity() error {
	if !oxylabs.IsUserAgentValid(opt.UserAgent) {
		return fmt.Errorf("invalid user agent parameter: %v", opt.UserAgent)
	}

	if opt.Render != "" && !oxylabs.IsRenderValid(opt.Render) {
		return fmt.Errorf("invalid render parameter: %v", opt.Render)
	}

	return nil
}

// checkParameterValidity checks validity of ScrapeGoogleHotels parameters.
func (opt *GoogleHotelsOpts) checkParameterValidity(ctx ContextOption) error {
	if !oxylabs.IsUserAgentValid(opt.UserAgent) {
		return fmt.Errorf("invalid user agent parameter: %v", opt.UserAgent)
	}

	if opt.Render != "" && !oxylabs.IsRenderValid(opt.Render) {
		return fmt.Errorf("invalid render parameter: %v", opt.Render)
	}

	if opt.Limit <= 0 || opt.Pages <= 0 || opt.StartPage <= 0 {
		return fmt.Errorf("limit, pages and start_page parameters must be greater than 0")
	}

	if ctx["hotel_occupancy"] != nil && ctx["hotel_occupancy"].(int) < 0 {
		return fmt.Errorf("invalid hotel_occupancy parameter: %v", ctx["hotel_occupancy"])
	}

	return nil
}

// checkParameterValidity checks validity of ScrapeGoogleTravelHotels parameters.
func (opt *GoogleTravelHotelsOpts) checkParameterValidity(ctx ContextOption) error {
	if !oxylabs.IsUserAgentValid(opt.UserAgent) {
		return fmt.Errorf("invalid user agent parameter: %v", opt.UserAgent)
	}

	if opt.Render != "" && !oxylabs.IsRenderValid(opt.Render) {
		return fmt.Errorf("invalid render parameter: %v", opt.Render)
	}

	if opt.StartPage <= 0 {
		return fmt.Errorf("start_page must be greater than 0")
	}

	if ctx["hotel_occupancy"] != nil && ctx["hotel_occupancy"].(int) < 0 {
		return fmt.Errorf("invalid hotel_occupancy parameter: %v", ctx["hotel_occupancy"])
	}

	if ctx["hotel_classes"] != nil {
		for _, value := range ctx["hotel_classes"].([]int) {
			if value < 2 || value > 5 {
				return fmt.Errorf("invalid hotel_classes parameter: %v", value)
			}
		}
	}

	return nil
}

// checkParameterValidity checks validity of ScrapeGoogleTrendsExplore parameters.
func (opt *GoogleTrendsExploreOpts) checkParameterValidity(ctx ContextOption) error {
	if !oxylabs.IsUserAgentValid(opt.UserAgent) {
		return fmt.Errorf("invalid user agent parameter: %v", opt.UserAgent)
	}

	if ctx["search_type"] != nil && !oxylabs.InList(ctx["search_type"].(string), AcceptedSearchTypeParameters) {
		return fmt.Errorf("invalid search_type parameter: %v", ctx["search_type"])
	}

	if ctx["category_id"] != nil && ctx["category_id"].(int) < 0 {
		return fmt.Errorf("invalid category_id")
	}

	return nil
}

// checkParameterValidity checks validity of google images parameters.
func (opt *GoogleImagesOpts) checkParameterValidity(ctx ContextOption) error {

	if opt.Render != "" && !oxylabs.IsRenderValid(opt.Render) {
		return fmt.Errorf("invalid render parameter: %v", opt.Render)
	}

	if opt.Pages <= 0 || opt.StartPage <= 0 {
		return fmt.Errorf("pages and start_page parameters must be greater than 0")
	}

	return nil
}

// GoogleSearchOpts contains all the query parameters available for google_search.
type GoogleSearchOpts struct {
	Domain            oxylabs.Domain
	StartPage         int
	Pages             int
	Limit             int
	Locale            oxylabs.Locale
	GeoLocation       string
	UserAgent         oxylabs.UserAgent
	Render            oxylabs.Render
	CallbackURL       string
	Parse             bool
	ParseInstructions *map[string]interface{}
	WaitTime          time.Duration
	Context           []func(ContextOption)
}

// ScrapeGoogleSearch scrapes google via Oxylabs SERP API with google_search as source.
func (c *SerpClient) ScrapeGoogleSearch(
	query string,
	opts ...*GoogleSearchOpts,
) (*Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), oxylabs.DefaultTimeout)
	defer cancel()

	return c.ScrapeGoogleSearchCtx(ctx, query, opts...)
}

// ScrapeGoogleSearchCtx scrapes google via Oxylabs SERP API with google_search as source.
// The provided context allows customization of the HTTP request, including setting timeouts.
func (c *SerpClient) ScrapeGoogleSearchCtx(
	ctx context.Context,
	query string,
	opts ...*GoogleSearchOpts,
) (*Response, error) {
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
		return nil, fmt.Errorf(
			"limit, start_page and pages parameters cannot be used together with limit_per_page context parameter",
		)
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
		"source":          oxylabs.GoogleSearch,
		"domain":          opt.Domain,
		"query":           query,
		"locale":          opt.Locale,
		"geo_location":    opt.GeoLocation,
		"user_agent_type": opt.UserAgent,
		"parse":           opt.Parse,
		"render":          opt.Render,
		"callback_url":    opt.CallbackURL,
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

	// Request.
	res, err := c.Req(ctx, jsonPayload, opt.Parse, customParserFlag, "POST")
	if err != nil {
		return nil, err
	}

	return res, nil
}

// GoogleUrlOpts contains all the query parameters available for google.
type GoogleUrlOpts struct {
	GeoLocation       string
	UserAgent         oxylabs.UserAgent
	Render            oxylabs.Render
	Parse             bool
	ParseInstructions *map[string]interface{}
	CallbackUrl       string
	WaitTime          time.Duration
}

// ScrapeGoogleUrl scrapes google via Oxylabs SERP API with google as source.
func (c *SerpClient) ScrapeGoogleUrl(
	url string,
	opts ...*GoogleUrlOpts,
) (*Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), oxylabs.DefaultTimeout)
	defer cancel()

	return c.ScrapeGoogleUrlCtx(ctx, url, opts...)
}

// ScrapeGoogleUrlCtx scrapes google via Oxylabs SERP API with google as source.
// The provided context allows customization of the HTTP request, including setting timeouts.
func (c *SerpClient) ScrapeGoogleUrlCtx(
	ctx context.Context,
	url string,
	opts ...*GoogleUrlOpts,
) (*Response, error) {
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

	// Request.
	res, err := c.Req(ctx, jsonPayload, opt.Parse, customParserFlag, "POST")
	if err != nil {
		return nil, err
	}

	return res, nil
}

// GoogleAdsOpts contains all the query parameters available for google_ads.
type GoogleAdsOpts struct {
	Domain            oxylabs.Domain
	StartPage         int
	Pages             int
	Locale            string
	GeoLocation       string
	UserAgent         oxylabs.UserAgent
	Render            oxylabs.Render
	CallbackURL       string
	Parse             bool
	ParseInstructions *map[string]interface{}
	WaitTime          time.Duration
	Context           []func(ContextOption)
}

// ScrapeGoogleAds scrapes google via Oxylabs SERP API with google_ads as source.
func (c *SerpClient) ScrapeGoogleAds(
	query string,
	opts ...*GoogleAdsOpts,
) (*Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), oxylabs.DefaultTimeout)
	defer cancel()

	return c.ScrapeGoogleAdsCtx(ctx, query, opts...)
}

// ScrapeGoogleAdsCtx scrapes google via Oxylabs SERP API with google_ads as source.
// The provided context allows customization of the HTTP request, including setting timeouts.
func (c *SerpClient) ScrapeGoogleAdsCtx(
	ctx context.Context,
	query string,
	opts ...*GoogleAdsOpts,
) (*Response, error) {
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
	SetDefaultPages(&opt.Pages)
	SetDefaultUserAgent(&opt.UserAgent)

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
		"callback_url":    opt.CallbackURL,
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

	// Request.
	res, err := c.Req(ctx, jsonPayload, opt.Parse, customParserFlag, "POST")
	if err != nil {
		return nil, err
	}

	return res, nil
}

// GoogleSuggestionsOpts contains all the query parameters available for google_shopping.
type GoogleSuggestionsOpts struct {
	Locale            string
	GeoLocation       string
	UserAgent         oxylabs.UserAgent
	Render            oxylabs.Render
	ParseInstructions *map[string]interface{}
	WaitTime          time.Duration
	CallbackUrl       string
}

// ScrapeGoogleSuggestions scrapes google via Oxylabs SERP API with google_suggestions as source.
func (c *SerpClient) ScrapeGoogleSuggestions(
	query string,
	opts ...*GoogleSuggestionsOpts,
) (*Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), oxylabs.DefaultTimeout)
	defer cancel()

	return c.ScrapeGoogleSuggestionsCtx(ctx, query, opts...)
}

// ScrapeGoogleSuggestionsCtx scrapes google via  Oxylabs SERP API with google_suggestions as source.
// The provided context allows customization of the HTTP request, including setting timeouts.
func (c *SerpClient) ScrapeGoogleSuggestionsCtx(
	ctx context.Context,
	query string,
	opts ...*GoogleSuggestionsOpts,
) (*Response, error) {
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

	// Request.
	res, err := c.Req(ctx, jsonPayload, customParserFlag, customParserFlag, "POST")
	if err != nil {
		return nil, err
	}

	return res, nil
}

// GoogleHotelsOpts contains all the query parameters available for google_hotels.
type GoogleHotelsOpts struct {
	Domain            oxylabs.Domain
	StartPage         int
	Pages             int
	Limit             int
	Locale            string
	GeoLocation       string
	UserAgent         oxylabs.UserAgent
	Render            oxylabs.Render
	CallbackURL       string
	ParseInstructions *map[string]interface{}
	WaitTime          time.Duration
	Context           []func(ContextOption)
}

// ScrapeGoogleHotels scrapes google via Oxylabs SERP API with google_hotels as source.
func (c *SerpClient) ScrapeGoogleHotels(
	query string,
	opts ...*GoogleHotelsOpts,
) (*Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), oxylabs.DefaultTimeout)
	defer cancel()

	return c.ScrapeGoogleHotelsCtx(ctx, query, opts...)
}

// ScrapeGoogleHotelsCtx scrapes google via the google_hotels source.
// The provided context allows customization of the HTTP request, including setting timeouts.
func (c *SerpClient) ScrapeGoogleHotelsCtx(
	ctx context.Context,
	query string,
	opts ...*GoogleHotelsOpts,
) (*Response, error) {
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
	setDefaultHotelOccupancy(context)

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
		"callback_url":    opt.CallbackURL,
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

	// Request.
	res, err := c.Req(ctx, jsonPayload, customParserFlag, customParserFlag, "POST")
	if err != nil {
		return nil, err
	}

	return res, nil
}

// GoogleTravelHotelsOpts contains all the query parameters available for google_travel_hotels.
type GoogleTravelHotelsOpts struct {
	Domain            oxylabs.Domain
	StartPage         int
	Locale            string
	GeoLocation       string
	UserAgent         oxylabs.UserAgent
	Render            oxylabs.Render
	CallbackURL       string
	ParseInstructions *map[string]interface{}
	WaitTime          time.Duration
	Context           []func(ContextOption)
}

// ScrapeGoogleTravelHotels scrapes google via Oxylabs SERP API with google_travel_hotels as source.
func (c *SerpClient) ScrapeGoogleTravelHotels(
	query string,
	opts ...*GoogleTravelHotelsOpts,
) (*Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), oxylabs.DefaultTimeout)
	defer cancel()

	return c.ScrapeGoogleTravelHotelsCtx(ctx, query, opts...)
}

// ScrapeGoogleTravelHotelsCtx scrapes google via Oxylabs SERP API with google_travel_hotels as source.
// The provided context allows customization of the HTTP request, including setting timeouts.
func (c *SerpClient) ScrapeGoogleTravelHotelsCtx(
	ctx context.Context,
	query string,
	opts ...*GoogleTravelHotelsOpts,
) (*Response, error) {
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

	// Request.
	res, err := c.Req(ctx, jsonPayload, customParserFlag, customParserFlag, "POST")
	if err != nil {
		return nil, err
	}

	return res, nil
}

// GoogleImagesOpts contains all the query parameters available for google_images.
type GoogleImagesOpts struct {
	Domain            oxylabs.Domain
	StartPage         int
	Pages             int
	Locale            string
	GeoLocation       string
	UserAgent         oxylabs.UserAgent
	Render            oxylabs.Render
	CallbackURL       string
	ParseInstructions *map[string]interface{}
	WaitTime          time.Duration
	Context           []func(ContextOption)
}

// ScrapeGoogleImages scrapes google via Oxylabs SERP API with google_images as source.
func (c *SerpClient) ScrapeGoogleImages(
	url string,
	opts ...*GoogleImagesOpts,
) (*Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), oxylabs.DefaultTimeout)
	defer cancel()

	return c.ScrapeGoogleImagesCtx(ctx, url, opts...)
}

// ScrapeGoogleImagesCtx scrapes google via Oxylabs SERP API with google_images as source.
// The provided context allows customization of the HTTP request, including setting timeouts.
func (c *SerpClient) ScrapeGoogleImagesCtx(
	ctx context.Context,
	url string,
	opts ...*GoogleImagesOpts,
) (*Response, error) {
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
		"source":          oxylabs.GoogleImages,
		"domain":          opt.Domain,
		"query":           url,
		"start_page":      opt.StartPage,
		"pages":           opt.Pages,
		"locale":          opt.Locale,
		"geo_location":    opt.GeoLocation,
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

	// Request.
	res, err := c.Req(ctx, jsonPayload, customParserFlag, customParserFlag, "POST")
	if err != nil {
		return nil, err
	}

	return res, nil
}

// GoogleTrendsExploreOpts contains all the query parameters available for google_trends_explore.
type GoogleTrendsExploreOpts struct {
	GeoLocation       string
	Context           []func(ContextOption)
	UserAgent         oxylabs.UserAgent
	CallbackURL       string
	ParseInstructions *map[string]interface{}
	WaitTime          time.Duration
}

// ScrapeGoogleTrendsExplore scrapes google via Oxylabs SERP API with google_trends_explore as source.
func (c *SerpClient) ScrapeGoogleTrendsExplore(
	query string,
	opts ...*GoogleTrendsExploreOpts,
) (*Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), oxylabs.DefaultTimeout)
	defer cancel()

	return c.ScrapeGoogleTrendsExploreCtx(ctx, query, opts...)
}

// ScrapeGoogleTrendsExploreCtx scrapes google via Oxylabs SERP API with google_trends_explore as source.
// The provided context allows customization of the HTTP request, including setting timeouts.
func (c *SerpClient) ScrapeGoogleTrendsExploreCtx(
	ctx context.Context,
	query string,
	opts ...*GoogleTrendsExploreOpts,
) (*Response, error) {
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
		"callback_url":    opt.CallbackURL,
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

	// Request.
	res, err := c.Req(ctx, jsonPayload, true, customParserFlag, "POST")
	if err != nil {
		return nil, err
	}

	return res, nil
}
