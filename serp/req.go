package serp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Request to the API.
func (c *SerpClient) Req(
	jsonPayload []byte,
	parse bool,
) (interface{}, error) {
	// Prepare requst.
	request, _ := http.NewRequest(
		"POST",
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
		return nil, fmt.Errorf("error with status code: %s : %s", response.Status, responseBody)
	}

	// Unmarshal the JSON object.
	var resp interface{}
	if parse {
		resp = &ParseTrueResponse{}
	} else {
		resp = &ParseFalseResponse{}
	}
	if err := json.Unmarshal(responseBody, resp); err != nil {
		return nil, fmt.Errorf("failed to parse JSON object: %v", err)
	}

	// Use type assertion to check the type and set status fields.
	switch r := resp.(type) {
	case *ParseTrueResponse:
		r.StatusCode = response.StatusCode
		r.Status = response.Status
	case *ParseFalseResponse:
		r.StatusCode = response.StatusCode
		r.Status = response.Status
	default:
		return nil, fmt.Errorf("unexpected response type")
	}

	return resp, nil
}
