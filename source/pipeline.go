package source

import (
	"bufio"
	"fmt"
	"github.com/goplog/configparser"
	"log"
	"os"
)

func PipeYieldField(out chan<- string, parser *configparser.ConfigParser) {
	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadSlice('\n')
		if err != nil {
			log.Fatal("Read error")
		}
		out <- fmt.Sprintf("%s", line)
	}
}
