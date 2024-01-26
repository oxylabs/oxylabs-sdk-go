package internal

import (
	"bytes"
	"context"
	"fmt"
	"net"
	"net/http"
)

// Req to the API.
// Ctx is the context of the req.
// JsonPayload is the payload for the req.
// Method is the HTTP method of the req.
func (c *Client) Req(
	ctx context.Context,
	jsonPayload []byte,
	method string,
) (*http.Response, error) {
	// Prepare req.
	req, err := http.NewRequestWithContext(
		ctx,
		method,
		c.BaseUrl,
		bytes.NewBuffer(jsonPayload),
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(c.ApiCredentials.Username, c.ApiCredentials.Password)

	// Get resp.
	resp, err := c.HttpClient.Do(req)
	if e, ok := err.(net.Error); ok && e.Timeout() {
		return nil, fmt.Errorf("timeout error: %v", err)
	} else if err != nil {
		return nil, err
	}

	return resp, nil
}
