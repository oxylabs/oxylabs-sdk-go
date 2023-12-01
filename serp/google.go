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

	if ctx["tbm"] != nil && !inList(ctx["tbm"].(string), AcceptedTbmParameters) {
		return fmt.Errorf("invalid tbm parameter: %v", ctx["tbm"])
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

func (c *SerpClient) ScrapeGoogleSearch(
	query string,
	opts ...*GoogleSearchOpts,
) (interface{}, error) {
	// Prepare options.
	opt := &GoogleSearchOpts{}
	if len(opts) > 0 && opts[len(opts)-1] != nil {
		opt = opts[len(opts)-1]
	}

	// Initialize the context map apply each provided context modifier function.
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
			"geolocation":     opt.Geolocation,
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
			"geolocation":     opt.Geolocation,
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

	res, err := c.Req(jsonPayload, opt.Parse)
	if err != nil {
		return nil, err
	} else {
		return res, nil
	}
}
