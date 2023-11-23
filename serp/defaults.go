package serp

import "github.com/mslmio/oxylabs-sdk-go/oxylabs"

// Default values for Yandex search source.
func (opt *YandexSearchOpts) setDefaults() {
	if opt.Domain == "" {
		opt.Domain = oxylabs.DOMAIN_COM
	}
	if opt.StartPage == 0 {
		opt.StartPage = 1
	}
	if opt.Pages == 0 {
		opt.Pages = 1
	}
	if opt.Limit == 0 {
		opt.Limit = 10
	}
	if opt.UserAgent == "" {
		opt.UserAgent = oxylabs.UA_DESKTOP
	}
}

// Default values for Yandex url source.
func (opt *YandexUrlOpts) setDefaults() {
	if opt.UserAgent == "" {
		opt.UserAgent = oxylabs.UA_DESKTOP
	}
}
