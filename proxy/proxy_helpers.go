package proxy

import (
	"net/http"

	"github.com/mslmio/oxylabs-sdk-go/oxylabs"
)

// AddUserAgentHeader adds the user_agent_type header to the req.
func AddUserAgentHeader(req *http.Request, userAgent oxylabs.UserAgent) {
	req.Header.Add("x-oxylabs-user-agent-type", string(userAgent))
}

// AddRenderHeader adds the render header to the req.
func AddRenderHeader(req *http.Request, render oxylabs.Render) {
	req.Header.Add("x-oxylabs-render", string(render))
}

// AddParseHeader adds the parse and parser headers to the req.
func AddParseHeader(req *http.Request, parser string) {
	req.Header.Add("x-oxylabs-parse", "1")
	req.Header.Add("x-oxylabs-parser", parser)
}
