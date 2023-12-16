package serp

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/mslmio/oxylabs-sdk-go/oxylabs"
)

type ScrapeProxyOpts struct {
	UserAgent   oxylabs.UserAgent
	GeoLocation string
	Render      oxylabs.Render
	Parser      *string
}

// checkParameterValidity checks validity of google search parameters.
func (opt *ScrapeProxyOpts) checkParameterValidity() error {
	if opt.UserAgent != "" && !oxylabs.IsUserAgentValid(opt.UserAgent) {
		return fmt.Errorf("invalid user agent parameter: %v", opt.UserAgent)
	}

	if opt.Render != "" && !oxylabs.IsRenderValid(opt.Render) {
		return fmt.Errorf("invalid render parameter: %v", opt.Render)
	}

	return nil
}

// ScrapeProxyUrl scrapes via proxy endpoint.
func (c *SerpClientProxy) ScrapeProxyUrl(
	url string,
	opts ...*ScrapeProxyOpts,
) (*ResponseProxy, error) {
	// Prepare options.
	opt := &ScrapeProxyOpts{}
	if len(opts) > 0 && opts[len(opts)-1] != nil {
		opt = opts[len(opts)-1]
	}

	// Check validity of parameters.
	if err := opt.checkParameterValidity(); err != nil {
		return nil, err
	}

	// Prepare request.
	request, err := http.NewRequest(
		"GET",
		url,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	// If options are provided, add them to the request.
	if opt.UserAgent != "" {
		request.Header.Add("x-oxylabs-user-agent-type", string(opt.UserAgent))
	}
	if opt.GeoLocation != "" {
		request.Header.Add("x-oxylabs-geo-location", opt.GeoLocation)
	}
	if opt.Render != "" {
		request.Header.Add("x-oxylabs-render", string(opt.Render))
	}
	if opt.Parser != nil {
		request.Header.Add("x-oxylabs-parse", "1")
		request.Header.Add("x-oxylabs-parser", *opt.Parser)
	}

	request.SetBasicAuth(c.ApiCredentials.Username, c.ApiCredentials.Password)

	// Get response.
	response, err := c.HttpClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer response.Body.Close()

	// Read response body.
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	// Send back error message.
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("error with status code %s: %s", response.Status, responseBody)
	}

	// Prepare response.
	resp := &ResponseProxy{}
	if opt.Parser != nil {
		json.Unmarshal(responseBody, &resp.ContentParsed)
	} else {
		resp.Content = string(responseBody)
	}

	// Set status code and status.
	resp.StatusCode = response.StatusCode
	resp.Status = response.Status

	return resp, nil
}
