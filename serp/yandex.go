package serp

import (
	"encoding/json"
	"fmt"

	"github.com/mslmio/oxylabs-sdk-go/oxylabs"
)

// Accepted parameters for yandex.
var yandexSearchAcceptedDomainParameters = []oxylabs.Domain{
	oxylabs.DOMAIN_COM,
	oxylabs.DOMAIN_RU,
	oxylabs.DOMAIN_UA,
	oxylabs.DOMAIN_BY,
	oxylabs.DOMAIN_KZ,
	oxylabs.DOMAIN_TR,
}
var yandexSearchAcceptedLocaleParameters = []oxylabs.Locale{
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

// checkParameterValidity checks validity of yandex search parameters.
func (opt *YandexSearchOpts) checkParameterValidity() error {
	if !oxylabs.InList(opt.Domain, yandexSearchAcceptedDomainParameters) {
		return fmt.Errorf("invalid domain parameter: %s", opt.Domain)
	}

	if opt.Locale != "" && !oxylabs.InList(opt.Locale, yandexSearchAcceptedLocaleParameters) {
		return fmt.Errorf("invalid locale parameter: %s", opt.Locale)
	}

	if !oxylabs.IsUserAgentValid(opt.UserAgent) {
		return fmt.Errorf("invalid user agent parameter: %v", opt.UserAgent)
	}

	return nil
}

// checkParameterValidity checks validity of yandex url parameters.
func (opt *YandexUrlOpts) checkParameterValidity() error {
	if !oxylabs.IsUserAgentValid(opt.UserAgent) {
		return fmt.Errorf("invalid user agent parameter: %v", opt.UserAgent)
	}

	if opt.Render != "" && !oxylabs.IsRenderValid(opt.Render) {
		return fmt.Errorf("invalid render parameter: %v", opt.Render)
	}

	return nil
}

type YandexSearchOpts struct {
	Domain      oxylabs.Domain
	StartPage   int
	Pages       int
	Limit       int
	Locale      oxylabs.Locale
	GeoLocation string
	UserAgent   oxylabs.UserAgent
	CallbackUrl string
}

// Scrapes Yandex via its search engine.
func (c *SerpClient) ScrapeYandexSearch(
	query string,
	opts ...*YandexSearchOpts,
) (*Response, error) {
	// Prepare options
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
		"source":          "yandex_search",
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

type YandexUrlOpts struct {
	UserAgent   oxylabs.UserAgent
	Render      oxylabs.Render
	CallbackUrl string
}

// Scrapes Yandex via provided url.
func (c *SerpClient) ScrapeYandexUrl(
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
		"source":          "yandex",
		"url":             url,
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
