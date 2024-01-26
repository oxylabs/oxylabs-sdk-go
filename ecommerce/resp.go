package ecommerce

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Resp is the response struct for all ecommerce sources.
type Resp struct {
	Parse             bool      `json:"parse"`
	ParseInstructions bool      `json:"parse_instructions"`
	Results           []Results `json:"results"`
	Job               Job       `json:"job"`
	StatusCode        int       `json:"status_code"`
	Status            string    `json:"status"`
}

type Results struct {
	CustomContentParsed map[string]interface{}
	ContentParsed       Content
	Content             string
	CreatedAt           string `json:"created_at"`
	UpdatedAt           string `json:"updated_at"`
	Page                int    `json:"page"`
	Url                 string `json:"url"`
	JobID               string `json:"job_id"`
	StatusCode          int    `json:"status_code"`
	ParserType          string `json:"parser_type"`
}

type Content struct {
	Url                    string                         `json:"url"`
	Title                  string                         `json:"title"`
	Pages                  int                            `json:"pages"`
	Query                  string                         `json:"query"`
	Images                 interface{}                    `json:"images"`
	Variants               Variants                       `json:"variants"`
	Highlights             []string                       `json:"highlights"`
	Description            string                         `json:"description"`
	RelatedItems           RelatedItems                   `json:"related_items"`
	Specifications         Specifications                 `json:"specifications"`
	Page                   int                            `json:"page"`
	Errors                 interface{}                    `json:"_errors"`
	Results                Result                         `json:"results"`
	Rating                 float64                        `json:"rating"`
	Pricing                []Pricing                      `json:"pricing"`
	Ads                    []AmazonProductAds             `json:"ads"`
	Asin                   string                         `json:"asin"`
	Price                  float64                        `json:"price"`
	Stock                  string                         `json:"stock"`
	Coupon                 string                         `json:"coupon"`
	Category               []AmazonProductCategory        `json:"category"`
	Currency               string                         `json:"currency"`
	Delivery               []AmazonProductDelivery        `json:"delivery"`
	Warnings               []string                       `json:"_warnings,omitempty"`
	DealType               string                         `json:"deal_type"`
	PageType               string                         `json:"page_type"`
	PriceSns               int                            `json:"price_sns"`
	Variation              interface{}                    `json:"variation"`
	HasVideos              bool                           `json:"has_videos"`
	SalesRank              []AmazonProductSalesRank       `json:"sales_rank"`
	TopReview              string                         `json:"top_review"`
	AsinInUrl              string                         `json:"asin_in_url"`
	PriceUpper             float64                        `json:"price_upper"`
	PricingStr             string                         `json:"pricing_str"`
	PricingURL             string                         `json:"pricing_url"`
	DiscountEnd            string                         `json:"discount_end"`
	Manufacturer           string                         `json:"manufacturer"`
	MaxQuantity            int                            `json:"max_quantity"`
	PriceBuybox            float64                        `json:"price_buybox"`
	ProductName            string                         `json:"product_name"`
	BulletPoints           string                         `json:"bullet_points"`
	IsAddonItem            bool                           `json:"is_addon_item"`
	PriceInitial           int                            `json:"price_initial"`
	PricingCount           int                            `json:"pricing_count"`
	ReviewsCount           int                            `json:"reviews_count"`
	SNSDiscounts           []interface{}                  `json:"sns_discounts"`
	DeveloperInfo          []interface{}                  `json:"developer_info"`
	LightningDeal          interface{}                    `json:"lightning_deal"`
	PriceShipping          int                            `json:"price_shipping"`
	IsPrimePantry          bool                           `json:"is_prime_pantry"`
	ProductDetails         ProductDetails                 `json:"product_details"`
	FeaturedMerchant       []interface{}                  `json:"featured_merchant"`
	IsPrimeEligible        bool                           `json:"is_prime_eligible"`
	ProductDimensions      string                         `json:"product_dimensions"`
	RefurbishedProduct     AmazonRefurbishedProduct       `json:"refurbished_product"`
	AnsweredQuestionsCount int                            `json:"answered_questions_count"`
	RatingStarDistribution []AmazonRatingStarDistribution `json:"rating_star_distribution"`
	Reviews                []AmazonReviews                `json:"reviews"`
	Questions              AmazonQuestions                `json:"questions"`
	QuestionsTotal         int                            `json:"questions_total"`
	BusinessName           string                         `json:"business_name"`
	RecentFeedback         []RecentFeedback               `json:"recent_feedback"`
	BusinessAddress        string                         `json:"business_address"`
	FeedbackSummaryTable   FeedbackSummaryTable           `json:"feedback_summary_table"`
	ReviewCount            int                            `json:"review_count"`
	LastVisiblePage        int                            `json:"last_visible_page"`
	ParseStatusCode        int                            `json:"parse_status_code"`
}

