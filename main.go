package main

import (
	"flag"
	"fmt"
	"github.com/goplog/run"
	"log"
	"os"
	"strings"
)

func main() {
	var cfg string
	fs := flag.NewFlagSet("goplog", flag.ExitOnError)
	fs.Usage = func() {
		fmt.Println(Help())
		os.Exit(0)
	}

	fs.StringVar(&cfg, "c", "cfg.confg", "configuration file")
	fs.Parse(os.Args[1:])

	f, err := os.Open(cfg)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()
	run.Run(cfg, false)
}

func Help() string {
	helpText := `
Usage: goplog [options]
Options:
    -c      configusration file default:cfg.conf
`
	return strings.TrimSpace(helpText)
}
