package produce

import (
	"log"

	"github.com/hpcloud/tail"
)

type ReadFromFile struct {
	Path string
}

func (r *ReadFromFile) Read(channel chan string) {
	config := tail.Config{
		Follow:   true,
		ReOpen:   true,
		Poll:     true,
		Location: &tail.SeekInfo{Offset: 0, Whence: 2},
	}
	t, err := tail.TailFile(r.Path, config)
	if err != nil {
		log.Print("read file fail. reason is %s.", err.Error())
	}
	for line := range t.Lines {
		channel <- line.Text
	}
}
