package serp

import (
	"encoding/json"
	"fmt"

	"github.com/mslmio/oxylabs-sdk-go/oxylabs"
)

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

// checkParameterValidity checks validity of google search parameters.
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

// checkParameterValidity checks validity of google url parameters.
func (opt *GoogleUrlOpts) checkParameterValidity() error {
	if !oxylabs.IsUserAgentValid(opt.UserAgent) {
		return fmt.Errorf("invalid user agent parameter: %v", opt.UserAgent)
	}

	if opt.Render != "" && !oxylabs.IsRenderValid(opt.Render) {
		return fmt.Errorf("invalid render parameter: %v", opt.Render)
	}

	return nil
}

// checkParameterValidity checks validity of google ads parameters.
func (opt *GoogleAdsOpts) checkParameterValidity(ctx ContextOption) error {
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

// checkParameterValidity checks validity of google suggestions parameters.
func (opt *GoogleSuggestionsOpts) checkParameterValidity() error {
	if !oxylabs.IsUserAgentValid(opt.UserAgent) {
		return fmt.Errorf("invalid user agent parameter: %v", opt.UserAgent)
	}

	if opt.Render != "" && !oxylabs.IsRenderValid(opt.Render) {
		return fmt.Errorf("invalid render parameter: %v", opt.Render)
	}

	return nil
}

// checkParameterValidity checks validity of google hotels parameters.
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

	return nil
}

// checkParameterValidity checks validity of google travel hotels parameters.
func (opt *GoogleTravelHotelsOpts) checkParameterValidity(ctx ContextOption) error {

	if opt.Render != "" && !oxylabs.IsRenderValid(opt.Render) {
		return fmt.Errorf("invalid render parameter: %v", opt.Render)
	}

	if opt.Limit <= 0 || opt.Pages <= 0 || opt.StartPage <= 0 {
		return fmt.Errorf("limit, pages and start_page parameters must be greater than 0")
	}

	return nil
}

// checkParameterValidity checks validity of google trends explore parameters.
func (opt *GoogleTrendsExploreOpts) checkParameterValidity(ctx ContextOption) error {

	if !oxylabs.IsUserAgentValid(opt.UserAgent) {
		return fmt.Errorf("invalid user agent parameter: %v", opt.UserAgent)
	}

	if ctx["search_type"] != nil && !oxylabs.InList(ctx["search_type"].(string), AcceptedSearchTypeParameters) {
		return fmt.Errorf("invalid search_type parameter: %v", ctx["search_type"])
	}

	return nil
}

type GoogleSearchOpts struct {
	Domain      oxylabs.Domain
	StartPage   int
	Pages       int
	Limit       int
	Locale      oxylabs.Locale
	Geolocation string
	UserAgent   oxylabs.UserAgent
	Render      oxylabs.Render
	Parse       bool
	Context     []func(ContextOption)
}

// Scrapes Google via its search engine.
func (c *SerpClient) ScrapeGoogleSearch(
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

	// Prepare payload.
	var payload map[string]interface{}

	// If user sends limit_per_page context parameter, use it instead of limit, start_page and pages parameters.
	if context["limit_per_page"] != nil {
		payload = map[string]interface{}{
			"source":          "google_search",
			"domain":          opt.Domain,
			"query":           query,
			"geo_location":    opt.Geolocation,
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
					"key":   "limit_per_page",
					"value": context["limit_per_page"],
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
	} else {
		payload = map[string]interface{}{
			"source":          "google_search",
			"domain":          opt.Domain,
			"query":           query,
			"start_page":      opt.StartPage,
			"pages":           opt.Pages,
			"limit":           opt.Limit,
			"geo_location":    opt.Geolocation,
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
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("error marshalling payload: %v", err)
	}

	res, err := c.Req(jsonPayload, opt.Parse, "POST")
	if err != nil {
		return nil, err
	}

	return res, nil
}

type GoogleUrlOpts struct {
	GeoLocation string
	UserAgent   oxylabs.UserAgent
	Render      oxylabs.Render
	Parse       bool
	CallbackUrl string
}

// ScrapeGoogleUrl scrapes google vith google as source.
func (c *SerpClient) ScrapeGoogleUrl(
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
		"source":          "google",
		"url":             url,
		"user_agent_type": opt.UserAgent,
		"render":          opt.Render,
		"callback_url":    opt.CallbackUrl,
		"geo_location":    opt.GeoLocation,
		"parse":           opt.Parse,
	}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("error marshalling payload: %v", err)
	}

	res, err := c.Req(jsonPayload, opt.Parse, "POST")
	if err != nil {
		return nil, err
	}

	return res, nil
}

