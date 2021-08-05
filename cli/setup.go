package cli

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/akamensky/argparse"
	"github.com/mkamadeus/nicscraper/models"
	"github.com/mkamadeus/nicscraper/utils/constants"
	"github.com/sirupsen/logrus"
)

func (a App) Setup() (*models.Arguments, error) {
	// Scraping parameters
	argPrefixes := a.Parser.String("p", "prefix", &argparse.Options{Required: true, Help: "Prefix of major/faculty (e.g: \"135\", \"165,182\")"})
	argYears := a.Parser.String("y", "year", &argparse.Options{Required: true, Help: "Year with format of YY (e.g: \"18\", \"19,20\")"})
	limit := a.Parser.Int("l", "limit", &argparse.Options{Help: "Set scraping limit (default)", Default: 20})

	// Output parameters
	format := a.Parser.Selector("f", "format", []string{"json", "csv"}, &argparse.Options{Help: "Output file format", Default: "json"})
	filename := a.Parser.String("o", "output", &argparse.Options{Help: "Output filename", Default: fmt.Sprintf("%d.json", time.Now().Unix())})
	argToken := a.Parser.String("t", "token", &argparse.Options{Required: false, Help: "NIC session token"})
	verbose := a.Parser.Flag("v", "verbose", &argparse.Options{Help: "Verbose output", Default: false})

	envToken := os.Getenv("NIC_CI_TOKEN")

	err := a.Parser.Parse(os.Args)
	if err != nil {
		return nil, err
	}

	if *verbose {
		logrus.SetLevel(logrus.DebugLevel)
	}

	var prefixes []string
	if *argPrefixes == "ALL" {
		prefixes = constants.TPBCodes[:]
	} else {
		prefixes = strings.Split(strings.Replace(*argPrefixes, " ", "", -1), ",")
	}

	years := strings.Split(strings.Replace(*argYears, " ", "", -1), ",")

	var token string
	if envToken != "" {
		token = envToken
	} else if *argToken != "" {
		token = *argToken
	} else {
		return nil, errors.New("token not found")
	}

	args := &models.Arguments{
		Token:          token,
		Prefixes:       prefixes,
		Years:          years,
		Limit:          *limit,
		Format:         *format,
		OutputFilename: *filename,
		Verbose:        *verbose,
	}
	logrus.Debugf("Arguments: %s", args)

	return args, nil
}
