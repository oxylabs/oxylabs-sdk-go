package oxylabs

type ContextOption map[string]interface{}

type PageLimit struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

type KeyValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// LimitPerPage sets the limits_per_page context option.
func LimitPerPage(limits []PageLimit) func(ContextOption) {
	return func(ctx ContextOption) {
		ctx["limit_per_page"] = limits
	}
}

// Content sets the limits_per_page context option.
func Content(content string) func(ContextOption) {
	return func(ctx ContextOption) {
		ctx["content"] = content
	}
}

// Cookies sets the cookies context option.
func Cookies(cookies []KeyValue) func(ContextOption) {
	return func(ctx ContextOption) {
		ctx["cookies"] = cookies
	}
}

// FollowRedirects sets the follow_redirects context option.
func FollowRedirects(follow bool) func(ContextOption) {
	return func(ctx ContextOption) {
		ctx["follow_redirects"] = follow
	}
}

// Headers sets the headers context option.
func Headers(headers map[string]string) func(ContextOption) {
	return func(ctx ContextOption) {
		ctx["headers"] = headers
	}
}

// HttpMethod sets the http_method context option.
func HttpMethod(method string) func(ContextOption) {
	return func(ctx ContextOption) {
		ctx["http_method"] = method
	}
}

// SessionId sets the session_id context option.
func SessionId(id string) func(ContextOption) {
	return func(ctx ContextOption) {
		ctx["session_id"] = id
	}
}

// SuccessfulStatusCodes sets the successful_status_codes context option.
func SuccessfulStatusCodes(codes []int) func(ContextOption) {
	return func(ctx ContextOption) {
		ctx["successful_status_codes"] = codes
	}
}

// ResultsLanguage sets the results_language context option.
func ResultsLanguage(lang string) func(ContextOption) {
	return func(ctx ContextOption) {
		ctx["results_language"] = lang
	}
}

// Filter sets the filter context option.
func Filter(filter int) func(ContextOption) {
	return func(ctx ContextOption) {
		ctx["filter"] = filter
	}
}

// Nfpr sets the nfpr context option.
func Nfpr(nfpr bool) func(ContextOption) {
	return func(ctx ContextOption) {
		ctx["nfpr"] = nfpr
	}
}

// SafeSearch sets the safe_search context option.
func SafeSearch(safeSearch bool) func(ContextOption) {
	return func(ctx ContextOption) {
		ctx["safe_search"] = safeSearch
	}
}

// Fpstate sets the fpstate context option.
func Fpstate(fpstate string) func(ContextOption) {
	return func(ctx ContextOption) {
		ctx["fpstate"] = fpstate
	}
}

// Tbm sets the tbm context option.
func Tbm(tbm string) func(ContextOption) {
	return func(ctx ContextOption) {
		ctx["tbm"] = tbm
	}
}

// Tbs sets the tbs context option.
func Tbs(tbs string) func(ContextOption) {
	return func(ctx ContextOption) {
		ctx["tbs"] = tbs
	}
}

// HotelOccupancy sets the hotel_occupancy context option.
func HotelOccupancy(num int) func(ContextOption) {
	return func(ctx ContextOption) {
		ctx["hotel_occupancy"] = num
	}
}

// HotelDates sets the hotel_dates context option.
func HotelDates(dates string) func(ContextOption) {
	return func(ctx ContextOption) {
		ctx["hotel_dates"] = dates
	}
}

// HotelClasses sets the hotel_classes context option.
func HotelClasses(classes []int) func(ContextOption) {
	return func(ctx ContextOption) {
		ctx["hotel_classes"] = classes
	}
}

// SearchType sets the search_type context option.
func SearchType(searchType string) func(ContextOption) {
	return func(ctx ContextOption) {
		ctx["search_type"] = searchType
	}
}

// DateFrom sets the date_from context option.
func DateFrom(dateFrom string) func(ContextOption) {
	return func(ctx ContextOption) {
		ctx["date_from"] = dateFrom
	}
}

// DateTo sets the date_to context option.
func DateTo(dateTo string) func(ContextOption) {
	return func(ctx ContextOption) {
		ctx["date_to"] = dateTo
	}
}

// CategoryId sets the category_id context option.
func CategoryId(categoryId int) func(ContextOption) {
	return func(ctx ContextOption) {
		ctx["category_id"] = categoryId
	}
}

// SortBy sets the sort_by context option.
func SortBy(sortBy string) func(ContextOption) {
	return func(ctx ContextOption) {
		ctx["sort_by"] = sortBy
	}
}

// MinPrice sets the min_price context option.
func MinPrice(minPrice int) func(ContextOption) {
	return func(ctx ContextOption) {
		ctx["min_price"] = minPrice
	}
}

// MaxPrice sets the max_price context option.
func MaxPrice(maxPrice int) func(ContextOption) {
	return func(ctx ContextOption) {
		ctx["max_price"] = maxPrice
	}
}

// MerchantId sets the category_id context option.
func MerchantId(merchantId int) func(ContextOption) {
	return func(ctx ContextOption) {
		ctx["merchant_id"] = merchantId
	}
}

// AutoselectVariant sets the autoselect_variant context option.
func AutoselectVariant(variant bool) func(ContextOption) {
	return func(ctx ContextOption) {
		ctx["autoselect_variant"] = variant
	}
}
