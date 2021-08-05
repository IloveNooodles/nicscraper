package scraper

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

func (s Scraper) Start() {
	for _, prefix := range s.Args.Prefixes {
		for _, year := range s.Args.Years {
			for i := 1; i <= s.Args.Limit; i++ {
				nim := fmt.Sprintf("%s%s%03d", prefix, year, i)
				go func(nim string) {
					student, err := s.GetByNIM(nim)
					logrus.Debugf("Student: %s Error: %s", student, err)
					if err != nil {
						logrus.Warnf("Failed to fetch %s, reason: %s", nim, err)
						s.Failed <- nim
						return
					}
					s.Students <- *student
					logrus.Infof("Scraped %s", nim)
				}(nim)
			}
		}
	}
}
