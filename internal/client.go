package internal

import "net/http"

type ApiCredentials struct {
	Username string
	Password string
}

type Client struct {
	BaseUrl        string
	ApiCredentials *ApiCredentials
	HttpClient     *http.Client
}
