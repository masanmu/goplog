package sink

import (
	"fmt"
	"github.com/goplog/configparser"
)

func CalculateItem(out <-chan string, parser *configparser.ConfigParser) {
	for v := range out {
		fmt.Println(v)
	}
}
