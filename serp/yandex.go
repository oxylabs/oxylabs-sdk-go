package serp

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mslmio/oxylabs-sdk-go/oxylabs"
)

// Accepted parameters for yandex.
var YandexSearchAcceptedDomainParameters = []oxylabs.Domain{
	oxylabs.DOMAIN_COM,
	oxylabs.DOMAIN_RU,
	oxylabs.DOMAIN_UA,
	oxylabs.DOMAIN_BY,
	oxylabs.DOMAIN_KZ,
	oxylabs.DOMAIN_TR,
}
var YandexSearchAcceptedLocaleParameters = []oxylabs.Locale{
	oxylabs.LOCALE_EN,
	oxylabs.LOCALE_RU,
	oxylabs.LOCALE_BY,
	oxylabs.LOCALE_DE,
	oxylabs.LOCALE_FR,
	oxylabs.LOCALE_ID,
	oxylabs.LOCALE_KK,
	oxylabs.LOCALE_TT,
	oxylabs.LOCALE_TR,
	oxylabs.LOCALE_UK,
}

// checkParameterValidity checks validity of ScrapeYandexSearch parameters.
func (opt *YandexSearchOpts) checkParameterValidity() error {
	if !oxylabs.InList(opt.Domain, YandexSearchAcceptedDomainParameters) {
		return fmt.Errorf("invalid domain parameter: %s", opt.Domain)
	}

	if opt.Locale != "" && !oxylabs.InList(opt.Locale, YandexSearchAcceptedLocaleParameters) {
		return fmt.Errorf("invalid locale parameter: %s", opt.Locale)
	}

	if !oxylabs.IsUserAgentValid(opt.UserAgent) {
		return fmt.Errorf("invalid user agent parameter: %v", opt.UserAgent)
	}

	if opt.Limit <= 0 || opt.Pages <= 0 || opt.StartPage <= 0 {
		return fmt.Errorf("limit, pages and start_page parameters must be greater than 0")
	}

	return nil
}

// checkParameterValidity checks validity of ScrapeYandexUrl parameters.
func (opt *YandexUrlOpts) checkParameterValidity() error {
	if !oxylabs.IsUserAgentValid(opt.UserAgent) {
		return fmt.Errorf("invalid user agent parameter: %v", opt.UserAgent)
	}

	if opt.Render != "" && !oxylabs.IsRenderValid(opt.Render) {
		return fmt.Errorf("invalid render parameter: %v", opt.Render)
	}

	return nil
}

// YandexSearchOpts contains all the query parameters available for yandex_search.
type YandexSearchOpts struct {
	Domain            oxylabs.Domain
	StartPage         int
	Pages             int
	Limit             int
	Locale            oxylabs.Locale
	GeoLocation       string
	UserAgent         oxylabs.UserAgent
	CallbackUrl       string
	ParseInstructions *map[string]interface{}
}

// ScrapeYandexSearch scrapes yandex via Oxylabs SERP API with yandex_search as source.
func (c *SerpClient) ScrapeYandexSearch(
	query string,
	opts ...*YandexSearchOpts,
) (*Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), oxylabs.DefaultTimeout)
	defer cancel()

	return c.ScrapeYandexSearchCtx(ctx, query, opts...)
}

// ScrapeYandexSearchCtx scrapes yandex via Oxylabs SERP API with yandex_search as source.
// The provided context allows customization of the HTTP request, including setting timeouts.
func (c *SerpClient) ScrapeYandexSearchCtx(
	ctx context.Context,
	query string,
	opts ...*YandexSearchOpts,
) (*Response, error) {
	// Prepare options.
	opt := &YandexSearchOpts{}
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

	// Request.
	res, err := c.Req(ctx, jsonPayload, customParserFlag, customParserFlag, "POST")
	if err != nil {
		return nil, err
	}

	return res, nil
}

// YandexUrlOpts contains all the query parameters available for yandex.
type YandexUrlOpts struct {
	UserAgent         oxylabs.UserAgent
	Render            oxylabs.Render
	CallbackUrl       string
	ParseInstructions *map[string]interface{}
}

// ScrapeYandexUrl scrapes a yandex url via Oxylabs SERP API with yandex as source.
func (c *SerpClient) ScrapeYandexUrl(
	url string,
	opts ...*YandexUrlOpts,
) (*Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), oxylabs.DefaultTimeout)
	defer cancel()

	return c.ScrapeYandexUrlCtx(ctx, url, opts...)
}

// ScapeYandexUrlCtx scrapes a yandex url via Oxylabs SERP API with yandex as source.
// The provided context allows customization of the HTTP request, including setting timeouts.
func (c *SerpClient) ScrapeYandexUrlCtx(
	ctx context.Context,
	url string,
	opts ...*YandexUrlOpts,
) (*Response, error) {
	// Check validity of url.
	err := oxylabs.ValidateURL(url, "yandex")
	if err != nil {
		return nil, err
	}

	// Prepare options.
	opt := &YandexUrlOpts{}
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

	// Request.
	res, err := c.Req(ctx, jsonPayload, customParserFlag, customParserFlag, "POST")
	if err != nil {
		return nil, err
	}

	return res, nil
}
