package proxy

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
)

// Init for Proxy runtime model.
func Init(
	username string,
	password string,
) (*http.Client, error) {
	// Prepare proxy url.
	proxyUrl, err := url.Parse(
		fmt.Sprintf(
			"http://%s:%s@realtime.oxylabs.io:60000",
			username,
			password,
		),
	)
	if err != nil {
		return nil, fmt.Errorf("error parsing proxy url: %v", err)
	}

	// Prepare custom transport.
	customTransport := &http.Transport{Proxy: http.ProxyURL(proxyUrl)}
	customTransport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	client := &http.Client{Transport: customTransport}

	return client, nil
}
