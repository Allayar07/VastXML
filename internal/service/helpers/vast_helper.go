package helpers

import (
	"github.com/haxqer/vast"
	"golang_vast/internal/model"
	"strconv"
	"time"
)

func NewSkippableVastStructure(ad model.VastModel) vast.VAST {
	id := strconv.Itoa(ad.Id)
	d := vast.Duration(time.Duration(ad.SkipTime) * time.Second)
	return vast.VAST{
		Version: "3.0",
		Ads: []vast.Ad{
			{
				ID: id,
				InLine: &vast.InLine{
					AdSystem: &vast.AdSystem{Name: "DSP"},
					AdTitle:  vast.CDATAString{CDATA: ad.Title},
					Impressions: []vast.Impression{
						{ID: id, URI: ""},
					},

					Creatives: []vast.Creative{

						{
							ID:       id,
							Sequence: 0,
							Linear: &vast.Linear{
								SkipOffset: &vast.Offset{Duration: &d},
								Duration:   vast.Duration(time.Duration(ad.AdsDuration) * time.Second),
								TrackingEvents: []vast.Tracking{
									{Event: vast.Event_type_start, URI: "http://82.180.154.172:8090/api/uploads/open/startevent.jpg"},
									{Event: vast.Event_type_skip, URI: "http://82.180.154.172:8090/api/uploads/open/skip.jpg"},
									{Event: vast.Event_type_firstQuartile, URI: "http://82.180.154.172:8090/api/uploads/open/25percent.jpg"},
									{Event: vast.Event_type_pause, URI: "http://82.180.154.172:8090/api/uploads/open/pause.jpg"},
									{Event: vast.Event_type_midpoint, URI: "http://82.180.154.172:8090/api/uploads/open/midpoint.jpg"},
									{Event: vast.Event_type_thirdQuartile, URI: "http://82.180.154.172:8090/api/uploads/open/75percent.jpg"},
									{Event: vast.Event_type_complete, URI: "http://82.180.154.172:8090/api/uploads/open/complete.jpg"},
								},

								MediaFiles: []vast.MediaFile{
									{
										ID:       id,
										Delivery: "progressive",
										Type:     "video/mp4",
										Width:    ad.AdsWidth,
										Height:   ad.AdsHeight,
										URI:      ad.MediaURI,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func NewUnskippableVastStructure(ad model.VastModel) vast.VAST {
	id := strconv.Itoa(ad.Id)
	return vast.VAST{
		Version: "3.0",
		Ads: []vast.Ad{
			{
				ID: id,
				InLine: &vast.InLine{
					AdSystem: &vast.AdSystem{Name: "DSP"},
					AdTitle:  vast.CDATAString{CDATA: ad.Title},
					Impressions: []vast.Impression{
						{ID: id, URI: ""},
					},

					Creatives: []vast.Creative{

						{
							ID:       id,
							Sequence: 0,
							Linear: &vast.Linear{
								Duration: vast.Duration(time.Duration(ad.AdsDuration) * time.Second),
								TrackingEvents: []vast.Tracking{
									{Event: vast.Event_type_start, URI: "http://82.180.154.172:8090/api/uploads/open/startevent.jpg"},
									{Event: vast.Event_type_skip, URI: "http://82.180.154.172:8090/api/uploads/open/skip.jpg"},
									{Event: vast.Event_type_firstQuartile, URI: "http://82.180.154.172:8090/api/uploads/open/25percent.jpg"},
									{Event: vast.Event_type_pause, URI: "http://82.180.154.172:8090/api/uploads/open/pause.jpg"},
									{Event: vast.Event_type_midpoint, URI: "http://82.180.154.172:8090/api/uploads/open/midpoint.jpg"},
									{Event: vast.Event_type_thirdQuartile, URI: "http://82.180.154.172:8090/api/uploads/open/75percent.jpg"},
									{Event: vast.Event_type_complete, URI: "http://82.180.154.172:8090/api/uploads/open/complete.jpg"},
								},

								MediaFiles: []vast.MediaFile{
									{
										ID:       id,
										Delivery: "progressive",
										Type:     "video/mp4",
										Width:    ad.AdsWidth,
										Height:   ad.AdsHeight,
										URI:      ad.MediaURI,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}
