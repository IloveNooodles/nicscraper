package main

import (
	"strings"
	"time"

	"github.com/alexflint/go-arg"
	"github.com/mkamadeus/nicscraper/models"
	"github.com/mkamadeus/nicscraper/scraper"
	"github.com/mkamadeus/nicscraper/utils/file"
	"github.com/sirupsen/logrus"
)

var start time.Time

func main() {
	var args models.Arguments
	res := arg.MustParse(&args)

	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		DisableColors: false,
	})

	if args.Verbose {
		logrus.SetLevel(logrus.DebugLevel)
	}

	logrus.Debugln(args)

	// Setup timer
	start = time.Now()

	method := strings.ToLower(args.Connection)

	if method == "teams" {
		useTeams(args)
	} else if method == "nic" {
		useNic(args)
	} else {
		res.Fail("Error: Invalid connection. Available: nic/teams")
	}
}

func useNic(args models.Arguments) {
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

func useTeams(args models.Arguments) {

	s, err := scraper.NewTeams(args)

	if err != nil {
		logrus.Fatalf("Error: %s", err.Error())
		return
	}

	// Start scraping
	logrus.Infoln("Starting to scrape...")

	s.StartTeams()
	data, failed := s.ReceiveTeams()

	// Stop timer
	elapsed := time.Since(start)

	logrus.Infof("Time elapsed: %.2fs", elapsed.Seconds())
	logrus.Infof("Failed to fetch: %s", failed)

	// Output filename
	if s.Args.Format == "json" {
		file.OutputTeamsJSON(s.Args.OutputFilename, data)
	} else if s.Args.Format == "csv" {
		file.OutputTeamsCSV(s.Args.OutputFilename, data)
	}
}
