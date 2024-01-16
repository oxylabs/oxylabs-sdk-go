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

// Helper function for handling resp parsing and error checking.
func (c *Client) GetResp(
	jobID string,
	parse bool,
	parseInstructions bool,
	respChan chan *Resp,
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
		close(respChan)
		return
	}

	// Read the resp body into a buffer.
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		err = fmt.Errorf("error reading resp body: %v", err)
		errChan <- err
		close(respChan)
		return
	}
	resp.Body.Close()

	// Check status code.
	if resp.StatusCode != 200 {
		err = fmt.Errorf("error with status code %s: %s", resp.Status, respBody)
		errChan <- err
		close(respChan)
		return
	}

	// Unmarshal the JSON object.
	res := &Resp{}
	res.Parse = parse
	if err := res.UnmarshalJSON(respBody); err != nil {
		err = fmt.Errorf("failed to parse JSON object: %v", err)
		errChan <- err
		close(respChan)
		return
	}
	res.StatusCode = resp.StatusCode
	res.Status = resp.Status
	close(errChan)
	respChan <- res
}

// PollJobStatus polls the job status and manages the resp/error channels.
// Ctx is the context of the req.
// JsonPayload is the payload for the req.
// Parse indicates whether to parse the resp.
// ParseInstructions indicates whether to parse the resp with custom parsing instructions.
// PollInterval is the time to wait between each subsequent polling req.
// respChan and errChan are the channels for the resp and error respectively.
func (c *Client) PollJobStatus(
	ctx context.Context,
	jobID string,
	parse bool,
	parseInstructions bool,
	pollInterval time.Duration,
	respChan chan *Resp,
	errChan chan error,
) {
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
			close(respChan)
			return
		}

		// Read the resp body into a buffer.
		respBody, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			err = fmt.Errorf("error reading resp body: %v", err)
			errChan <- err
			close(respChan)
			return
		}

		// Unmarshal into job.
		job := &Job{}
		if err = json.Unmarshal(respBody, &job); err != nil {
			err = fmt.Errorf("error unmarshalling job resp body: %v", err)
			errChan <- err
			close(respChan)
			return
		}

		// Check job status.
		if job.Status == "done" {
			c.GetResp(job.ID, parse, parseInstructions, respChan, errChan)
			return
		} else if job.Status == "faulted" {
			err = fmt.Errorf("there was an error processing your query")
			errChan <- err
			close(respChan)
			return
		}

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
