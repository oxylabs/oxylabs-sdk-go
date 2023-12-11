package serp

import (
	"net/http"
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
