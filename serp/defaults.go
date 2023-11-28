package serp

import (
	"github.com/mslmio/oxylabs-sdk-go/oxylabs"
)

func SetDefaultDomain(domain *oxylabs.Domain) {
	if *domain == "" {
		*domain = oxylabs.DOMAIN_COM
	}
}

func SetDefaultStartPage(startPage *int) {
	if *startPage == 0 {
		*startPage = 1
	}
}

func SetDefaultPages(pages *int) {
	if *pages == 0 {
		*pages = 1
	}
}

func SetDefaultLimit(limit *int) {
	if *limit == 0 {
		*limit = 10
	}
}

func SetDefaultUserAgent(userAgent *oxylabs.UserAgent) {
	if *userAgent == "" {
		*userAgent = oxylabs.UA_DESKTOP
	}
}
