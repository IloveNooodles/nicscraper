package scraper

import (
	"errors"

	"github.com/mkamadeus/nicscraper/models"
)

type Scraper struct {
	IsVerbose bool
	Students  chan models.Student
	Failed    chan string
	Args      *models.Arguments
}

func New(args *models.Arguments) (*Scraper, error) {
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
