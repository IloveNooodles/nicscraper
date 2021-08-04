package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"time"

	"github.com/akamensky/argparse"
	"github.com/mkamadeus/nicscraper/scraper"
	"github.com/mkamadeus/nicscraper/utils/file"
)

func main() {
	token := os.Getenv("NIC_CI_TOKEN")
	if token == "" {
		log.Fatalln("Error: NIC_CI_TOKEN is not set.")
		return
	}

	parser := argparse.NewParser("nicscraper", "Scrapes from NIC")

	// Scraping parameters
	prefixes := parser.List("p", "prefix", &argparse.Options{Required: true, Help: "Prefix of major/faculty (e.g: 135, 165)"})
	years := parser.List("y", "year", &argparse.Options{Required: true, Help: "Year with format of YY (e.g: 18, 19, 20)"})
	limit := parser.Int("l", "limit", &argparse.Options{Help: "Set scraping limit (default)", Default: 20})

	// Output parameters
	format := parser.Selector("f", "format", []string{"json", "csv"}, &argparse.Options{Help: "Output file format", Default: "json"})
	filename := parser.String("o", "output", &argparse.Options{Help: "Output filename", Default: fmt.Sprintf("%d.json", time.Now().Unix())})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		return
	}

	start := time.Now()

	s := scraper.New(token, *limit, false)

	// Test connection
	if !s.IsConnected() {
		log.Fatalln("Error: Not connected to site. Your token may be expired.")
		return
	} else {
		log.Println("Connected. Starting...")
	}

	s.Start(*prefixes, *years)
	data, errors := s.Receive()

	elapsed := time.Since(start)
	log.Printf("Time elapsed: %.2fs", elapsed.Seconds())

	log.Printf("Failed to fetch: %s", errors)

	sort.Slice(data[:], func(i, j int) bool {
		return data[i].Username < data[j].Username
	})

	if *format == "json" {
		file.OutputJSON(*filename, data)
	} else if *format == "csv" {
		file.OutputCSV(*filename, data)
	}

}
