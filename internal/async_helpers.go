package internal

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Helper function to make a POST request and retrieve the Job ID.
func (c *Client) GetJobID(
	jsonPayload []byte,
) (string, error) {
	request, _ := http.NewRequest(
		"POST",
		c.BaseUrl,
		bytes.NewBuffer(jsonPayload),
	)
	request.Header.Add("Content-type", "application/json")
	request.SetBasicAuth(
		c.ApiCredentials.Username,
		c.ApiCredentials.Password,
	)
	response, err := c.HttpClient.Do(request)
	if err != nil {
		return "", fmt.Errorf("error performing request: %v", err)
	}

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %v", err)
	}
	response.Body.Close()

	// Unmarshal into job.
	job := &Job{}
	if err = json.Unmarshal(responseBody, &job); err != nil {
		return "", fmt.Errorf("error unmarshalling job response body: %v", err)
	}

	return job.ID, nil
}

// Helper function for handling response parsing and error checking.
func (c *Client) GetResponse(
	jobID string,
	parse bool,
	parseInstructions bool,
	respChan chan *Resp,
	errChan chan error,
) {
	request, _ := http.NewRequest(
		"GET",
		fmt.Sprintf("https://data.oxylabs.io/v1/queries/%s/results", jobID),
		nil,
	)
	request.Header.Add("Content-type", "application/json")
	request.SetBasicAuth(
		c.ApiCredentials.Username,
		c.ApiCredentials.Password,
	)
	response, err := c.HttpClient.Do(request)
	if err != nil {
		errChan <- err
		close(respChan)
		return
	}

	// Read the response body into a buffer.
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		err = fmt.Errorf("error reading response body: %v", err)
		errChan <- err
		close(respChan)
		return
	}
	response.Body.Close()

	// Check status code.
	if response.StatusCode != 200 {
		err = fmt.Errorf("error with status code %s: %s", response.Status, responseBody)
		errChan <- err
		close(respChan)
		return
	}

	// Unmarshal the JSON object.
	resp := &Resp{}
	resp.Parse = parse
	if err := resp.UnmarshalJSON(responseBody); err != nil {
		err = fmt.Errorf("failed to parse JSON object: %v", err)
		errChan <- err
		close(respChan)
		return
	}
	resp.StatusCode = response.StatusCode
	resp.Status = response.Status
	close(errChan)
	respChan <- resp
}

// PollJobStatus polls the job status and manages the response/error channels.
// Ctx is the context of the request.
// JsonPayload is the payload for the request.
// Parse indicates whether to parse the response.
// ParseInstructions indicates whether to parse the response with custom parsing instructions.
// PollInterval is the time to wait between each subsequent polling request.
// respChan and errChan are the channels for the response and error respectively.
func (c *Client) PollJobStatus(
	ctx context.Context,
	jobID string,
	parse bool,
	parseInstructions bool,
	pollInterval time.Duration,
	respChan chan *Resp,
	errChan chan error,
) {
	// Add default timeout if ctx has no deadline.
	if _, ok := ctx.Deadline(); !ok {
		context, cancel := context.WithTimeout(ctx, DefaultTimeout)
		defer cancel()
		ctx = context
	}

	for {
		// Perform a request to query job status.
		request, _ := http.NewRequest(
			"GET",
			fmt.Sprintf("https://data.oxylabs.io/v1/queries/%s", jobID),
			nil,
		)
		request.Header.Add("Content-type", "application/json")
		request.SetBasicAuth(
			c.ApiCredentials.Username,
			c.ApiCredentials.Password,
		)
		response, err := c.HttpClient.Do(request)
		if err != nil {
			errChan <- err
			close(respChan)
			return
		}

		// Read the response body into a buffer.
		responseBody, err := io.ReadAll(response.Body)
		response.Body.Close()
		if err != nil {
			err = fmt.Errorf("error reading response body: %v", err)
			errChan <- err
			close(respChan)
			return
		}

		// Unmarshal into job.
		job := &Job{}
		if err = json.Unmarshal(responseBody, &job); err != nil {
			err = fmt.Errorf("error unmarshalling job response body: %v", err)
			errChan <- err
			close(respChan)
			return
		}

		// Check job status.
		if job.Status == "done" {
			c.GetResponse(job.ID, parse, parseInstructions, respChan, errChan)
			return
		} else if job.Status == "faulted" {
			err = fmt.Errorf("there was an error processing your query")
			errChan <- err
			close(respChan)
			return
		}

		// Set wait time between requests.
		sleepTime := DefaultPollInterval
		if pollInterval != 0 {
			sleepTime = pollInterval
		}

		select {
		case <-ctx.Done():
			err = fmt.Errorf("timeout exceeded")
			errChan <- err
			close(respChan)
			return
		default:
			time.Sleep(sleepTime)
		}
	}
}