type Result struct {
	Paid                   []Paid                   `json:"paid"`
	Filters                []Filters                `json:"filters"`
	Organic                []Organic                `json:"organic"`
	SearchInformation      SearchInformation        `json:"search_information"`
	Suggested              []SuggestedAmazonSearch  `json:"suggested"`
	AmazonChoices          []AmazonChoices          `json:"amazon_choices"`
	InstantRecommendations []InstantRecommendations `json:"instant_recommendations"`
	Pos                    int                      `json:"pos"`
	Url                    string                   `json:"url"`
	Asin                   string                   `json:"asin"`
	Price                  float64                  `json:"price"`
	Title                  string                   `json:"title"`
	Rating                 float64                  `json:"rating"`
	Currency               string                   `json:"currency"`
	IsPrime                bool                     `json:"is_prime"`
	PriceStr               string                   `json:"price_str"`
	PriceUpper             float64                  `json:"price_upper"`
	RatingsCount           int                      `json:"ratings_count"`
}

type Paid struct {
	Pos                 int           `json:"pos"`
	Url                 string        `json:"url"`
	Desc                string        `json:"desc"`
	Title               string        `json:"title"`
	DataRw              string        `json:"data_rw"`
	DataPcu             []string      `json:"data_pcu"`
	Sitelinks           PaidSitelinks `json:"sitelinks"`
	UrlShown            string        `json:"url_shown"`
	Asin                string        `json:"asin"`
	Price               float64       `json:"price"`
	Rating              float64       `json:"rating"`
	RelPos              int           `json:"rel_pos"`
	Currency            string        `json:"currency"`
	UrlImage            string        `json:"url_image"`
	BestSeller          bool          `json:"best_seller"`
	PriceUpper          float64       `json:"price_upper"`
	IsSponsored         bool          `json:"is_sponsored"`
	Manufacturer        string        `json:"manufacturer"`
	PricingCount        int           `json:"pricing_count"`
	ReviewsCount        int           `json:"reviews_count"`
	IsAmazonsChoice     bool          `json:"is_amazons_choice"`
	NoPriceReason       string        `json:"no_price_reason"`
	SalesVolume         string        `json:"sales_volume"`
	IsPrime             bool          `json:"is_prime"`
	ShippingInformation string        `json:"shipping_information"`
	PosOverall          int           `json:"pos_overall"`
}

type PaidSitelinks struct {
	Expanded []Expanded `json:"expanded,omitempty"`
	Inline   []Inline   `json:"inline,omitempty"`
}

type Expanded struct {
	Url   string `json:"url"`
	Desc  string `json:"desc"`
	Title string `json:"title"`
}

type Inline struct {
	Url   string `json:"url"`
	Desc  string `json:"desc"`
	Title string `json:"title"`
}

type Filters struct {
	Name   string `json:"name"`
	Values []struct {
		Url   string `json:"url"`
		Value string `json:"value"`
	} `json:"values"`
}