// checkParameterValidity checks validity of google images parameters.
func (opt *GoogleImagesOpts) checkParameterValidity(ctx ContextOption) error {

	if opt.Render != "" && !oxylabs.IsRenderValid(opt.Render) {
		return fmt.Errorf("invalid render parameter: %v", opt.Render)
	}

	if opt.Pages <= 0 || opt.StartPage <= 0 {
		return fmt.Errorf("limit, pages and start_page parameters must be greater than 0")
	}

	return nil
}

type GoogleAdsOpts struct {
	Domain      oxylabs.Domain
	StartPage   int
	Pages       int
	Limit       int
	Locale      string
	GeoLocation string
	UserAgent   oxylabs.UserAgent
	Render      oxylabs.Render
	Parse       bool
	Context     []func(ContextOption)
}

// SrcapeGoogleAds scrapes google via the google_ads source.
func (c *SerpClient) ScrapeGoogleAds(
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
		"geo_location":    opt.GeoLocation,
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

	res, err := c.Req(jsonPayload, opt.Parse, "POST")
	if err != nil {
		return nil, err
	}

	return res, nil
}

type GoogleSuggestionsOpts struct {
	Locale      string
	GeoLocation string
	UserAgent   oxylabs.UserAgent
	Render      oxylabs.Render
	CallbackUrl string
}

// ScrapeGoogleSuggestions scrapes google via the google_suggestions source.
func (c *SerpClient) ScrapeGoogleSuggestions(
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
		"source":          "google_suggestions",
		"query":           query,
		"geo_location":    opt.GeoLocation,
		"user_agent_type": opt.UserAgent,
		"render":          opt.Render,
		"callback_url":    opt.CallbackUrl,
	}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("error marshalling payload: %v", err)
	}

	res, err := c.Req(jsonPayload, false, "POST")
	if err != nil {
		return nil, err
	}

	return res, nil
}

type GoogleHotelsOpts struct {
	Domain          oxylabs.Domain
	StartPage       int
	Pages           int
	Limit           int
	Locale          string
	ResultsLanguage string
	GeoLocation     string
	UserAgent       oxylabs.UserAgent
	Render          oxylabs.Render
	CallbackURL     string
	Context         []func(ContextOption)
}

// ScrapeGoogleHotels scrapes google via the google_hotels source.
func (c *SerpClient) ScrapeGoogleHotels(
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
		"geo_location":     opt.GeoLocation,
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

	res, err := c.Req(jsonPayload, false, "POST")
	if err != nil {
		return nil, err
	}

	return res, nil
}

type GoogleTravelHotelsOpts struct {
	Domain      oxylabs.Domain
	StartPage   int
	Pages       int
	Limit       int
	Locale      string
	GeoLocation string
	Render      oxylabs.Render
	CallbackURL string
	Context     []func(ContextOption)
}

// ScrapeGoogleTravelHotels scrapes google via the google_travel_hotels source.
func (c *SerpClient) ScrapeGoogleTravelHotels(
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
	SetDefaultLimit(&opt.Limit)
	SetDefaultPages(&opt.Pages)

	// Check validity of parameters.
	err := opt.checkParameterValidity(context)
	if err != nil {
		return nil, err
	}

	// Prepare payload.
	payload := map[string]interface{}{
		"source":       "google_travel_hotels",
		"domain":       opt.Domain,
		"query":        query,
		"start_page":   opt.StartPage,
		"pages":        opt.Pages,
		"limit":        opt.Limit,
		"locale":       opt.Locale,
		"geo_location": opt.GeoLocation,
		"render":       opt.Render,
		"callback_url": opt.CallbackURL,
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

	res, err := c.Req(jsonPayload, false, "POST")
	if err != nil {
		return nil, err
	}

	return res, nil
}

type GoogleImagesOpts struct {
	Domain      oxylabs.Domain
	StartPage   int
	Pages       int
	Locale      string
	GeoLocation string
	UserAgent   oxylabs.UserAgent
	Render      oxylabs.Render
	CallbackURL string
	Context     []func(ContextOption)
}

// ScrapeGoogleImages scrapes google via the google_images source.
func (c *SerpClient) ScrapeGoogleImages(
	query string,
	opts ...*GoogleImagesOpts,
) (*Response, error) {
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
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("error marshalling payload: %v", err)
	}

	res, err := c.Req(jsonPayload, false, "POST")
	if err != nil {
		return nil, err
	}

	return res, nil
}

type GoogleTrendsExploreOpts struct {
	GeoLocation string
	Context     []func(ContextOption)
	UserAgent   oxylabs.UserAgent
	CallbackURL string
}

// ScrapeGoogleTrendsExplore scrapes google via the google_trends_explore source.
func (c *SerpClient) ScrapeGoogleTrendsExplore(
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

	// Prepare payload.
	payload := map[string]interface{}{
		"source":       "google_trends_explore",
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
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("error marshalling payload: %v", err)
	}
	fmt.Printf("%+v\n\n", payload)
	res, err := c.Req(jsonPayload, false, "POST")
	if err != nil {
		return nil, err
	}

	return res, nil
}
