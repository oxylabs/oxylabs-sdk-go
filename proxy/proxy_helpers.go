package proxy

import (
	"net/http"

	"github.com/mslmio/oxylabs-sdk-go/oxylabs"
)

// AddUserAgentHeader adds the user_agent_type header to the request.
func AddUserAgentHeader(request *http.Request, userAgent oxylabs.UserAgent) {
	request.Header.Add("x-oxylabs-user-agent-type", string(userAgent))
}

// AddRenderHeader adds the render header to the request.
func AddRenderHeader(request *http.Request, render oxylabs.Render) {
	request.Header.Add("x-oxylabs-render", string(render))
}

// AddParseHeader adds the parse and parser headers to the request.
func AddParseHeader(request *http.Request, parser string) {
	request.Header.Add("x-oxylabs-parse", "1")
	request.Header.Add("x-oxylabs-parser", parser)
}
