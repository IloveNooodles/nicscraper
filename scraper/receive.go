package scraper

import "github.com/mkamadeus/nicscraper/models"

func (s Scraper) Receive(limit int) []models.Student {
	result := make([]models.Student, 0)
	for i := 0; i < limit; {
		s := <-s.Students
		result = append(result, *s)
		i++
	}
	return result
}
