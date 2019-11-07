package models

type SearchParams struct {
	Slug      string
	StartDate string
	Page      int
	Limit     int
	OrderBy   []string
}

type SearchResponse struct {
	Results []SiteSearchResult `json:"results"`
	Action  map[string]string  `json:"action"`
}

//SiteSearchResult is for full text search result
type SiteSearchResult struct {
	Type  string `json:"type"`
	Slug  string `json:"slug"`
	Title string `gorm:"type:tsvector" json:"title"`
}

func (s SiteSearchResult) TableName() string {
	return "searches"
}
