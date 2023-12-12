package serp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/mslmio/oxylabs-sdk-go/oxylabs"
)

// ScrapeBingSearch scrapes bing with bing_search as source with async polling runtime.
func (c *SerpClientAsync) ScrapeBingSearch(
	query string,
	opts ...*BingSearchOpts,
) (chan *Response, error) {
	responseChan := make(chan *Response)
	errChan := make(chan error)

	// Prepare options
	opt := &BingSearchOpts{}
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
		"source":          "bing_search",
		"domain":          opt.Domain,
		"query":           query,
		"start_page":      opt.StartPage,
		"pages":           opt.Pages,
		"limit":           opt.Limit,
		"locale":          opt.Locale,
		"geo_location":    opt.GeoLocation,
		"user_agent_type": opt.UserAgent,
		"callback_url":    opt.CallbackUrl,
		"render":          opt.Render,
	}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("error marshalling payload: %v", err)
	}

	request, _ := http.NewRequest(
		"POST",
		c.BaseUrl,
		bytes.NewBuffer(jsonPayload),
	)

	request.Header.Add("Content-type", "application/json")
	request.SetBasicAuth(c.ApiCredentials.Username, c.ApiCredentials.Password)
	response, err := c.HttpClient.Do(request)
	if err != nil {
		return nil, err
	}

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}
	response.Body.Close()

	// Unmarshal into job.
	job := &Job{}
	json.Unmarshal(responseBody, &job)

	go func() {
		startNow := time.Now()

		for {
			request, _ = http.NewRequest(
				"GET",
				fmt.Sprintf("https://data.oxylabs.io/v1/queries/%s", job.ID),
				nil,
			)
			request.Header.Add("Content-type", "application/json")
			request.SetBasicAuth(c.ApiCredentials.Username, c.ApiCredentials.Password)
			response, err = c.HttpClient.Do(request)
			if err != nil {
				errChan <- err
				close(responseChan)
				return
			}

			responseBody, err = io.ReadAll(response.Body)
			if err != nil {
				err = fmt.Errorf("error reading response body: %v", err)
				errChan <- err
				close(responseChan)
				return
			}
			response.Body.Close()

			json.Unmarshal(responseBody, &job)

			if job.Status == "done" {
				JobId := job.ID
				request, _ = http.NewRequest(
					"GET",
					fmt.Sprintf("https://data.oxylabs.io/v1/queries/%s/results", JobId),
					nil,
				)
				request.Header.Add("Content-type", "application/json")
				request.SetBasicAuth(c.ApiCredentials.Username, c.ApiCredentials.Password)
				response, err = c.HttpClient.Do(request)
				if err != nil {
					errChan <- err
					close(responseChan)
					return
				}

				// Read the response body into a buffer.
				responseBody, err := io.ReadAll(response.Body)
				if err != nil {
					err = fmt.Errorf("error reading response body: %v", err)
					errChan <- err
					close(responseChan)
					return
				}
				response.Body.Close()

				// Send back error message.
				if response.StatusCode != 200 {
					err = fmt.Errorf("error with status code %s: %s", response.Status, responseBody)
					errChan <- err
					close(responseChan)
					return
				}

				// Unmarshal the JSON object.
				resp := &Response{}
				if err := resp.UnmarshalJSON(responseBody); err != nil {
					err = fmt.Errorf("failed to parse JSON object: %v", err)
					errChan <- err
					close(responseChan)
					return
				}
				resp.StatusCode = response.StatusCode
				resp.Status = response.Status
				close(errChan)
				responseChan <- resp
			} else if job.Status == "faulted" {
				err = fmt.Errorf("There was an error processing your query")
				errChan <- err
				close(responseChan)
				return
			}

			if time.Since(startNow) > oxylabs.DefaultTimeout {
				err = fmt.Errorf("timeout exceeded: %v", oxylabs.DefaultTimeout)
				errChan <- err
				close(responseChan)
				return
			}

			time.Sleep(oxylabs.DefaultWaitTime)
		}
	}()

	err = <-errChan
	if err != nil {
		return nil, err
	}

	return responseChan, nil
}

