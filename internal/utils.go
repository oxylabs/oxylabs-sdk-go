package internal

import (
	"fmt"
	"net/url"
	"runtime"
	"strings"
)

var (
	sdkVersion    = "1.0.0" // SDK version. Needs to be updated manually along with new tag.
	sdkIdentifier = fmt.Sprintf(identifierTmpl, sdkVersion, runtime.Version(), runtime.GOOS, runtime.GOARCH)
)

// sdkIdentifier template. E.g. "oxylabs-sdk-go/1.0.0 (go1.22.6; linux/amd64)"
const identifierTmpl = "oxylabs-sdk-go/%s (%s; %s/%s)"

// InList checks if a value is present in the given slice.
func InList[T comparable](val T, list []T) bool {
	for _, item := range list {
		if item == val {
			return true
		}
	}

	return false
}

// ValidateUrl validates non-empty URL's scheme, host, and matches expected domain or host.
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
