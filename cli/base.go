package cli

import (
	"github.com/akamensky/argparse"
)

type App struct {
	Parser *argparse.Parser
}

func New() *App {
	parser := argparse.NewParser("nicscraper", "Scrapes student data from nic.itb.ac.id. Make sure that NIC_CI_TOKEN is set via environment variable or using -t/--token")

	return &App{
		Parser: parser,
	}
}
