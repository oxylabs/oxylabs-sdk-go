package serp

type Response struct {
	Results    []Result `json:"results"`
	Job        Job      `json:"job"`
	StatusCode int      `json:"status_code"`
	Status     string   `json:"status"`
}

type Result struct {
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
	Value string `json:"value"`
}

type Link struct {
	Rel    string `json:"rel"`
	Href   string `json:"href"`
	Method string `json:"method"`
}
