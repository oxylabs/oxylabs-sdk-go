package oxylabs

import (
	"fmt"
	"net/url"
	"strings"
)

func ValidateURL(
	inputURL string,
	host string,
) error {
	// Check if url is empty.
	if inputURL == "" {
		return fmt.Errorf("url parameter is empty")
	}

	// Parse the URL
	parsedURL, err := url.ParseRequestURI(inputURL)
	if err != nil {
		return fmt.Errorf("failed to parse URL: %v", err)
	}

	// Check if the scheme (protocol) is present and non-empty.
	if parsedURL.Scheme == "" {
		return fmt.Errorf("url is missing scheme")
	}

	// Check if the Host is present and non-empty.
	if parsedURL.Host == "" {
		return fmt.Errorf("url is missing a host")
	}

	// Check if the Host matches the expected domain/host.
	if !strings.Contains(parsedURL.Host, host) {
		return fmt.Errorf("url does not belong to %s", host)
	}

	return nil
}
