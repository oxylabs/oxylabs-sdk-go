package proxy

import (
	"net/http"

	"github.com/mslmio/oxylabs-sdk-go/oxylabs"
)

// Helper functions to add headers to request object.s

func AddGeoLocationHeader(request *http.Request, geoLocation string) {
	request.Header.Add("x-oxylabs-geo-location", geoLocation)
}

func AddUserAgentHeader(request *http.Request, userAgent oxylabs.UserAgent) {
	request.Header.Add("x-oxylabs-user-agent-type", string(userAgent))
}

func AddRenderHeader(request *http.Request, render oxylabs.Render) {
	request.Header.Add("x-oxylabs-render", string(render))
}

func AddParseHeader(request *http.Request, parser string) {
	request.Header.Add("x-oxylabs-parse", "1")
	request.Header.Add("x-oxylabs-parser", parser)
}
