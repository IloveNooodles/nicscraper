package scraper

import (
	"github.com/mkamadeus/nicscraper/models"
)

func (s Scraper) Receive() ([]models.Student, []string) {
	result := make([]models.Student, 0)
	failed := make([]string, 0)

	for i := 0; i < s.Limit; {
		select {
		case student := <-s.Students:
			result = append(result, *student)
		case err := <-s.Failed:
			failed = append(failed, err)
		}
		i++
	}

	return result, failed
}
