package scraper

import (
	"sort"

	"github.com/mkamadeus/nicscraper/models"
	"github.com/sirupsen/logrus"
)

func (s Scraper) Receive() ([]models.Student, []string) {
	result := make([]models.Student, 0)
	failed := make([]string, 0)

	logrus.Debug(len(s.Args.Prefixes) * len(s.Args.Years) * s.Args.Limit)

	for i := 0; i < len(s.Args.Prefixes)*len(s.Args.Years)*s.Args.Limit; {
		select {
		case student := <-s.Students:
			logrus.Debugf("Received student: %s", student)
			result = append(result, student)
		case err := <-s.Failed:
			logrus.Debugf("Received failed: %s", err)
			failed = append(failed, err)
		}
		i++
	}

	// Sort output
	sort.Slice(result[:], func(i, j int) bool {
		return result[i].Username < result[j].Username
	})
	sort.Slice(failed[:], func(i, j int) bool {
		return failed[i] < failed[j]
	})

	return result, failed
}
