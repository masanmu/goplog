package source

import (
	"github.com/goplog/common"
	"github.com/goplog/configparser"
	"github.com/hpcloud/tail"
	"log"
)

func FielYieldLine(out chan<- interface{}, parser *configparser.ConfigParser) {
	var startPosition int
	fileName, err := parser.Get("source", "source_file")
	if err != nil {
		log.Fatal(err)
	}

	if ok := common.IsFile(fileName); !ok {
		log.Fatalf("File %s not found", fileName)
	}

	startPosition, err = parser.Getint("source", "start_position")
	if err != nil {
		startPosition = 0
	}
	lines, err := tail.TailFile(fileName, tail.Config{Follow: true, Location: &tail.SeekInfo{Offset: 0, Whence: startPosition}, Poll: true, ReOpen: true})
	if err != nil {
		log.Fatalf("Failed to open %s file", fileName)
	}
	for line := range lines.Lines {
		out <- line.Text
	}
	close(out)
}
