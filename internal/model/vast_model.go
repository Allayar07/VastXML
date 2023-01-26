package model

type VastModel struct {
	ID          string `json:"id"`
	IsSkipable  bool   `json:"is_Skipable"`
	SkipTime    int    `json:"skipTime"`
	AdsDuration int64  `json:"adsDuration"`
	Title       string `json:"title"`
	MediaURI    string `json:"mediaURI"`
}