type Organic struct {
	Pos      int     `json:"pos"`
	Url      string  `json:"url"`
	Type     string  `json:"type"`
	Price    float64 `json:"price"`
	Title    string  `json:"title"`
	Currency string  `json:"currency"`
	Merchant struct {
		Url  string `json:"url"`
		Name string `json:"name"`
	} `json:"merchant"`
	PriceStr        string       `json:"price_str"`
	ProductId       string       `json:"product_id"`
	Asin            string       `json:"asin"`
	Rating          float64      `json:"rating"`
	UrlImage        string       `json:"url_image"`
	BestSeller      bool         `json:"best_seller"`
	PriceUpper      float64      `json:"price_upper"`
	IsSponsored     bool         `json:"is_sponsored"`
	Manufacturer    string       `json:"manufacturer"`
	PricingCount    int          `json:"pricing_count"`
	ReviewsCount    int          `json:"reviews_count"`
	IsAmazonsChoice bool         `json:"is_amazons_choice"`
	NoPriceReason   string       `json:"no_price_reason"`
	IsPrime         bool         `json:"is_prime"`
	SalesVolume     string       `json:"sales_volume"`
	Variations      []Variations `json:"variations"`
	PosOVerall      int          `json:"pos_overall"`
}

type Variations struct {
	Asin               string  `json:"asin"`
	Title              string  `json:"title"`
	Price              float64 `json:"price"`
	PriceStrikethrough float64 `json:"price_strikethrough"`
	NotAvailable       bool    `json:"not_available"`
}

type SearchInformation struct {
	Query             string `json:"query"`
	ShowingResultsFor string `json:"showing_results_for"`
}

type Variants struct {
	Type  string `json:"type"`
	Items []struct {
		Value     string `json:"value"`
		Selected  bool   `json:"selected"`
		Available bool   `json:"available"`
		ProductId string `json:"product_id"`
	} `json:"items"`
}

type RelatedItems struct {
	Items []struct {
		Url          string  `json:"url"`
		Price        float64 `json:"price"`
		Title        string  `json:"title"`
		Rating       float64 `json:"rating"`
		Currency     string  `json:"currency"`
		ReviewsCount int     `json:"reviews_count"`
	} `json:"items"`
}

type Specifications struct {
	Items []struct {
		Title string `json:"title"`
		Value string `json:"value"`
	} `json:"items"`
	SectionTitle string `json:"section_title"`
}

type Pricing struct {
	Price         float64 `json:"price"`
	Seller        string  `json:"seller"`
	Details       string  `json:"details"`
	Currency      string  `json:"currency"`
	Condition     string  `json:"condition"`
	PriceTax      float64 `json:"price_tax"`
	PriceTotal    float64 `json:"price_total"`
	SellerLink    string  `json:"seller_link"`
	PriceShipping float64 `json:"price_shipping"`

	Delivery        string      `json:"delivery"`
	SellerId        string      `json:"seller_id"`
	RatingCount     int         `json:"rating_count"`
	DeliveryOptions interface{} `json:"delivery_options"`
}

type SuggestedAmazonSearch struct {
	Url                 string  `json:"url"`
	Asin                string  `json:"asin"`
	Price               float64 `json:"price"`
	Title               string  `json:"title"`
	Rating              float64 `json:"rating"`
	Currency            string  `json:"currency"`
	UrlImage            string  `json:"url_image"`
	BestSeller          bool    `json:"best_seller"`
	PriceUpper          float64 `json:"price_upper"`
	IsSponsored         bool    `json:"is_sponsored"`
	Manufacturer        string  `json:"manufacturer"`
	PricingCount        int     `json:"pricing_count"`
	ReviewsCount        int     `json:"reviews_count"`
	IsAmazonsChoice     bool    `json:"is_amazons_choice"`
	Pos                 int     `json:"pos"`
	ShippingInformation string  `json:"shipping_information"`
	SalesVolume         string  `json:"sales_volume"`
	NoPriceReason       string  `json:"no_price_reason"`
	SuggestedQuery      string  `json:"suggested_query"`
}

