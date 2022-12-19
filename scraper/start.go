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

const CHUNK_SIZE int = 20
const SKIP_THRESHOLD int = 5

func (s Scraper) Start() {
	progress := Progress{count: 0, limit: s.Args.Limit * len(s.Args.Years.Arr) * len(s.Args.Prefixes.Arr)}

	for _, prefix := range s.Args.Prefixes.Arr {
		for _, year := range s.Args.Years.Arr {

			nimPrefix := fmt.Sprintf("%s%s", prefix, year)

			// Split goroutines into chunks
			chunkCount := int(math.Ceil(float64(s.Args.Limit) / float64(CHUNK_SIZE)))
			logrus.Debugf("chunkcount: %d %s%s", chunkCount, prefix, year)

			for i := 0; i < chunkCount; i++ {
				go func(offset int, prefix string) {
					lb := (offset * CHUNK_SIZE) + 1
					ub := ((offset + 1) * CHUNK_SIZE)
					if ub > s.Args.Limit {
						ub = s.Args.Limit
					}

					isSkipping := false
					notFoundStreak := 0

					for suffix := lb; suffix <= ub; suffix++ {
						nim := fmt.Sprintf("%s%03d", prefix, suffix)

						// If skippable (more than SKIP_THRESHOLD)
						if isSkipping {
							s.Failed <- nim
							continue
						}

						student, err := s.GetByNIM(nim)
						logrus.Debugf("stud: %s err: %s", student, err)

						if err != nil {
							logrus.Warnf("Failed to fetch %s, reason: %s", nim, err)
							s.Failed <- nim

							notFoundStreak++
							if notFoundStreak > SKIP_THRESHOLD {
								isSkipping = true
								logrus.Warnf("Skipping %s - %s", nim, fmt.Sprintf("%s%03d", prefix, ub))
							}
							continue
						}
						notFoundStreak = 0
						s.Students <- student
					}

					logrus.Infof("Fetched %s - %s", fmt.Sprintf("%s%03d", prefix, lb), fmt.Sprintf("%s%03d", prefix, ub))
					progress.m.Lock()
					progress.count += ub - lb + 1
					progress.m.Unlock()
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

func (s TeamsScrapper) StartTeams() {
	progress := Progress{count: 0, limit: s.Args.Limit * len(s.Args.Years.Arr) * len(s.Args.Prefixes.Arr)}

	for _, prefix := range s.Args.Prefixes.Arr {
		for _, year := range s.Args.Years.Arr {

			nimPrefix := fmt.Sprintf("%s%s", prefix, year)

			// Split goroutines into chunks
			chunkCount := int(math.Ceil(float64(s.Args.Limit) / float64(CHUNK_SIZE)))
			logrus.Debugf("chunkcount: %d %s%s", chunkCount, prefix, year)

			for i := 0; i < chunkCount; i++ {
				go func(offset int, prefix string) {
					lb := (offset * CHUNK_SIZE) + 1
					ub := ((offset + 1) * CHUNK_SIZE)
					if ub > s.Args.Limit {
						ub = s.Args.Limit
					}

					isSkipping := false
					notFoundStreak := 0

					for suffix := lb; suffix <= ub; suffix++ {
						nim := fmt.Sprintf("%s%03d", prefix, suffix)

						// If skippable (more than SKIP_THRESHOLD)
						if isSkipping {
							s.Failed <- nim
							continue
						}

						student, err := s.GetByNIMTeams(nim)
						logrus.Debugf("stud: %s err: %s", student, err)

						if err != nil {
							logrus.Warnf("Failed to fetch %s, reason: %s", nim, err)
							s.Failed <- nim

							notFoundStreak++
							if notFoundStreak > SKIP_THRESHOLD {
								isSkipping = true
								logrus.Warnf("Skipping %s - %s", nim, fmt.Sprintf("%s%03d", prefix, ub))
							}
							continue
						}
						notFoundStreak = 0
						s.Students <- student
					}

					logrus.Infof("Fetched %s - %s", fmt.Sprintf("%s%03d", prefix, lb), fmt.Sprintf("%s%03d", prefix, ub))
					progress.m.Lock()
					progress.count += ub - lb + 1
					progress.m.Unlock()
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
