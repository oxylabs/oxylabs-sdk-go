package serp

type ContextOption map[string]interface{}

type PageLimit struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

func LimitPerPage(limits []PageLimit) func(ContextOption) {
	return func(ctx ContextOption) {
		ctx["limit_per_page"] = limits
	}
}

func ResultsLanguage(lang string) func(ContextOption) {
	return func(ctx ContextOption) {
		ctx["results_language"] = lang
	}
}

func Filter(filter int) func(ContextOption) {
	return func(ctx ContextOption) {
		ctx["filter"] = filter
	}
}

func Nfpr(nfpr bool) func(ContextOption) {
	return func(ctx ContextOption) {
		ctx["nfpr"] = nfpr
	}
}

func SafeSearch(safeSearch bool) func(ContextOption) {
	return func(ctx ContextOption) {
		ctx["safe_search"] = safeSearch
	}
}

func Fpstate(fpstate string) func(ContextOption) {
	return func(ctx ContextOption) {
		ctx["fpstate"] = fpstate
	}
}

func Tbm(tbm string) func(ContextOption) {
	return func(ctx ContextOption) {
		ctx["tbm"] = tbm
	}
}

func Tbs(tbs string) func(ContextOption) {
	return func(ctx ContextOption) {
		ctx["tbs"] = tbs
	}
}

func HotelOccupancy(num int) func(ContextOption) {
	return func(ctx ContextOption) {
		ctx["hotel_occupancy"] = num
	}
}

func HotelDates(dates string) func(ContextOption) {
	return func(ctx ContextOption) {
		ctx["hotel_dates"] = dates
	}
}

func HotelClasses(classes []int) func(ContextOption) {
	return func(ctx ContextOption) {
		ctx["hotel_classes"] = classes
	}
}

func SearchType(searchType string) func(ContextOption) {
	return func(ctx ContextOption) {
		ctx["search_type"] = searchType
	}
}

func DateFrom(dateFrom string) func(ContextOption) {
	return func(ctx ContextOption) {
		ctx["date_from"] = dateFrom
	}
}

func DateTo(dateTo string) func(ContextOption) {
	return func(ctx ContextOption) {
		ctx["date_to"] = dateTo
	}
}

func CategoryId(categoryId int) func(ContextOption) {
	return func(ctx ContextOption) {
		ctx["category_id"] = categoryId
	}
}
