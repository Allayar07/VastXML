package service

import (
	"encoding/xml"
	"github.com/akrfjmt/vmap"
	"github.com/rs/vast"

	"os"
	"time"
)

const VmapXml = "test_vmap.xml"

var fld = false
var tru = true

func (s *VastService) GenerateVMAP() error {

	d := vast.Duration(3 * time.Second)
	vmp := vmap.VMAP{
		Version: "1.0",

		AdBreaks: []vmap.AdBreak{
			{
				BreakID:    "preroll",
				TimeOffset: vmap.Offset{Duration: &d},
				BreakType:  "linear",
				AdSource: &vmap.AdSource{
					ID:               "preroll-ads",
					FollowRedirects:  &tru,
					AllowMultipleAds: &fld,
					AdTagURI: &vmap.AdTagURI{
						TemplateType: "vast3",
						URI:          "https://opencdn.b-cdn.net/pub/5.0/e-v-1/preroll.xml",
					},
				},
			},
		},
	}
	f, err := os.Create(VmapXml)
	if err != nil {
		return err
	}

	defer f.Close()

	Encoded := xml.NewEncoder(f)
	if err = Encoded.Encode(vmp); err != nil {
		return err
	}

	return nil

}
