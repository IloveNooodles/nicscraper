package main

import (
	"time"

	"github.com/alexflint/go-arg"
	"github.com/mkamadeus/nicscraper/models"
	"github.com/mkamadeus/nicscraper/scraper"
	"github.com/mkamadeus/nicscraper/utils/file"
	"github.com/sirupsen/logrus"
)

func main() {
	var args models.Arguments
	arg.MustParse(&args)

	if args.Verbose {
		logrus.SetLevel(logrus.DebugLevel)
	}
	logrus.Debugln(args)

	// Setup timer
	start := time.Now()

	s, err := scraper.New(args)
	if err != nil {
		logrus.Fatalf("Error: %s", err.Error())
		return
	}

	// Start scraping
	logrus.Infoln("Starting to scrape...")
	s.Start()
	data, failed := s.Receive()

	// Stop timer
	elapsed := time.Since(start)

	logrus.Infof("Time elapsed: %.2fs", elapsed.Seconds())
	logrus.Infof("Failed to fetch: %s", failed)

	// Output filename
	if s.Args.Format == "json" {
		file.OutputJSON(s.Args.OutputFilename, data)
	} else if s.Args.Format == "csv" {
		file.OutputCSV(s.Args.OutputFilename, data)
	}

}
