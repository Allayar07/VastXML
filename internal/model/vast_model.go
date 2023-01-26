package model

type VastModel struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	AdsDuration int64  `json:"adsDuration"`
	AdsHeight   int    `json:"adsHeight"`
	AdsWidth    int    `json:"adsWidth"`
	SkipTime    int    `json:"skipTime"`
	MediaURI    string `json:"mediaURI"`
	IsSkippable bool   `json:"isSkippable"`
}
