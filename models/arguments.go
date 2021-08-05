package models

type Arguments struct {
	Token          string
	Prefixes       []string
	Years          []string
	Limit          int
	Format         string
	OutputFilename string
	Verbose        bool
}
