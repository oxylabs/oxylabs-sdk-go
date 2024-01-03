package serp

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

// Request to the API.
func (c *SerpClient) Req(
	jsonPayload []byte,
	parse bool,
	method string,
) (*Response, error) {
	// Prepare requst.
	request, _ := http.NewRequest(
		method,
		c.BaseUrl,
		bytes.NewBuffer(jsonPayload),
	)
	request.Header.Set("Content-Type", "application/json")
	request.SetBasicAuth(c.ApiCredentials.Username, c.ApiCredentials.Password)

	// Get response.
	response, err := c.HttpClient.Do(request)
	if err != nil {
		return nil, err
	}

	// Read the response body into a buffer.
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	// Send back error message.
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("error with status code %s: %s", response.Status, responseBody)
	}

	// Unmarshal the JSON object.
	resp := &Response{}
	resp.Parse = parse
	if err := resp.UnmarshalJSON(responseBody); err != nil {
		return nil, fmt.Errorf("failed to parse JSON object: %v", err)
	}

	// Set status code and status.
	resp.StatusCode = response.StatusCode
	resp.Status = response.Status

	return resp, nil
}
