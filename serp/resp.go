package serp

import (
	"encoding/json"
)

// Custom function to unmarshal into the Response struct.
// Because of different return types depending on the parse option.
func (r *Response) UnmarshalJSON(data []byte) error {
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
					URL           string  `json:"url"`
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
					URL:           result.URL,
					JobID:         result.JobID,
					StatusCode:    result.StatusCode,
				})
			} else if r.Parse && r.ParseInstructions {
				var result struct {
					CustomContentParsed map[string]interface{} `json:"content"`
					CreatedAt           string                 `json:"created_at"`
					UpdatedAt           string                 `json:"updated_at"`
					Page                int                    `json:"page"`
					URL                 string                 `json:"url"`
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
					URL:                 result.URL,
					JobID:               result.JobID,
					StatusCode:          result.StatusCode,
				})
			} else if !r.Parse {
				var result struct {
					Content    string `json:"content"`
					CreatedAt  string `json:"created_at"`
					UpdatedAt  string `json:"updated_at"`
					Page       int    `json:"page"`
					URL        string `json:"url"`
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
					URL:        result.URL,
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

type Response struct {
	Parse             bool     `json:"parse"`
	ParseInstructions bool     `json:"parse_instructions"`
	Results           []Result `json:"results"`
	Job               Job      `json:"job"`
	StatusCode        int      `json:"status_code"`
	Status            string   `json:"status"`
}

type ResponseProxy struct {
	ContentParsed Content
	Content       string
	StatusCode    int
	Status        string
}

type Job struct {
	CallbackURL         string        `json:"callback_url"`
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
	URL                 interface{}   `json:"url"`
	Query               string        `json:"query"`
	Source              string        `json:"source"`
	StartPage           int           `json:"start_page"`
	Status              string        `json:"status"`
	StorageType         interface{}   `json:"storage_type"`
	StorageURL          interface{}   `json:"storage_url"`
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
	URL                 string `json:"url"`
	JobID               string `json:"job_id"`
	StatusCode          int    `json:"status_code"`
	ParserType          string `json:"parser_type"`
}

type Content struct {
	URL             string  `json:"url"`
	Page            int     `json:"page"`
	Results         Results `json:"results"`
	LastVisiblePage int     `json:"last_visible_page"`
	ParseStatusCode int     `json:"parse_status_code"`
}

type Results struct {
	Pla               Pla               `json:"pla"`
	Paid              []Paid            `json:"paid"`
	Organic           []Organic         `json:"organic"`
	Knowledge         Knowledge         `json:"knowledge"`
	LocalPack         LocalPack         `json:"local_pack"`
	InstantAnswers    []InstantAnswer   `json:"instant_answers"`
	RelatedSearches   RelatedSearches   `json:"related_searches"`
	SearchInformation SearchInformation `json:"search_information"`
	TotalResultsCount int               `json:"total_results_count"`
}

type Pla struct {
	Items      []PlaItem `json:"items"`
	PosOverall int       `json:"pos_overall,omitempty"`
}

type PlaItem struct {
	Pos       int    `json:"pos"`
	URL       string `json:"url"`
	Price     string `json:"price"`
	Title     string `json:"title"`
	Seller    string `json:"seller"`
	URLImage  string `json:"url_image"`
	ImageData string `json:"image_data"`
}

type InstantAnswer struct {
	Type       string `json:"type"`
	Parsed     bool   `json:"_parsed"`
	PosOverall int    `json:"pos_overall"`
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
	URL   string `json:"url"`
	Title string `json:"title"`
}

type RelatedSearch struct {
	URL          string `json:"url"`
	Title        string `json:"title"`
	SectionTitle string `json:"section_title"`
}

type LocalPack struct {
	Items      []Item `json:"items"`
	PosOverall int    `json:"pos_overall"`
}

type Item struct {
	Cid         string        `json:"cid"`
	Pos         int           `json:"pos"`
	Links       []LinkElement `json:"links"`
	Title       string        `json:"title"`
	Rating      int           `json:"rating"`
	Address     string        `json:"address"`
	RatingCount int           `json:"rating_count"`
}

type Organic struct {
	Pos        int              `json:"pos"`
	URL        string           `json:"url"`
	Desc       string           `json:"desc"`
	Title      string           `json:"title"`
	Sitelinks  OrganicSitelinks `json:"sitelinks,omitempty"`
	URLShown   string           `json:"url_shown"`
	PosOverall int              `json:"pos_overall"`
}

type OrganicSitelinks struct {
	Expanded []Profile `json:"expanded,omitempty"`
	Inline   []Profile `json:"inline,omitempty"`
}

type Paid struct {
	Pos        int           `json:"pos"`
	URL        string        `json:"url"`
	Desc       string        `json:"desc"`
	Title      string        `json:"title"`
	DataRw     string        `json:"data_rw"`
	DataPcu    []string      `json:"data_pcu"`
	Sitelinks  PaidSitelinks `json:"sitelinks,omitempty"`
	URLShown   string        `json:"url_shown"`
	PosOverall int           `json:"pos_overall"`
}

type PaidSitelinks struct {
	Expanded []Expanded `json:"expanded,omitempty"`
	Inline   []Profile  `json:"inline,omitempty"`
}

type Expanded struct {
	URL   string `json:"url"`
	Desc  string `json:"desc"`
	Title string `json:"title"`
}

type RelatedSearches struct {
	PosOverall      int      `json:"pos_overall"`
	RelatedSearches []string `json:"related_searches"`
}

type SearchInformation struct {
	Query             string `json:"query"`
	ShowingResultsFor string `json:"showing_results_for"`
	TotalResultsCount int    `json:"total_results_count"`
}
