package ecommerce

type ContextOption map[string]interface{}

// Nfpr sets the nfpr context option.
func Nfpr(nfpr bool) func(ContextOption) {
	return func(ctx ContextOption) {
		ctx["nfpr"] = nfpr
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
