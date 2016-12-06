package channel

import (
	"github.com/goplog/configparser"
)

func ParseLine(out chan<- string, in <-chan string, parser *configparser.ConfigParser) {
	for v := range in {
		out <- v
	}
}
