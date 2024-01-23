package serp

import (
	"net/http"

	"github.com/mslmio/oxylabs-sdk-go/internal"
)

type SerpResp struct {
	internal.Resp
}
type BingResp struct {
	internal.BingResp
}

type GoogleSearchResp struct {
	internal.GoogleSearchResp
}

type GoogleUrlResp struct {
	internal.GoogleUrlResp
}

type GoogleAdsResp struct {
	internal.GoogleAdsResp
}

type GoogleImagesResp struct {
	internal.GoogleImagesResp
}

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
			BaseUrl: internal.SyncBaseUrl,
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
			BaseUrl: internal.AsyncBaseUrl,
			ApiCredentials: &internal.ApiCredentials{
				Username: username,
				Password: password,
			},
			HttpClient: &http.Client{},
		},
	}
}
