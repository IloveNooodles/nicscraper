package models

import (
	"strings"

	"github.com/mkamadeus/nicscraper/utils/constants"
)

type PrefixesSeparated struct {
	Arr []string
}

type YearsSeparated struct {
	Arr []string
}

type Arguments struct {
	Prefixes       PrefixesSeparated `arg:"-p,--prefix,required" help:"Prefix of major/faculty (e.g: \"135\", \"165,182\"), can specify \"ALL\" to scrape all registered"`
	Years          YearsSeparated    `arg:"-y,--years,required" help:"Year with format of YY (e.g: \"18\", \"19,20\")"`
	Limit          int               `arg:"-l,--limit,required" help:"Set scraping limit"`
	UseTeams       bool              `arg:"-u,--use-teams" default:"False" help:"Scrapping use teams, Warning: Must provide jwt-token and Cvid"`
	Jwt            string            `arg:"-j,--jwt" default:"" help:"Set jwt token"`
	Cvid           string            `arg:"-c,--cvid" default:"" help:"Set cvid token"`
	Token          string            `arg:"-t,--token,env:NIC_CI_TOKEN" help:"Set token via argument, also accepts from NIC_CI_TOKEN environment variable"`
	Format         string            `arg:"-f,--format" default:"json" help:""`
	OutputFilename string            `arg:"-o,--output" default:"result.json"`
	Verbose        bool              `arg:"-v,--verbose"`
}

func (Arguments) Description() string {
	return "Scrapes student data from nic.itb.ac.id. Made with <3"
}

func (p *PrefixesSeparated) UnmarshalText(b []byte) error {
	var prefixes []string

	if string(b) == "ALL" {
		prefixes = constants.TPBCodes[:]
	} else {
		prefixes = strings.Split(strings.Replace(string(b), " ", "", -1), ",")
	}

	p.Arr = prefixes
	return nil
}

func (y *YearsSeparated) UnmarshalText(b []byte) error {
	years := strings.Split(strings.Replace(string(b), " ", "", -1), ",")
	y.Arr = years
	return nil
}
