package scraper

import "github.com/mkamadeus/nicscraper/models"

type Scraper struct {
	Token     string
	IsVerbose bool
	Students  chan *models.Student
	Failed    chan string
}

func New(token string, verbose bool) *Scraper {
	return &Scraper{
		Token:     token,
		IsVerbose: verbose,
		Students:  make(chan *models.Student),
		Failed:    make(chan string),
	}
}
