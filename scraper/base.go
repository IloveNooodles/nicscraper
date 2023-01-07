package scraper

import (
	"github.com/mkamadeus/nicscraper/models"
	"github.com/pkg/errors"
)

type Scraper struct {
	IsVerbose bool
	Students  chan models.Student
	Failed    chan string
	Args      models.Arguments
}

type TeamsScrapper struct {
	IsVerbose bool
	Students  chan models.TeamsStudent
	Failed    chan string
	Args      models.Arguments
}

func New(args models.Arguments) (*Scraper, error) {
	scraper := &Scraper{
		Students: make(chan models.Student),
		Failed:   make(chan string),
		Args:     args,
	}

	if !scraper.IsConnected() {
		return nil, errors.New("scraper not connected")
	}

	return scraper, nil
}

func NewTeams(args models.Arguments) (*TeamsScrapper, error) {
	scraper := &TeamsScrapper{
		Students: make(chan models.TeamsStudent),
		Failed:   make(chan string),
		Args:     args,
	}

	if scraper.Args.CVID == "" {
		return nil, errors.New("Please input Cvid Token")
	}

	if scraper.Args.JWT == "" {
		return nil, errors.New("Please input Jwt Token")
	}

	if !scraper.IsConnected() {
		return nil, errors.New("scraper not connected")
	}

	return scraper, nil
}
