package produce

import (
    "time"
    "os"
	"github.com/Shopify/sarama"
    "github.com/hpcloud/tail"
    "log-shiper/tool"
	
)



func WriteToChannel(c chan sarama.ProducerMessage, filename string, listentip string, topic string){
	// 读文件
	t, err := tail.TailFile(filename, tail.Config{
        Follow:     true,
        ReOpen:    true,
        Poll:      true,
        Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
    })
    if err != nil {
        tool.Logger.Error(err.Error())
        os.Exit(-1)
    }

    for line := range t.Lines{
        var t = time.Now()
        key := listentip + "_" + t.Format("20060102150405")
        msg := &sarama.ProducerMessage{
            Topic: topic,
            Key: sarama.StringEncoder(line.Text),
            Value: sarama.StringEncoder(key),
        }
        c <- *msg
    }

}