package serp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *SerpClient) Req(
	jsonPayload []byte,
) (*Response, error) {
	// Prepare requst.
	request, _ := http.NewRequest("POST",
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

	// Unmarshal the JSON object.
	resp := &Response{}
	if err := json.Unmarshal(responseBody, resp); err != nil {
		return nil, fmt.Errorf("failed to parse JSON object: %v", err)
	}

	// Set status.
	resp.StatusCode = response.StatusCode
	resp.Status = response.Status

	return resp, nil
}
