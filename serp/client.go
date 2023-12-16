package serp

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
)

type ApiCredentials struct {
	Username string
	Password string
}

type SerpClient struct {
	HttpClient     *http.Client
	ApiCredentials *ApiCredentials
	BaseUrl        string
}

// Init for Sync runtime model.
func Init(
	username string,
	password string,
) *SerpClient {
	return &SerpClient{
		ApiCredentials: &ApiCredentials{
			Username: username,
			Password: password,
		},
		HttpClient: &http.Client{},
		BaseUrl:    "https://realtime.oxylabs.io/v1/queries",
	}
}

type SerpClientAsync struct {
	HttpClient     *http.Client
	ApiCredentials *ApiCredentials
	BaseUrl        string
}

// Init for Async runtime model.
func InitAsync(
	username string,
	password string,
) *SerpClientAsync {
	return &SerpClientAsync{
		ApiCredentials: &ApiCredentials{
			Username: username,
			Password: password,
		},
		HttpClient: &http.Client{},
		BaseUrl:    "https://data.oxylabs.io/v1/queries",
	}
}

type SerpClientProxy struct {
	HttpClient     *http.Client
	ApiCredentials *ApiCredentials
}

// Init for Proxy runtime model.
func InitProxy(
	username string,
	password string,
) *SerpClientProxy {
	// Prepare API credentials.
	apiCredentials := &ApiCredentials{
		Username: username,
		Password: password,
	}

	// Prepare proxy url.
	proxyUrl, err := url.Parse(
		fmt.Sprintf(
			"http://%s:%s@realtime.oxylabs.io:60000",
			apiCredentials.Username,
			apiCredentials.Password,
		),
	)
	if err != nil {
		fmt.Printf("error parsing proxy url: %v", err)
	}

	// Prepare custom transport.
	customTransport := &http.Transport{Proxy: http.ProxyURL(proxyUrl)}
	customTransport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	client := &http.Client{Transport: customTransport}

	return &SerpClientProxy{
		ApiCredentials: &ApiCredentials{
			Username: username,
			Password: password,
		},
		HttpClient: client,
	}
}
