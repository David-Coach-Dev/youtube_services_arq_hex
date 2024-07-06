package models

type DataYT struct {
	Kind           string      `json:"kind"`
	Etag           string      `json:"etag"`
	NextPageToken  string      `json:"nextPageToken"`
	PrevPageToken  string      `json:"prevPageToken"`
	PageInfo       []PageInfo  `json:"pageInfo"`
	Items          []Items     `json:"items"`
}
type PageInfo struct {
	TotalResults    int  `json:"totalResults"`
	ResultsPerPage  int  `json:"resultsPerPage"`
}
type Items struct {
	ID           string  `json:"id"`
	Title        string  `json:"title"`
	Description  string  `json:"description"`
	Channel      string  `json:"channel"`
	Live         string  `json:"live"`
	URL          string  `json:"url"`
	Thumbnails   string  `json:"thumbnails"`
}
