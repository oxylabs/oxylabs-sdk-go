package internal

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
)

// Req to the API.
// Ctx is the context of the req.
// JsonPayload is the payload for the req.
// Parse indicates whether to parse the resp.
// ParseInstructions indicates whether to parse the resp
// with custom parsing instructions.
// Method is the HTTP method of the req.
func (c *Client) Req(
	ctx context.Context,
	jsonPayload []byte,
	parse bool,
	parseInstructions bool,
	method string,
) (*Resp, error) {
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
	defer resp.Body.Close()

	// Read the resp body into a buffer.
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// If status code not 200, return error.
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error with status code %s: %s", resp.Status, respBody)
	}

	// Unmarshal the JSON object.
	res := &Resp{}
	res.Parse = parse
	res.ParseInstructions = parseInstructions
	if err := res.UnmarshalJSON(respBody); err != nil {
		return nil, fmt.Errorf("failed to parse JSON object: %v", err)
	}

	// Set status code and status.
	res.StatusCode = resp.StatusCode
	res.Status = resp.Status

	return res, nil
}