type AmazonChoices struct {
	Url                 string       `json:"url"`
	Asin                string       `json:"asin"`
	Price               float64      `json:"price"`
	Title               string       `json:"title"`
	Rating              float64      `json:"rating"`
	Currency            string       `json:"currency"`
	UrlImage            string       `json:"url_image"`
	BestSeller          bool         `json:"best_seller"`
	PriceUpper          float64      `json:"price_upper"`
	IsSponsored         bool         `json:"is_sponsored"`
	Manufacturer        string       `json:"manufacturer"`
	PricingCount        int          `json:"pricing_count"`
	ReviewsCount        int          `json:"reviews_count"`
	IsAmazonsChoice     bool         `json:"is_amazons_choice"`
	Pos                 int          `json:"pos"`
	IsPrime             bool         `json:"is_prime"`
	ShippingInformation string       `json:"shipping_information"`
	SalesVolume         string       `json:"sales_volume"`
	NoPriceReason       string       `json:"no_price_reason"`
	Variations          []Variations `json:"variations"`
}

type InstantRecommendations struct {
	Url             string  `json:"url"`
	Asin            string  `json:"asin"`
	Price           float64 `json:"price"`
	Title           string  `json:"title"`
	Rating          float64 `json:"rating"`
	Currency        string  `json:"currency"`
	UrlImage        string  `json:"url_image"`
	BestSeller      bool    `json:"best_seller"`
	PriceUpper      float64 `json:"price_upper"`
	IsSponsored     bool    `json:"is_sponsored"`
	Manufacturer    string  `json:"manufacturer"`
	PricingCount    int     `json:"pricing_count"`
	ReviewsCount    int     `json:"reviews_count"`
	IsAmazonsChoice bool    `json:"is_amazons_choice"`
	Pos             int     `json:"pos"`
	SalesVolume     string  `json:"sales_volume"`
	NoPriceReason   string  `json:"no_price_reason"`
}

type AmazonProductAds struct {
	Pos             int      `json:"pos"`
	Asin            string   `json:"asin"`
	Type            string   `json:"type"`
	Price           float64  `json:"price"`
	Title           string   `json:"title"`
	Images          []string `json:"images"`
	Rating          float64  `json:"rating"`
	Location        string   `json:"location"`
	PriceUpper      float64  `json:"price_upper"`
	ReviewsCount    int      `json:"reviews_count"`
	IsPrimeEligible bool     `json:"is_prime_eligible"`
}

type AmazonProductCategory struct {
	Ladder []struct {
		Url  string `json:"url"`
		Name string `json:"name"`
	} `json:"ladder"`
}

type AmazonProductDelivery struct {
	Date struct {
		By   string `json:"by"`
		From string `json:"from"`
	} `json:"date"`
	Type string `json:"type"`
}

type AmazonProductSalesRank struct {
	Rank   int `json:"rank"`
	Ladder []struct {
		Url  string `json:"url"`
		Name string `json:"name"`
	} `json:"ladder"`
}

type ProductDetails struct {
	Asin                         string `json:"asin"`
	Batteries                    string `json:"batteries"`
	ItemWeight                   string `json:"item_weight"`
	Manufacturer                 string `json:"manufacturer"`
	CustomerReviews              string `json:"customer_reviews"`
	BestSellersRank              string `json:"best_sellers_rank"`
	CountryOfOrigin              string `json:"country_of_origin"`
	ItemModelNumber              string `json:"item_model_number"`
	ProductDimensions            string `json:"product_dimensions"`
	DateFirstAvailable           string `json:"date_first_available"`
	IsDiscontinuedByManufacturer string `json:"is_discontinued_by_manufacturer"`
}

type AmazonRefurbishedProduct struct {
	Link struct {
		Url   string `json:"url"`
		Title string `json:"title"`
	} `json:"link"`
	ConditionTitle string `json:"condition_title"`
}

