package model

var BaseURL string
var APIKey string

// Request struct for searching - https://www.giantbomb.com/api/documentation/#toc-0-41
type SearchGamesRequest struct {
	FieldList []string `json:"field_list,omitempty"`
	Limit     string   `json:"limit,omitempty"`
	Page      string   `json:"page,omitempty"`
	Query     string   `json:"query"`
	Resources string   `json:"resources"`
}

// Search games response struct
type SearchGameResponse struct {
    Results   []Game `json:"results"`
}

// Request struct for game
type GameRequest struct {
    FieldList []string `json:"field_list,omitempty"`
    GUID      string   `json:"guid"`
}

// Response struct for game
type GameResponse struct {
    Results   Game `json:"results"` 
}

// Request struct for DLCs
type DLCsRequest struct {
    FieldList []string `json:"field_list,omitempty"`
    Filter    string   `json:"filter"`
    Sort      string   `json:"sort"`
}

// Response struct for DLCs
type DLCsResponse struct {
    Results   []DLC `json:"Results"`
}

// platform struct - https://www.giantbomb.com/api/documentation/#toc-0-29
type Platform struct {
    GUID string `json:"guid"`
    Name string `json:"name"`
}

// dlc struct - https://www.giantbomb.com/api/documentation/#toc-0-12
type DLC struct {
    Deck            string   `json:"deck"`
    Game            Game     `json:"game"`
    GUID            string   `json:"guid"`
    ID              int      `json:"id"`
    Name            string   `json:"name"`
    Platform        Platform `json:"platform"`
    ReleaseDate     string   `json:"release_date"`
}

// game struct - https://www.giantbomb.com/api/documentation/#toc-0-16
type Game struct {
    Deck                   string    `json:"deck"`
    Description            string    `json:"description"`
    ExpectedReleaseYear    int       `json:"expected_release_year"`
	GUID                   string    `json:"guid"`
	ID                     int       `json:"id"`
    Name                   string    `json:"name"`
    DLCs                   []DLC     `json:"dlcs"`
}
