package run

import (
	"github.com/goplog/channel"
	"github.com/goplog/configparser"
	"github.com/goplog/sink"
	"github.com/goplog/source"
	"log"
)

func Run(cfg string, debug bool) {
	run(cfg, debug)
}

func run(cfg string, debug bool) {
	var Out = make(chan interface{})
	var In = make(chan interface{})
	var sourceMap = map[string]func(chan<- interface{}, *configparser.ConfigParser){"file": source.FielYieldLine, "pipe": source.PipeYieldField}
	var channelMap = map[string]func(chan<- interface{}, <-chan interface{}, *configparser.ConfigParser){"json": channel.ParseLine}
	var sinkMap = map[string]func(<-chan interface{}, *configparser.ConfigParser){"zabbix": sink.CalculateItem}

	parser := configparser.New()
	parser.ReadFile(cfg)

	sourceModuleName, err := parser.Get("source", "source_module")
	if err != nil {
		log.Fatal(err)
	}
	channelModuleName, err := parser.Get("channel", "channel_module")
	if err != nil {
		log.Fatal(err)
	}
	sinkModuleName, err := parser.Get("sink", "sink_module")
	if err != nil {
		log.Fatal(err)
	}

	go sourceMap[sourceModuleName](In, parser)
	go channelMap[channelModuleName](Out, In, parser)
	sinkMap[sinkModuleName](Out, parser)
	//	select {}
}
