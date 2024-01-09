package oxylabs

import (
	"fmt"
	"net/url"
	"strings"
	"time"
)

var (
	DefaultTimeout  = 50 * time.Second
	DefaultWaitTime = 2 * time.Second
)

func ValidateURL(
	inputURL string,
	host string,
) error {
	// Check if the URL is empty.
	if inputURL == "" {
		return fmt.Errorf("url parameter is empty")
	}

	// Parse the URL.
	parsedURL, err := url.ParseRequestURI(inputURL)
	if err != nil {
		return fmt.Errorf("failed to parse URL: %v", err)
	}

	// Check if the scheme (protocol) is present and not empty.
	if parsedURL.Scheme == "" {
		return fmt.Errorf("URL is missing scheme")
	}

	// Check if the host is present and not empty.
	if parsedURL.Host == "" {
		return fmt.Errorf("URL is missing a host")
	}

	// Check if the host matches the expected domain or host.
	if !strings.Contains(parsedURL.Host, host) {
		return fmt.Errorf("URL does not belong to %s", host)
	}

	return nil
}
