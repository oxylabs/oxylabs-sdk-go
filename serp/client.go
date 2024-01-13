package serp

import (
	"net/http"

	"github.com/mslmio/oxylabs-sdk-go/internal"
)

type SerpClient struct {
	InternalClient *internal.Client
}

// Init for Sync runtime model.
func Init(
	username string,
	password string,
) *SerpClient {
	return &SerpClient{
		InternalClient: &internal.Client{
			BaseURL: "https://realtime.oxylabs.io/v1/queries",
			ApiCredentials: &internal.ApiCredentials{
				Username: username,
				Password: password,
			},
			HttpClient: &http.Client{},
		},
	}
}

type SerpClientAsync struct {
	InternalClient *internal.Client
}

// Init for Async runtime model.
func InitAsync(
	username string,
	password string,
) *SerpClientAsync {
	return &SerpClientAsync{
		InternalClient: &internal.Client{
			BaseURL: "https://data.oxylabs.io/v1/queries",
			ApiCredentials: &internal.ApiCredentials{
				Username: username,
				Password: password,
			},
			HttpClient: &http.Client{},
		},
	}
}
