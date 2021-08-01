package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/akamensky/argparse"
	"github.com/mkamadeus/nicscraper/scraper"
)

func main() {
	start := time.Now()

	parser := argparse.NewParser("nicscraper", "Scrapes from NIC")
	limit := parser.Int("l", "limit", &argparse.Options{Help: "Set scraping limit (default)", Default: 20})
	prefix := parser.String("p", "prefix", &argparse.Options{Required: true, Help: "Prefix of major/faculty (e.g: 135, 165)"})
	year := parser.String("y", "year", &argparse.Options{Required: true, Help: "Year with format of YY (e.g: 18, 19, 20)"})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		return
	}

	token := os.Getenv("NIC_CI_TOKEN")
	if token == "" {
		panic("no token found")
	}

	s := scraper.New(token, false)

	for i := 1; i <= *limit; i++ {
		nim := fmt.Sprintf("%s%s%03d", *prefix, *year, i)
		go func(nim string) {
			s.GetByNIM(nim)
		}(nim)
	}
	s.Receive(*limit)

	elapsed := time.Since(start)
	log.Printf("Time elapsed: %.2fs", elapsed.Seconds())

}
