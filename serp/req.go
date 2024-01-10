package serp

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
)

// Request to the API.
// Ctx is the context of the request.
// JsonPayload is the payload for the request.
// Parse indicates whether to parse the response.
// ParseInstructions indicates whether to parse the response
// with custom parsing instructions.
// Method is the HTTP method of the request.
func (c *SerpClient) Req(
	ctx context.Context,
	jsonPayload []byte,
	parse bool,
	parseInstructions bool,
	method string,
) (*Response, error) {
	// Prepare request.
	request, _ := http.NewRequestWithContext(
		ctx,
		method,
		c.BaseUrl,
		bytes.NewBuffer(jsonPayload),
	)
	request.Header.Set("Content-Type", "application/json")
	request.SetBasicAuth(c.ApiCredentials.Username, c.ApiCredentials.Password)

	// Get response.
	response, err := c.HttpClient.Do(request)
	if e, ok := err.(net.Error); ok && e.Timeout() {
		return nil, fmt.Errorf("timeout error: %v", err)
	} else if err != nil {
		return nil, err
	}

	// Read the response body into a buffer.
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	// If status code not 200, return error.
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("error with status code %s: %s", response.Status, responseBody)
	}

	// Unmarshal the JSON object.
	resp := &Response{}
	resp.Parse = parse
	resp.ParseInstructions = parseInstructions
	if err := resp.UnmarshalJSON(responseBody); err != nil {
		return nil, fmt.Errorf("failed to parse JSON object: %v", err)
	}

	// Set status code and status.
	resp.StatusCode = response.StatusCode
	resp.Status = response.Status

	return resp, nil
}
