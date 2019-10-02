package models

type SearchParams struct {
	Slug      string
	StartDate string
	Page      int
	Limit     int
	OrderBy   []string
}
