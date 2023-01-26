package service

import (
	"encoding/xml"
	"github.com/haxqer/vast"
	"golang_vast/internal/model"
	"golang_vast/internal/repository"
	"log"
	"os"
	"time"
)

const (
	VastXML = "./vast.xml"
)

type VastService struct {
	repos repository.VastRepository
}

func NewVastService(repo repository.VastRepository) *VastService {
	return &VastService{
		repos: repo,
	}
}

func (s *VastService) Generate_Vast(ad model.VastModel) error {

	vastModel, err := s.repos.AdVast(ad)
	if err != nil {
		return err
	}

	if vastModel.IsSkipable == true {
		d := vast.Duration(time.Duration(vastModel.SkipTime) * time.Second)

		v := vast.VAST{
			Version: "3.0",
			Ads: []vast.Ad{
				{
					ID: vastModel.ID,
					InLine: &vast.InLine{
						AdSystem: &vast.AdSystem{Name: "DSP"},
						AdTitle:  vast.CDATAString{CDATA: vastModel.Title},
						Impressions: []vast.Impression{
							{ID: "1", URI: ""},
						},

						Creatives: []vast.Creative{

							{
								ID:       "1",
								Sequence: 0,
								Linear: &vast.Linear{
									SkipOffset: &vast.Offset{Duration: &d},
									Duration:   vast.Duration(time.Duration(vastModel.AdsDuration) * time.Second),
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
											ID:       "1",
											Delivery: "progressive",
											Type:     "video/mp4",
											Width:    vastModel.AdsWidth,
											Height:   vastModel.AdsHeight,
											URI:      vastModel.MediaURI,
										},
									},
								},
							},
						},
					},
				},
			},
		}

		file, err := os.Create(VastXML)
		if err != nil {
			log.Fatalln(err)
			return err
		}

		defer file.Close()
		xmlFile := xml.NewEncoder(file)

		if err := xmlFile.Encode(v); err != nil {
			log.Fatalln(err)
			return err
		}
	} else {
		v := vast.VAST{
			Version: "3.0",
			Ads: []vast.Ad{
				{
					ID: vastModel.ID,
					InLine: &vast.InLine{
						AdSystem: &vast.AdSystem{Name: "DSP"},
						AdTitle:  vast.CDATAString{CDATA: vastModel.Title},
						Impressions: []vast.Impression{
							{ID: "1", URI: ""},
						},

						Creatives: []vast.Creative{

							{
								ID:       "1",
								Sequence: 0,
								Linear: &vast.Linear{
									Duration: vast.Duration(time.Duration(vastModel.AdsDuration) * time.Second),
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
											ID:       "1",
											Delivery: "progressive",
											Type:     "video/mp4",
											Width:    vastModel.AdsWidth,
											Height:   vastModel.AdsHeight,
											URI:      vastModel.MediaURI,
										},
									},
								},
							},
						},
					},
				},
			},
		}

		file, err := os.Create(VastXML)
		if err != nil {
			log.Fatalln(err)
			return err
		}

		defer file.Close()
		xmlFile := xml.NewEncoder(file)

		if err := xmlFile.Encode(v); err != nil {
			log.Fatalln(err)
			return err
		}
	}

	return nil
}
