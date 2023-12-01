package serp

type ParseTrueResponse struct {
	Results    []ResultParseTrue `json:"results"`
	Job        Job               `json:"job"`
	StatusCode int               `json:"status_code"`
	Status     string            `json:"status"`
}

type ParseFalseResponse struct {
	Results    []ResultParseFalse `json:"results"`
	Job        Job                `json:"job"`
	StatusCode int                `json:"status_code"`
	Status     string             `json:"status"`
}

type ResultParseTrue struct {
	Content    Content `json:"content"`
	CreatedAt  string  `json:"created_at"`
	UpdatedAt  string  `json:"updated_at"`
	Page       int     `json:"page"`
	URL        string  `json:"url"`
	JobID      string  `json:"job_id"`
	StatusCode int     `json:"status_code"`
}

type Content struct {
	Url             string `json:"url"`
	Page            int    `json:"page"`
	Result          Result `json:"results"`
	LastVisiblePage int    `json:"last_visible_page"`
	ParseStatusCode int    `json:"parse_status_code"`
}

type Result struct {
	Pla               Pla                    `json:"pla"`
	Paid              []Paid                 `json:"paid"`
	Images            Images                 `json:"images"`
	Organic           []Organic              `json:"organic"`
	Knowledge         Knowledge              `json:"knowledge"`
	InstantAnswers    []InstantAnswers       `json:"instant_answers"`
	RelatedSearches   RelatedSearchesResults `json:"related_searches"`
	SearchInformation SearchInformation      `json:"search_information"`
	TotalResultsCount int                    `json:"total_results_count"`
	LastVisiblePage   int                    `json:"last_visible_page"`
	ParseStatusCode   int                    `json:"parse_status_code"`
}
type Pla struct {
	Items      []PlaItem `json:"items"`
	PosOverall *int64    `json:"pos_overall,omitempty"`
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

type Paid struct {
	Position    int       `json:"pos"`
	Url         string    `json:"url"`
	Description string    `json:"desc"`
	Title       string    `json:"title"`
	DataRw      string    `json:"data_rw"`
	DataPcu     []string  `json:"data_pcu"`
	SiteLinks   SiteLinks `json:"sitelinks"`
}

type SiteLinks struct {
	Expanded    []Expanded `json:"expanded"`
	UrlShown    string     `json:"url_shown"`
	PosOverrall int        `json:"pos_overall"`
}

type Expanded struct {
	Url         string `json:"url"`
	Description string `json:"desc"`
	Title       string `json:"title"`
}

type Images struct {
	Items       []Item `json:"items"`
	PosOverrall int    `json:"pos_overall"`
}

type Item struct {
	Alt    string `json:"alt"`
	Pos    int    `json:"pos"`
	Url    string `json:"url"`
	Source string `json:"source"`
}

type Organic struct {
	Position    int              `json:"pos"`
	Url         string           `json:"url"`
	Title       string           `json:"title"`
	Description string           `json:"desc"`
	SiteLinks   SiteLinksOrganic `json:"sitelinks"`
	UrlShown    string           `json:"url_shown"`
	PosOverrall int              `json:"pos_overall"`
}

type SiteLinksOrganic struct {
	Expanded []ExpandedOrganic `json:"expanded"`
}

type ExpandedOrganic struct {
	Url   string `json:"url"`
	Title string `json:"title"`
}

type Knowledge struct {
	Title           string            `json:"title"`
	Images          []string          `json:"images"`
	Factoids        []Factoids        `json:"factoids"`
	Profiles        []Profiles        `json:"profiles"`
	Subtitle        string            `json:"subtitle"`
	Description     string            `json:"description"`
	RelatedSearches []RelatedSearches `json:"related_searches"`
}

type Factoids struct {
	Links   []Links `json:"links"`
	Title   string  `json:"title"`
	Content string  `json:"content"`
}

type Links struct {
	Href  string `json:"href"`
	Title string `json:"title"`
}

type Profiles struct {
	Url   string `json:"url"`
	Title string `json:"title"`
}

type RelatedSearches struct {
	Url          string `json:"url"`
	Title        string `json:"title"`
	SectionTitle string `json:"section_title"`
}

type InstantAnswers struct {
	Type        string `json:"type"`
	Parsed      bool   `json:"_parsed"`
	PosOverrall int    `json:"pos_overall"`
}

type RelatedSearchesResults struct {
	PosOverrall     int      `json:"pos_overall"`
	RelatedSearches []string `json:"related_searches"`
}

type SearchInformation struct {
	Query             string `json:"query"`
	ShowingResultsFor string `json:"showing_results_for"`
	TotalResultsCount int    `json:"total_results_count"`
}

type ResultParseFalse struct {
	Content    string `json:"content"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
	Page       int    `json:"page"`
	URL        string `json:"url"`
	JobID      string `json:"job_id"`
	StatusCode int    `json:"status_code"`
}

type Job struct {
	CallbackURL         string    `json:"callback_url"`
	ClientID            int       `json:"client_id"`
	Context             []Context `json:"context"`
	CreatedAt           string    `json:"created_at"`
	Domain              string    `json:"domain"`
	GeoLocation         string    `json:"geo_location"`
	ID                  string    `json:"id"`
	Limit               int       `json:"limit"`
	Locale              string    `json:"locale"`
	Pages               int       `json:"pages"`
	Parse               bool      `json:"parse"`
	ParserType          string    `json:"parser_type"`
	ParsingInstructions string    `json:"parsing_instructions"`
	BrowserInstructions string    `json:"browser_instructions"`
	Render              string    `json:"render"`
	URL                 string    `json:"url"`
	Query               string    `json:"query"`
	Source              string    `json:"source"`
	StartPage           int       `json:"start_page"`
	Status              string    `json:"status"`
	StorageType         string    `json:"storage_type"`
	StorageURL          string    `json:"storage_url"`
	Subdomain           string    `json:"subdomain"`
	ContentEncoding     string    `json:"content_encoding"`
	UpdatedAt           string    `json:"updated_at"`
	UserAgentType       string    `json:"user_agent_type"`
	SessionInfo         string    `json:"session_info"`
	Statuses            []string  `json:"statuses"`
	ClientNotes         string    `json:"client_notes"`
	Links               []Link    `json:"_links"`
}

type Context struct {
	Key   string `json:"key"`
	Value any    `json:"value"`
}

type Link struct {
	Rel    string `json:"rel"`
	Href   string `json:"href"`
	Method string `json:"method"`
}