type AmazonRatingStarDistribution struct {
	Rating     float64 `json:"rating"`
	Percentage int     `json:"percentage"`
}

type AmazonReviews struct {
	Id                string  `json:"id"`
	Title             string  `json:"title"`
	Author            string  `json:"author"`
	Rating            float64 `json:"rating"`
	Content           string  `json:"content"`
	Timestamp         string  `json:"timestamp"`
	IsVerified        bool    `json:"is_verified"`
	ProductAttributes string  `json:"product_attributes"`
}

type AmazonQuestions struct {
	Title   string `json:"title"`
	Votes   int    `json:"votes"`
	Answers []struct {
		Author    string `json:"author"`
		Content   string `json:"content"`
		Timestamp string `json:"timestamp"`
	} `json:"answers"`
}

type RecentFeedback struct {
	Feedback    string `json:"feedback"`
	RatedBy     string `json:"rated_by"`
	RatingStars int    `json:"rating_stars"`
}

type FeedbackSummaryTable struct {
	Counts struct {
		ThirtyDays   int `json:"30_days"`
		NinetyDays   int `json:"90_days"`
		AllTime      int `json:"all_time"`
		TwelveMonths int `json:"12_months"`
	} `json:"counts"`
	Neutral struct {
		ThirtyDays   string `json:"30_days"`
		NinetyDays   string `json:"90_days"`
		AllTime      string `json:"all_time"`
		TwelveMonths string `json:"12_months"`
	} `json:"neutral"`
	Negative struct {
		ThirtyDays   string `json:"30_days"`
		NinetyDays   string `json:"90_days"`
		AllTime      string `json:"all_time"`
		TwelveMonths string `json:"12_months"`
	} `json:"negative"`
	Positive struct {
		ThirtyDays   string `json:"30_days"`
		NinetyDays   string `json:"90_days"`
		AllTime      string `json:"all_time"`
		TwelveMonths string `json:"12_months"`
	} `json:"positive"`
}

type Job struct {
	CallbackUrl string `json:"callback_url"`
	ClientID    int    `json:"client_id"`
	Context     []struct {
		Key   string      `json:"key"`
		Value interface{} `json:"value"`
	} `json:"context,omitempty"`
	CreatedAt           string        `json:"created_at"`
	Domain              string        `json:"domain"`
	GeoLocation         interface{}   `json:"geo_location"`
	ID                  string        `json:"id"`
	Limit               int           `json:"limit"`
	Locale              interface{}   `json:"locale"`
	Pages               int           `json:"pages"`
	Parse               bool          `json:"parse"`
	ParserType          interface{}   `json:"parser_type"`
	ParsingInstructions interface{}   `json:"parsing_instructions"`
	BrowserInstructions interface{}   `json:"browser_instructions"`
	Render              interface{}   `json:"render"`
	Url                 interface{}   `json:"url"`
	Query               string        `json:"query"`
	Source              string        `json:"source"`
	StartPage           int           `json:"start_page"`
	Status              string        `json:"status"`
	StorageType         interface{}   `json:"storage_type"`
	StorageUrl          interface{}   `json:"storage_url"`
	Subdomain           string        `json:"subdomain"`
	ContentEncoding     string        `json:"content_encoding"`
	UpdatedAt           string        `json:"updated_at"`
	UserAgentType       string        `json:"user_agent_type"`
	SessionInfo         interface{}   `json:"session_info"`
	Statuses            []interface{} `json:"statuses"`
	ClientNotes         interface{}   `json:"client_notes"`
	Links               []struct {
		Rel    string `json:"rel"`
		Href   string `json:"href"`
		Method string `json:"method"`
	} `json:"_links,omitempty"`
}