// ScrapeBingUrl scrapes bing with bing as source with async polling runtime.
func (c *SerpClientAsync) ScrapeBingUrl(
	url string,
	opts ...*BingUrlOpts,
) (chan *Response, error) {
	responseChan := make(chan *Response)
	errChan := make(chan error)

	// Check validity of url.
	err := oxylabs.ValidateURL(url, "bing")
	if err != nil {
		return nil, err
	}

	// Prepare options.
	opt := &BingUrlOpts{}
	if len(opts) > 0 && opts[len(opts)-1] != nil {
		opt = opts[len(opts)-1]
	}

	// Set defaults.
	SetDefaultUserAgent(&opt.UserAgent)

	// Check validity of parameters.
	err = opt.checkParameterValidity()
	if err != nil {
		return nil, err
	}

	// Prepare payload.
	payload := map[string]interface{}{
		"source":          "bing",
		"url":             url,
		"user_agent_type": opt.UserAgent,
		"geo_location":    opt.GeoLocation,
		"render":          opt.Render,
		"callback_url":    opt.CallbackUrl,
	}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("error marshalling payload: %v", err)
	}

	request, _ := http.NewRequest(
		"POST",
		c.BaseUrl,
		bytes.NewBuffer(jsonPayload),
	)

	request.Header.Add("Content-type", "application/json")
	request.SetBasicAuth(c.ApiCredentials.Username, c.ApiCredentials.Password)
	response, err := c.HttpClient.Do(request)
	if err != nil {
		return nil, err
	}

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}
	response.Body.Close()

	// Unmarshal into job.
	job := &Job{}
	json.Unmarshal(responseBody, &job)

	go func() {
		startNow := time.Now()

		for {
			request, _ = http.NewRequest(
				"GET",
				fmt.Sprintf("https://data.oxylabs.io/v1/queries/%s", job.ID),
				nil,
			)
			request.Header.Add("Content-type", "application/json")
			request.SetBasicAuth(c.ApiCredentials.Username, c.ApiCredentials.Password)
			response, err = c.HttpClient.Do(request)
			if err != nil {
				errChan <- err
				close(responseChan)
				return
			}

			responseBody, err = io.ReadAll(response.Body)
			if err != nil {
				err = fmt.Errorf("error reading response body: %v", err)
				errChan <- err
				close(responseChan)
				return
			}
			response.Body.Close()

			json.Unmarshal(responseBody, &job)

			if job.Status == "done" {
				JobId := job.ID
				request, _ = http.NewRequest(
					"GET",
					fmt.Sprintf("https://data.oxylabs.io/v1/queries/%s/results", JobId),
					nil,
				)
				request.Header.Add("Content-type", "application/json")
				request.SetBasicAuth(c.ApiCredentials.Username, c.ApiCredentials.Password)
				response, err = c.HttpClient.Do(request)
				if err != nil {
					errChan <- err
					close(responseChan)
					return
				}

				// Read the response body into a buffer.
				responseBody, err := io.ReadAll(response.Body)
				if err != nil {
					err = fmt.Errorf("error reading response body: %v", err)
					errChan <- err
					close(responseChan)
					return
				}
				response.Body.Close()

				// Send back error message.
				if response.StatusCode != 200 {
					err = fmt.Errorf("error with status code %s: %s", response.Status, responseBody)
					errChan <- err
					close(responseChan)
					return
				}

				// Unmarshal the JSON object.
				resp := &Response{}
				if err := resp.UnmarshalJSON(responseBody); err != nil {
					err = fmt.Errorf("failed to parse JSON object: %v", err)
					errChan <- err
					close(responseChan)
					return
				}
				resp.StatusCode = response.StatusCode
				resp.Status = response.Status
				close(errChan)
				responseChan <- resp
			} else if job.Status == "faulted" {
				err = fmt.Errorf("There was an error processing your query")
				errChan <- err
				close(responseChan)
				return
			}

			if time.Since(startNow) > oxylabs.DefaultTimeout {
				err = fmt.Errorf("timeout exceeded: %v", oxylabs.DefaultTimeout)
				errChan <- err
				close(responseChan)
				return
			}

			time.Sleep(oxylabs.DefaultWaitTime)
		}
	}()

	err = <-errChan
	if err != nil {
		return nil, err
	}

	return responseChan, nil
}
