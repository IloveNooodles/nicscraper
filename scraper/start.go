package scraper

import (
	"fmt"
	"log"
)

func (s Scraper) Start(prefixes []string, years []string) {
	for _, prefix := range prefixes {
		for _, year := range years {
			for i := 1; i <= s.Limit; i++ {
				nim := fmt.Sprintf("%s%s%03d", prefix, year, i)
				go func(nim string) {
					student, err := s.GetByNIM(nim)
					if err != nil {
						log.Printf("Failed to fetch %s, reason: %s", nim, err)
						s.Failed <- nim
						return
					}
					s.Students <- student
					log.Printf("Scraped %s", nim)
				}(nim)
			}
		}
	}
}
