package service

import (
	"encoding/xml"
	"fmt"
	"github.com/haxqer/vast"
	"golang_vast/internal/repository"
	"golang_vast/internal/service/helpers"
	"log"
	"os"
)

const (
	Xml = ".xml"
)

type VastService struct {
	repos *repository.VastPostgres
}

func NewVastService(repo *repository.VastPostgres) *VastService {
	return &VastService{
		repos: repo,
	}
}

func (s *VastService) GenerateVast(id int) error {
	// Get ad from db
	ad, err := s.repos.GetAd(id)
	if err != nil {
		return err
	}

	var v vast.VAST
	fmt.Println(ad.IsSkippable)

	if ad.IsSkippable == true {
		fmt.Println("true")
		v = helpers.NewSkippableVastStructure(ad)
	} else {
		fmt.Println("false")
		v = helpers.NewUnskippableVastStructure(ad)
	}

	filename := fmt.Sprintf("%d%s", ad.Id, Xml)

	file, ErrOs := os.Create(filename)
	if ErrOs != nil {
		log.Fatalln(ErrOs)
		return ErrOs
	}

	defer file.Close()

	xmlFile := xml.NewEncoder(file)

	if err = xmlFile.Encode(v); err != nil {
		log.Fatalln(err)
		return err
	}

	return nil
}