// Custom function to unmarshal into the Resp struct.
// Because of different return types depending on the parse option.
func (r *Resp) UnmarshalJSON(data []byte) error {
	// Unmarshal json data into RawResp map.
	var rawResp map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawResp); err != nil {
		return err
	}

	// Unmarshal the results array.
	if resultsData, ok := rawResp["results"]; ok {
		// Slice to store raw JSON messages for each result.
		var resultsRawMessages []json.RawMessage
		if err := json.Unmarshal(resultsData, &resultsRawMessages); err != nil {
			return err
		}

		// Unmarshal each result into the Results slice.
		for _, resultRawMessage := range resultsRawMessages {
			if r.Parse && !r.ParseInstructions {
				var result struct {
					ContentParsed Content `json:"content"`
					CreatedAt     string  `json:"created_at"`
					UpdatedAt     string  `json:"updated_at"`
					Page          int     `json:"page"`
					Url           string  `json:"url"`
					JobID         string  `json:"job_id"`
					StatusCode    int     `json:"status_code"`
				}
				if err := json.Unmarshal(resultRawMessage, &result); err != nil {
					return err
				}
				r.Results = append(r.Results, Results{
					ContentParsed: result.ContentParsed,
					CreatedAt:     result.CreatedAt,
					UpdatedAt:     result.UpdatedAt,
					Page:          result.Page,
					Url:           result.Url,
					JobID:         result.JobID,
					StatusCode:    result.StatusCode,
				})
			} else if r.Parse && r.ParseInstructions {
				var result struct {
					CustomContentParsed map[string]interface{} `json:"content"`
					CreatedAt           string                 `json:"created_at"`
					UpdatedAt           string                 `json:"updated_at"`
					Page                int                    `json:"page"`
					Url                 string                 `json:"url"`
					JobID               string                 `json:"job_id"`
					StatusCode          int                    `json:"status_code"`
				}
				if err := json.Unmarshal(resultRawMessage, &result); err != nil {
					return err
				}
				r.Results = append(r.Results, Results{
					CustomContentParsed: result.CustomContentParsed,
					CreatedAt:           result.CreatedAt,
					UpdatedAt:           result.UpdatedAt,
					Page:                result.Page,
					Url:                 result.Url,
					JobID:               result.JobID,
					StatusCode:          result.StatusCode,
				})
			} else if !r.Parse {
				var result struct {
					Content    string `json:"content"`
					CreatedAt  string `json:"created_at"`
					UpdatedAt  string `json:"updated_at"`
					Page       int    `json:"page"`
					Url        string `json:"url"`
					JobID      string `json:"job_id"`
					StatusCode int    `json:"status_code"`
				}
				if err := json.Unmarshal(resultRawMessage, &result); err != nil {
					return err
				}
				r.Results = append(r.Results, Results{
					Content:    result.Content,
					CreatedAt:  result.CreatedAt,
					UpdatedAt:  result.UpdatedAt,
					Page:       result.Page,
					Url:        result.Url,
					JobID:      result.JobID,
					StatusCode: result.StatusCode,
				})
			}
		}
	}

	// Unmarshal the job object.
	if jobData, ok := rawResp["job"]; ok {
		var job Job
		if err := json.Unmarshal(jobData, &job); err != nil {
			return err
		}
		r.Job = job
	}

	return nil
}

// GetResp returns a Resp struct from the http.Response object.
// It will use the parse and customParserFlag parameters
// to determine how to parse the response.
func GetResp(
	httpResp *http.Response,
	parse bool,
	customParserFlag bool,
) (*Resp, error) {
	// Read the resp body into a buffer.
	respBody, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return nil, err
	}

	// If status code not 200, return error.
	if httpResp.StatusCode != 200 {
		return nil, fmt.Errorf("error with status code %s: %s", httpResp.Status, respBody)
	}

	// Unmarshal the JSON object.
	res := &Resp{}
	res.Parse = parse
	res.ParseInstructions = customParserFlag
	if err := res.UnmarshalJSON(respBody); err != nil {
		return nil, fmt.Errorf("failed to parse JSON object: %v", err)
	}

	// Set status code and status.
	res.StatusCode = httpResp.StatusCode
	res.Status = httpResp.Status

	return res, nil
}
