package channel

import (
	"encoding/json"
	"github.com/goplog/configparser"
	"log"
)

func ParseLine(out chan<- interface{}, in <-chan interface{}, parser *configparser.ConfigParser) {
	var f interface{}
	key, err := parser.Get("channel", "channel_keys")
	if err != nil {
		log.Fatalf("No Key in Define")
	}
	for v := range in {
		err := json.Unmarshal([]byte(v.(string)), &f)
		if err != nil {
			log.Fatal("Error json")
		}
		line := f.(map[string]interface{})
		out <- line[key]
	}
	close(out)
}
