package serp

import (
	"fmt"
	"reflect"

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

// Function to check validity of yandex search parameters.
func (opt *YandexSearchOpts) checkParameterValidity() error {

	if opt.Domain != "" && !inSlice(opt.Domain, yandexSearchAcceptedDomainParameters) {
		return fmt.Errorf("invalid domain parameter: %s", opt.Domain)
	}

	if opt.Locale != "" && !inSlice(opt.Locale, yandexSearchAcceptedLocaleParameters) {
		return fmt.Errorf("invalid locale parameter: %s", opt.Locale)
	}

	if !oxylabs.IsUserAgentValid(string(opt.UserAgent)) {
		return fmt.Errorf("invalid user agent parameter: %v", opt.UserAgent)
	}

	return nil
}

// Function to check validity of yandex url parameters.
func (opt *YandexUrlOpts) checkParameterValidity() error {

	if !oxylabs.IsUserAgentValid(string(opt.UserAgent)) {
		return fmt.Errorf("invalid user agent parameter: %v", opt.UserAgent)
	}

	if opt.Render != "" && !oxylabs.IsRenderValid(opt.Render) {
		return fmt.Errorf("invalid render parameter: %v", opt.Render)
	}

	return nil
}

// Functions to check other sources wll be here.

// inSlice checks if a value is in the slice.
func inSlice(val interface{}, list interface{}) bool {
	switch reflect.TypeOf(list).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(list)

		for i := 0; i < s.Len(); i++ {
			if s.Index(i).Interface() == val {
				return true
			}
		}
	}

	return false
}
