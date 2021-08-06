package scraper

import (
	"fmt"
	"math"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

type Progress struct {
	m     sync.Mutex
	count int
	limit int
}

const CHUNK_SIZE int = 50

func (s Scraper) Start() {
	progress := Progress{count: 0, limit: s.Args.Limit * len(s.Args.Years) * len(s.Args.Prefixes)}
	for _, prefix := range s.Args.Prefixes {
		for _, year := range s.Args.Years {
			nimPrefix := fmt.Sprintf("%s%s", prefix, year)
			for i := 0; i <= s.Args.Limit/CHUNK_SIZE; i++ {
				go func(offset int, prefix string) {
					batasAtas := int(math.Min((float64(offset)+1)*float64(CHUNK_SIZE), float64(s.Args.Limit)))
					start := (offset * CHUNK_SIZE) + 1
					isSkipping := false
					notFoundStreak := 0
					if start <= batasAtas {
						for each := start; each <= batasAtas; each++ {
							nim := fmt.Sprintf("%s%03d", prefix, each)
							if isSkipping {
								s.Failed <- nim
								continue
							}

							student, err := s.GetByNIM(nim)
							logrus.Debugf("Student: %s Error: %s", student, err)
							if err != nil {
								logrus.Warnf("Failed to fetch %s, reason: %s", nim, err)
								s.Failed <- nim
								notFoundStreak++
								if notFoundStreak > 5 {
									isSkipping = true
									logrus.Warnf("Skipping %s - %s", nim, fmt.Sprintf("%s%03d", prefix, batasAtas))
								}
								continue
							}
							notFoundStreak = 0
							s.Students <- student
						}
						logrus.Infof("Fetched %s - %s", fmt.Sprintf("%s%03d", prefix, start), fmt.Sprintf("%s%03d", prefix, batasAtas))
						progress.m.Lock()
						progress.count += batasAtas - start + 1
						progress.m.Unlock()
					}
				}(i, nimPrefix)
			}
		}
	}
	go func() {
		lastProgress := float64(0)
		for {
			progress.m.Lock()
			currentProgress := float64(progress.count) * 100 / float64(progress.limit)
			if lastProgress != currentProgress {
				logrus.Infof("Progress %.2f%%", currentProgress)
				lastProgress = currentProgress
			}
			progress.m.Unlock()
			time.Sleep(3 * time.Second)
		}
	}()
}
