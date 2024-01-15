package internal

import (
	"fmt"
	"net/url"
	"strings"
)

// Checks if the parameter is in the list of accepted parameters.
func InList[T comparable](val T, list []T) bool {
	for _, item := range list {
		if item == val {
			return true
		}
	}
	return false
}

// Checks if the provided URL is valid.
func ValidateUrl(
	inputUrl string,
	host string,
) error {
	// Check if the URL is empty.
	if inputUrl == "" {
		return fmt.Errorf("URL parameter is empty")
	}

	// Parse the URL.
	parsedUrl, err := url.ParseRequestURI(inputUrl)
	if err != nil {
		return fmt.Errorf("failed to parse URL: %v", err)
	}

	// Check if the scheme (protocol) is present and not empty.
	if parsedUrl.Scheme == "" {
		return fmt.Errorf("URL is missing scheme")
	}

	// Check if the host is present and not empty.
	if parsedUrl.Host == "" {
		return fmt.Errorf("URL is missing a host")
	}

	// Check if the host matches the expected domain or host.
	if !strings.Contains(parsedUrl.Host, host) {
		return fmt.Errorf("URL does not belong to %s", host)
	}

	return nil
}
