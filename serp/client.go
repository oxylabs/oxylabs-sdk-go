package serp

import (
	"net/http"

	"github.com/mslmio/oxylabs-sdk-go/internal"
)

const (
	SYNC_BASE_URL  = "https://realtime.oxylabs.io/v1/queries"
	ASYNC_BASE_URL = "https://data.oxylabs.io/v1/queries"
)

type SerpClient struct {
	C *internal.Client
}

// Init for Sync runtime model.
func Init(
	username string,
	password string,
) *SerpClient {
	return &SerpClient{
		C: &internal.Client{
			BaseUrl: SYNC_BASE_URL,
			ApiCredentials: &internal.ApiCredentials{
				Username: username,
				Password: password,
			},
			HttpClient: &http.Client{},
		},
	}
}

type SerpClientAsync struct {
	C *internal.Client
}

// Init for Async runtime model.
func InitAsync(
	username string,
	password string,
) *SerpClientAsync {
	return &SerpClientAsync{
		C: &internal.Client{
			BaseUrl: ASYNC_BASE_URL,
			ApiCredentials: &internal.ApiCredentials{
				Username: username,
				Password: password,
			},
			HttpClient: &http.Client{},
		},
	}
}
