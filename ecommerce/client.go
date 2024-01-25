package ecommerce

import (
	"net/http"

	"github.com/mslmio/oxylabs-sdk-go/internal"
)

type EcommerceClient struct {
	C *internal.Client
}

// Init for Sync runtime model.
func Init(
	username string,
	password string,
) *EcommerceClient {
	return &EcommerceClient{
		C: &internal.Client{
			BaseUrl: internal.SyncBaseUrl,
			ApiCredentials: &internal.ApiCredentials{
				Username: username,
				Password: password,
			},
			HttpClient: &http.Client{},
		},
	}
}

type EcommerceClientAsync struct {
	C *internal.Client
}

// Init for Async runtime model.
func InitAsync(
	username string,
	password string,
) *EcommerceClientAsync {
	return &EcommerceClientAsync{
		C: &internal.Client{
			BaseUrl: internal.AsyncBaseUrl,
			ApiCredentials: &internal.ApiCredentials{
				Username: username,
				Password: password,
			},
			HttpClient: &http.Client{},
		},
	}
}
