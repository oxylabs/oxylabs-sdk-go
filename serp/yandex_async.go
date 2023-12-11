package serp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func (c *SerpClientAsync) ScrapeYandexSearch(
	query string,
	opts ...*YandexSearchOpts,
) (chan *Response, error) {
	ResponseChan := make(chan *Response)

	opt := &YandexSearchOpts{}
	if len(opts) > 0 && opts[len(opts)-1] != nil {
		opt = opts[len(opts)-1]
	}

	// Set defaults.
	SetDefaultDomain(&opt.Domain)
	SetDefaultStartPage(&opt.StartPage)
	SetDefaultLimit(&opt.Limit)
	SetDefaultPages(&opt.Pages)
	SetDefaultUserAgent(&opt.UserAgent)

	// Check validity of parameters.
	err := opt.checkParameterValidity()
	if err != nil {
		return nil, err
	}

	// Prepare payload.
	payload := map[string]interface{}{
		"source":          "yandex_search",
		"domain":          opt.Domain,
		"query":           query,
		"start_page":      opt.StartPage,
		"pages":           opt.Pages,
		"limit":           opt.Limit,
		"locale":          opt.Locale,
		"geo_location":    opt.GeoLocation,
		"user_agent_type": opt.UserAgent,
		"callback_url":    opt.CallbackUrl,
	}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("error marshalling payload: %v", err)
	}

	client := &http.Client{}
	request, _ := http.NewRequest(
		"POST",
		c.BaseUrl,
		bytes.NewBuffer(jsonPayload),
	)

	request.Header.Add("Content-type", "application/json")
	request.SetBasicAuth(c.ApiCredentials.Username, c.ApiCredentials.Password)
	response, _ := client.Do(request)

	responseBody, _ := io.ReadAll(response.Body)

	// unmarshal into job object
	job := &Job{}
	json.Unmarshal(responseBody, &job)

	request, _ = http.NewRequest("GET",
		fmt.Sprintf("https://data.oxylabs.io/v1/queries/%s", job.ID),
		nil,
	)
	request.Header.Add("Content-type", "application/json")
	request.SetBasicAuth(c.ApiCredentials.Username, c.ApiCredentials.Password)

	go func() {
		for {
			response, _ = client.Do(request)
			responseBody, _ = io.ReadAll(response.Body)

			json.Unmarshal(responseBody, &job)

			if job.Status == "done" {
				JobId := job.ID
				request, _ = http.NewRequest("GET",
					fmt.Sprintf("https://data.oxylabs.io/v1/queries/%s/results", JobId),
					nil,
				)

				request.Header.Add("Content-type", "application/json")
				request.SetBasicAuth(c.ApiCredentials.Username, c.ApiCredentials.Password)
				response, _ = client.Do(request)

				// Read the response body into a buffer.
				responseBody, err := io.ReadAll(response.Body)
				if err != nil {
					err = fmt.Errorf("error reading response body: %v", err)
					return
				}

				// Send back error message.
				if response.StatusCode != 200 {
					err = fmt.Errorf("error with status code %s: %s", response.Status, responseBody)
					return
				}

				// Unmarshal the JSON object.
				resp := &Response{}
				if err := resp.UnmarshalJSON(responseBody); err != nil {
					err = fmt.Errorf("failed to parse JSON object: %v", err)
					return
				}
				resp.StatusCode = response.StatusCode
				resp.Status = response.Status
				ResponseChan <- resp
			}

			time.Sleep(2 * time.Second)
		}
	}()

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	return ResponseChan, nil
}
