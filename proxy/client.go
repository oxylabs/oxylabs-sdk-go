package proxy

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/oxylabs/oxylabs-sdk-go/internal"
)

// Init initializes and returns an HTTP client configured with Oxylabs proxy settings.
func Init(
	username string,
	password string,
) (*http.Client, error) {
	// Encode username and password.
	encodedUsername := url.QueryEscape(username)
	encodedPassword := url.QueryEscape(password)

	// Prepare proxy URL.
	proxyUrl, err := url.Parse(
		fmt.Sprintf(
			"https://%s:%s@realtime.oxylabs.io:60000",
			encodedUsername,
			encodedPassword,
		),
	)
	if err != nil {
		return nil, fmt.Errorf("error parsing proxy URL: %v", err)
	}

	// Prepare custom transport.
	customTransport := &http.Transport{Proxy: http.ProxyURL(proxyUrl)}
	customTransport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	client := &http.Client{Transport: customTransport}

	return client, nil
}

func NewRequest(method, url string, body io.Reader) (*http.Request, error) {
	return internal.NewRequest(method, url, body)
}
