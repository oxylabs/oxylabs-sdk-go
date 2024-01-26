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

// Helper function to make a POST req and retrieve the Job ID.
func (c *Client) GetJobID(
	jsonPayload []byte,
) (string, error) {
	req, _ := http.NewRequest(
		"POST",
		c.BaseUrl,
		bytes.NewBuffer(jsonPayload),
	)
	req.Header.Add("Content-type", "application/json")
	req.SetBasicAuth(
		c.ApiCredentials.Username,
		c.ApiCredentials.Password,
	)
	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("error performing req: %v", err)
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading resp body: %v", err)
	}
	resp.Body.Close()

	// Unmarshal into job.
	job := &Job{}
	if err = json.Unmarshal(respBody, &job); err != nil {
		return "", fmt.Errorf("error unmarshalling job resp body: %v", err)
	}

	return job.ID, nil
}

// Helper function for getting the http response from the request.
func (c *Client) GetHttpResp(
	jobID string,
	httpChan chan *http.Response,
	errChan chan error,
) {
	req, _ := http.NewRequest(
		"GET",
		fmt.Sprintf("https://data.oxylabs.io/v1/queries/%s/results", jobID),
		nil,
	)
	req.Header.Add("Content-type", "application/json")
	req.SetBasicAuth(
		c.ApiCredentials.Username,
		c.ApiCredentials.Password,
	)
	resp, err := c.HttpClient.Do(req)
	if err != nil {
		errChan <- err
		close(httpChan)
		return
	}

	// Return.
	close(errChan)
	httpChan <- resp
}

// PollJobStatus polls the job status and manages the resp/error channels.
// ctx is the context of the req.
// jsonPayload is the payload for the req.
// pollInterval is the time to wait between each subsequent polling req.
// httpRespChan and errChan are the channels for the http resp and error respectively.
func (c *Client) PollJobStatus(
	ctx context.Context,
	jobID string,
	pollInterval time.Duration,
	httpRespChan chan *http.Response,
	errChan chan error,
) {
	// Add default timeout if ctx has no deadline.
	if _, ok := ctx.Deadline(); !ok {
		context, cancel := context.WithTimeout(ctx, DefaultTimeout)
		defer cancel()
		ctx = context
	}

	// Set wait time between requests.
	sleepTime := DefaultPollInterval
	if pollInterval != 0 {
		sleepTime = pollInterval
	}

	for {
		// Perform a req to query job status.
		req, _ := http.NewRequest(
			"GET",
			fmt.Sprintf("https://data.oxylabs.io/v1/queries/%s", jobID),
			nil,
		)
		req.Header.Add("Content-type", "application/json")
		req.SetBasicAuth(
			c.ApiCredentials.Username,
			c.ApiCredentials.Password,
		)
		resp, err := c.HttpClient.Do(req)
		if err != nil {
			errChan <- err
			close(httpRespChan)
			return
		}

		// Read the resp body into a buffer.
		respBody, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			err = fmt.Errorf("error reading resp body: %v", err)
			errChan <- err
			close(httpRespChan)
			return
		}

		// Unmarshal into job.
		job := &Job{}
		if err = json.Unmarshal(respBody, &job); err != nil {
			err = fmt.Errorf("error unmarshalling job resp body: %v", err)
			errChan <- err
			close(httpRespChan)
			return
		}

		// Check job status.
		if job.Status == "done" {
			c.GetHttpResp(job.ID, httpRespChan, errChan)
			return
		} else if job.Status == "faulted" {
			err = fmt.Errorf("there was an error processing your query")
			errChan <- err
			close(httpRespChan)
			return
		}

		select {
		case <-ctx.Done():
			err = fmt.Errorf("timeout exceeded")
			errChan <- err
			close(httpRespChan)
			return
		default:
			time.Sleep(sleepTime)
		}
	}
}

// Job struct to get job id and status for the async polling.
type Job struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}
