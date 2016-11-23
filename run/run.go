package run

import (
	"fmt"
	"github.com/goplog/configparser"
	"log"
)

func Run(cfg string, debug bool) {
	run(cfg, debug)
}

func run(cfg string, debug bool) {
	parser := configparser.New()
	parser.ReadFile(cfg)
	sections := parser.Sections()
	for _, key := range sections {
		fmt.Println(key)
		options, err := parser.Options(key)
		if err != nil {
			log.Fatal(err)
		}
		for _, value := range options {
			fmt.Println(parser.Get(key, value))
		}
	}
}
