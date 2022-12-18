package scraper

import (
	"github.com/mkamadeus/nicscraper/models"
	"github.com/pkg/errors"
)

type Scraper struct {
	IsVerbose    bool
	Students     chan models.Student
	TeamsStudent chan models.TeamsStudent
	Failed       chan string
	Args         models.Arguments
}

func New(args models.Arguments) (*Scraper, error) {
	scraper := &Scraper{
		Students:     make(chan models.Student),
		TeamsStudent: make(chan models.TeamsStudent),
		Failed:       make(chan string),
		Args:         args,
	}

	if scraper.Args.UseTeams {

		if scraper.Args.Cvid == "" {
			return nil, errors.New("Please input Cvid Token")
		}

		if scraper.Args.Jwt == "" {
			return nil, errors.New("Please input Jwt Token")
		}
	}

	if !scraper.IsConnected() {
		return nil, errors.New("scraper not connected")
	}

	return scraper, nil
}
