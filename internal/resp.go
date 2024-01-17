package internal

import (
	"encoding/json"
)

// Custom function to unmarshal into the Response struct.
// Because of different return types depending on the parse option.
func (r *Resp) UnmarshalJSON(data []byte) error {
	// Unmarshal json data into RawResponse map.
	var rawResponse map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawResponse); err != nil {
		return err
	}

	// Unmarshal the results array.
	if resultsData, ok := rawResponse["results"]; ok {
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
				r.Results = append(r.Results, Result{
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
				r.Results = append(r.Results, Result{
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
				r.Results = append(r.Results, Result{
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
	if jobData, ok := rawResponse["job"]; ok {
		var job Job
		if err := json.Unmarshal(jobData, &job); err != nil {
			return err
		}
		r.Job = job
	}

	return nil
}

type Resp struct {
	Parse             bool     `json:"parse"`
	ParseInstructions bool     `json:"parse_instructions"`
	Results           []Result `json:"results"`
	Job               Job      `json:"job"`
	StatusCode        int      `json:"status_code"`
	Status            string   `json:"status"`
}

type Job struct {
	CallbackUrl         string        `json:"callback_url"`
	ClientID            int           `json:"client_id"`
	Context             []Context     `json:"context"`
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
	Links               []Link        `json:"_links"`
}

type Context struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

type Link struct {
	Rel    string `json:"rel"`
	Href   string `json:"href"`
	Method string `json:"method"`
}

type Result struct {
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
	Url             string  `json:"url"`
	Page            int     `json:"page"`
	Results         Results `json:"results"`
	LastVisiblePage int     `json:"last_visible_page"`
	ParseStatusCode int     `json:"parse_status_code"`
}

type Results struct {
	Pla                        Pla                          `json:"pla"`
	Paid                       []Paid                       `json:"paid"`
	Images                     Image                        `json:"images"`
	Organic                    []Organic                    `json:"organic"`
	Twitter                    []Twitter                    `json:"twitter"`
	Knowledge                  Knowledge                    `json:"knowledge"`
	LocalPack                  LocalPack                    `json:"local_pack"`
	TopStories                 TopStory                     `json:"top_stories"`
	PopularProducts            PopularProducts              `json:"popular_products"`
	RelatedSearches            RelatedSearches              `json:"related_searches"`
	RelatedQuestions           RelatedQuestions             `json:"related_questions"`
	SearchInformation          SearchInformation            `json:"search_information"`
	ItemCarousel               ItemCarousel                 `json:"item_carousel"`
	Recipes                    Recipes                      `json:"recipes"`
	Videos                     Videos                       `json:"videos"`
	FeaturedSnippet            []FeaturedSnippet            `json:"featured_snippet"`
	RelatedSearchesCategorized []RelatedSearchesCategorized `json:"related_searches_categorized"`
	Hotels                     Hotels                       `json:"hotels"`
	Flights                    Flights                      `json:"flights"`
	VideoBox                   VideoBox                     `json:"video_box"`
	LocalServiceAds            LocalServiceAds              `json:"local_service_ads"`
	TotalResultsCount          int                          `json:"total_results_count"`
}

type Pla struct {
	Items      []PlaItem `json:"items"`
	PosOverall int       `json:"pos_overall,omitempty"`
}

type PlaItem struct {
	Pos       int    `json:"pos"`
	Url       string `json:"url"`
	Price     string `json:"price"`
	Title     string `json:"title"`
	Seller    string `json:"seller"`
	UrlImage  string `json:"url_image"`
	ImageData string `json:"image_data"`
}

type Paid struct {
	Pos        int           `json:"pos"`
	Url        string        `json:"url"`
	Desc       string        `json:"desc"`
	Title      string        `json:"title"`
	DataRw     string        `json:"data_rw"`
	DataPcu    []string      `json:"data_pcu"`
	Sitelinks  PaidSitelinks `json:"sitelinks"`
	UrlShown   string        `json:"url_shown"`
	PosOverall int           `json:"pos_overall"`
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

type Image struct {
	Items      []ImageItem `json:"items"`
	PosOverall int         `json:"pos_overall"`
}

type ImageItem struct {
	Alt string `json:"alt"`
	Pos int    `json:"pos"`
	Url string `json:"url"`
}

type Organic struct {
	Pos        int              `json:"pos"`
	Url        string           `json:"url"`
	Desc       string           `json:"desc"`
	Title      string           `json:"title"`
	Images     []string         `json:"images"`
	Sitelinks  OrganicSitelinks `json:"sitelinks,omitempty"`
	UrlShown   string           `json:"url_shown"`
	PosOverall int              `json:"pos_overall"`
}

type OrganicSitelinks struct {
	Expanded []Expanded `json:"expanded,omitempty"`
	Inline   []Inline   `json:"inline,omitempty"`
}

type Twitter struct {
	Pos        int           `json:"pos"`
	Url        string        `json:"url"`
	Items      []TwitterItem `json:"items"`
	Title      string        `json:"title"`
	PosOverall int           `json:"pos_overall"`
}

type TwitterItem struct {
	Pos       int    `json:"pos"`
	Url       string `json:"url"`
	Content   string `json:"content"`
	TimeFrame string `json:"time_frame"`
}

type Knowledge struct {
	Title           string          `json:"title"`
	Images          []string        `json:"images"`
	Factoids        []Factoid       `json:"factoids"`
	Profiles        []Profile       `json:"profiles"`
	Subtitle        string          `json:"subtitle"`
	Description     string          `json:"description"`
	RelatedSearches []RelatedSearch `json:"related_searches"`
}

type Factoid struct {
	Links   []LinkElement `json:"links"`
	Title   string        `json:"title"`
	Content string        `json:"content"`
}

type LinkElement struct {
	Href  string `json:"href"`
	Title string `json:"title"`
}

type Profile struct {
	Url   string `json:"url"`
	Title string `json:"title"`
}

type RelatedSearch struct {
	Url          string `json:"url"`
	Title        string `json:"title"`
	SectionTitle string `json:"section_title"`
}

type LocalPack struct {
	Items      []LocalPackItem `json:"items"`
	PosOverall int             `json:"pos_overall"`
}

type LocalPackItem struct {
	Cid         string `json:"cid"`
	Pos         int    `json:"pos"`
	Title       string `json:"title"`
	Rating      int    `json:"rating"`
	Address     string `json:"address"`
	Subtitle    string `json:"subtitle"`
	RatingCount int    `json:"rating_count"`
}

type TopStory struct {
	Items      []TopStoryItem `json:"items"`
	PosOverall int            `json:"pos_overall"`
}

type TopStoryItem struct {
	Pos       int    `json:"pos"`
	Url       string `json:"url"`
	Title     string `json:"title"`
	Source    string `json:"source"`
	TimeFrame string `json:"time_frame"`
}

type PopularProducts struct {
	Items      []PopularProductsItem `json:"items"`
	PosOverall int                   `json:"pos_overall"`
}

type PopularProductsItem struct {
	Pos       int    `json:"pos"`
	Price     string `json:"price"`
	Title     string `json:"title"`
	ImageData string `json:"image_data"`
}

type RelatedSearches struct {
	PosOverall      int      `json:"pos_overall"`
	RelatedSearches []string `json:"related_searches"`
}

type RelatedQuestions struct {
	Items      []RelatedQuestionsItem `json:"items"`
	PosOverall int                    `json:"pos_overall"`
}

type RelatedQuestionsItem struct {
	Pos      int    `json:"pos"`
	Answer   string `json:"answer"`
	Source   Source `json:"source"`
	Question string `json:"question"`
}

type Source struct {
	Url      string `json:"url"`
	Title    string `json:"title"`
	UrlShown string `json:"url_shown"`
}

type SearchInformation struct {
	Query             string `json:"query"`
	ShowingResultsFor string `json:"showing_results_for"`
	TotalResultsCount int    `json:"total_results_count"`
}

type ItemCarousel struct {
	Items      []ItemCarouselItem `json:"items"`
	Title      string             `json:"title"`
	PosOverall int                `json:"pos_overall"`
}

type ItemCarouselItem struct {
	Pos      int    `json:"pos"`
	Href     string `json:"href"`
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
}

type Recipes struct {
	Items      []RecipesItem `json:"items"`
	PosOverall int           `json:"pos_overall"`
}

type RecipesItem struct {
	Pos      int    `json:"pos"`
	Url      string `json:"url"`
	Desc     string `json:"desc"`
	Title    string `json:"title"`
	Rating   int    `json:"rating"`
	Source   string `json:"source"`
	Duration string `json:"duration"`
}

type Videos struct {
	Items      []VideosItem `json:"items"`
	PosOverall int          `json:"pos_overall"`
}

type VideosItem struct {
	Pos    int    `json:"pos"`
	Url    string `json:"url"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Source string `json:"source"`
}

type FeaturedSnippet struct {
	Url        string `json:"url"`
	Desc       string `json:"desc"`
	Title      string `json:"title"`
	UrlShown   string `json:"url_shown"`
	PosOverall int    `json:"pos_overall"`
}

type RelatedSearchesCategorized struct {
	Items      []RelatedSearchesCategorizedItem `json:"items"`
	Category   Category                         `json:"category"`
	PosOverall int                              `json:"pos_overall"`
}

type RelatedSearchesCategorizedItem struct {
	Url   string `json:"url"`
	Title string `json:"title"`
}

type Category struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type Hotels struct {
	DateTo  string `json:"date_to"`
	Results []struct {
		Price       string `json:"price"`
		Title       string `json:"title"`
		Description string `json:"description"`
	} `json:"results"`
	DateFrom   string `json:"date_from"`
	PosOverall int    `json:"pos_overall"`
}

type Flights struct {
	To      string `json:"to"`
	From    string `json:"from"`
	DateTo  string `json:"date_to"`
	Results []struct {
		Url      string `json:"url"`
		Type     string `json:"type"`
		Price    string `json:"price"`
		Airline  string `json:"airline"`
		Duration string `json:"duration"`
	} `json:"results"`
	DateFrom   string `json:"date_from"`
	PosOverall int    `json:"pos_overall"`
}

type VideoBox struct {
	Url        string `json:"url"`
	Title      string `json:"title"`
	PosOverall int    `json:"pos_overall"`
}

type LocalServiceAds struct {
	Items []struct {
		Pos              int    `json:"pos"`
		Url              string `json:"url"`
		Title            string `json:"title"`
		Rating           int    `json:"rating"`
		ReviewsCount     int    `json:"reviews_count"`
		GoogleGuaranteed bool   `json:"google_guaranteed"`
	} `json:"items"`
	PosOverall int `json:"pos_overall"`
}
