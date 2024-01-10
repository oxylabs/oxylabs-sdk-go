package ecommerce

import (
	"net/http"
)

type ApiCredentials struct {
	Username string
	Password string
}

type EcommerceClient struct {
	HttpClient     *http.Client
	ApiCredentials *ApiCredentials
	BaseUrl        string
}

// Init for Sync runtime model.
func Init(
	username string,
	password string,
) *EcommerceClient {
	return &EcommerceClient{
		ApiCredentials: &ApiCredentials{
			Username: username,
			Password: password,
		},
		HttpClient: &http.Client{},
		BaseUrl:    "https://realtime.oxylabs.io/v1/queries",
	}
}

type EcommerceClientAsync struct {
	HttpClient     *http.Client
	ApiCredentials *ApiCredentials
	BaseUrl        string
}

// Init for Async runtime model.
func InitAsync(
	username string,
	password string,
) *EcommerceClientAsync {
	return &EcommerceClientAsync{
		ApiCredentials: &ApiCredentials{
			Username: username,
			Password: password,
		},
		HttpClient: &http.Client{},
		BaseUrl:    "https://data.oxylabs.io/v1/queries",
	}
}
