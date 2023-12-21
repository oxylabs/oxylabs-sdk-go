# Oxylabs SDK Go

This is a Go SDK for the [Oxylabs](https://oxylabs.io) [Scraper APIs](https://developers.oxylabs.io/scraper-apis/getting-started).

This will help simplify integrating with Oxylabs's APIs, which can help you with retrieving search engine results (SERP), eCommerce data, real estate data, and more.

Some technical features include but are not limited to:

### Simplified Interface

Abstracts away complexities, offering a straightforward user interface for interacting with the Oxylabs SERP API.

### Automated Request Management

Streamlines the handling of API requests and responses for enhanced efficiency and reliability.

### Error Handling

Provides meaningful error messages and handles common API errors, simplifying troubleshooting.

### Result Parsing

Streamlines the process of extracting relevant data from SERP results, allowing developers to focus on application logic.

## Getting Started

You will need an Oxylabs API username and password which you can get by signing up at https://oxylabs.io. You can check things out with a free trial at https://oxylabs.io/products/scraper-api/serp for a week.

### Setting Up

This SDK requires a minimum version of `go 1.21`.

You can check your go version by running the following command in your preferred terminal:

```sh
go version
``` 

If you need to install or update go you can do so by following the steps mentioned [here](https://go.dev/doc/install).

#### Initialize Project

```sh
$ mkdir ~/oxylabs-sdk
$ cd ~/oxylabs-sdk
$ go mod init oxylabs-sdk 
```

#### Install SDK package

```sh
$ go get github.com/mslmio/oxylabs-sdk-go
```

### Quick Start

Basic usage of the SDK.

```go
package main

import (
	"fmt"

	"github.com/mslmio/oxylabs-sdk-go/serp"
)

func main() {
	// Set your Oxylabs API Credentials.
	const username = "username"
	const password = "password"

	// Initialize the SERP realtime client with your credentials.
	c := serp.Init(username, password)

	// Use `google_search` as a source to scrape Google with adidas as a query.
	res, err := c.ScrapeGoogleSearch(
		"adidas",
	)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Results: %+v\n", res)
}
```

## General Information

### Integration Methods

There are three integration method for the Oxylabs SERP API.
    - Realtime (Sync)
    - Push-Pull (Async)
    - Proxy Endpoint

To use either them you can just use the following init functions respectively:

- `serp.Init(username,password)`

- `serp.InitAsync(username,password)`

- `proxy.Init(username,password)`

Learn more about integration methods [on the official documentation](https://developers.oxylabs.io/scraper-apis/getting-started/integration-methods) and how this SDk uses them [here](#integration-methods-1).

### Sources

The Oxylabs SERP API scrapes according to the source provided to the API. There are currently four search engines you can scrape with the Oxylabs SERP API all with different sources.
| Search Engine | Sources
| ------------- | --------------
| **Google**    | `google`, `google_search`, `google_ads`, `google_hotels`, `google_travel_hotels`, `google_images`, `google_suggest`, `google_trends_explore`  
| **Yandex**    | `yandex`, `yandex_search`  
| **Bing**      | `bing`, `bing_search`
| **Baidu**     | `baidu`, `baidu_search`

Our SDK makes it easy for you, you just need to call the relevant function name from the client. For example if you wish to scrape Yandex with `yandex_search` as a source you
just need to invoke:

```go
res, err := c.ScrapeYandexSearch(
	"football",
)
```

### Query Parameters

Each source has different accepted query parameters. For a detailed list of accepted parameters by each source you can head over to https://developers.oxylabs.io/scraper-apis/serp-scraper-api.

This SDK provides you with the option to query with default parameters by not sending anything as the second argument as seen in the above example. Lets say we want to send in some query parameters it is as simple as:

```go
res, err := c.ScrapeYandexSearch(
	"football",
	&serp.YandexSearchOpts{
		StartPage: 1,
		Pages:     3,
		Limit:     4,
		Domain:    "com",
		Locale:    "en",
	},
)
```

### Configurable Options

For consistency and ease of use, this SDK provides a list of pre-defined commonly used parameter values as constants in our library. You can use them by importing the oxylabs package.

```go
import (
	"github.com/mslmio/oxylabs-sdk-go/oxylabs"
)
```

Currently these are available for the `Render` and`UserAgent` parameters. For the full list you can check `oxylabs/types.go`. You can send in these values as strings too.

These can be used like this:

```go
res, err := c.ScrapeGoogleSearch(
	"adidas",
	&serp.GoogleSearchOpts{
		UserAgent: oxylabs.UA_DESKTOP_CHROME, //desktop_chrome
		Render:    oxylabs.HTML,              // html
		Domain:    oxylabs.DOMAIN_COM,        // com
	},
)
```

### Context Options for Google sources

The SDK easily allows you to send in context options relevant to google sources. 

Here is an example of how you could send context options for Google Search:

```go
res, err := c.ScrapeGoogleSearch(
	"adidas",
	&serp.GoogleSearchOpts{
		Parse: true,
		Context: []func(serp.ContextOption){
			serp.ResultsLanguage("en"),
			serp.Filter(1),
			serp.Tbm("isch"),
			serp.LimitPerPage([]serp.PageLimit{{Page: 1, Limit: 1}, {Page: 2, Limit: 6}}),
		},
	},
)
```

## Integration Methods

### Realtime Integration

Realtime is a synchronous integration method. This means that upon sending your job submission request, **you will have to keep the connection open** until we successfully finish your job or return an error.

The **TTL** of Realtime connections is **150 seconds**. There may be rare cases where your connection times out before you receive a response from us, for example, if our system is under heavier-than-usual load or the job you submitted was extremely hard to complete:

### Push Pull(Polling) Integration <a id="push-pull"></a>

Push-Pull is an asynchronous integration method. This SDK implements this integration with a polling technique to poll the endpoint for results after a set interval of time.

Using it as straightforward as using the realtime integration. The only difference is that it will return a channel with the Response. Below is an example of this integration method:

```go
package main

import (
	"fmt"

	"github.com/mslmio/oxylabs-sdk-go/oxylabs"
	"github.com/mslmio/oxylabs-sdk-go/serp"
)

func main() {
	const username = "username"
	const password = "password"

	// Initialize the SERP push-pull client with your credentials.
	c := serp.InitAsync(username, password)

	ch, err := c.ScrapeGoogleAds(
		"adidas shoes",
		&serp.GoogleAdsOpts{
			UserAgent: oxylabs.UA_DESKTOP,
			Parse:     true,
		},
	)
	if err != nil {
		panic(err)
	}

	res := <-ch
	fmt.Printf("Results: %+v\n", res)
}
```

### Proxy Endpoint

This method is also synchronous (like Realtime), but instead of using our service via a RESTful interface, you **can use our endpoint like a proxy**. Use Proxy Endpoint if you've used proxies before and would just like to get unblocked content from us.

Since the parameters in this method are sent as as headers there are only a few parameters which this integration method accepts. You can find those parameters at
https://developers.oxylabs.io/scraper-apis/getting-started/integration-methods/proxy-endpoint#accepted-parameters.

The proxy endpoint integration is very open ended allowing many different use cases. To cater this, the user is provided a pre-configured `http.Client` and they can use it as they deem fit:

```go
package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/mslmio/oxylabs-sdk-go/oxylabs"
	"github.com/mslmio/oxylabs-sdk-go/proxy"
)

func main() {
	const username = "username"
	const password = "password"

	// Init returns an http client pre configured with the proxy settings.
	c, _ := proxy.Init(username, password)

	request, _ := http.NewRequest(
		"GET",
		"https://www.example.com",
		nil,
	)

	// Add relevant Headers.
	proxy.AddUserAgentHeader(request, oxylabs.UA_DESKTOP)
	proxy.AddRenderHeader(request, oxylabs.HTML)
	proxy.AddParseHeader(request, "google_search")

	request.SetBasicAuth(username, password)
	response, _ := c.Do(request)

	resp, _ := io.ReadAll(response.Body)
	fmt.Println(string(resp))
}
```

## About Oxylabs

Established in 2015, Oxylabs are a market-leading web intelligence collection platform, driven by the highest business, ethics, and compliance standards, enabling companies worldwide to unlock data-driven insights.

[![image](https://oxylabs.io/images/og-image.png)](https://oxylabs.io/)