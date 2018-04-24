package produce

import (
    "time"
    "os"
	"github.com/Shopify/sarama"
    "github.com/hpcloud/tail"
    "github.com/log-shiper/tool"	
)

func WriteToChannel(c chan sarama.ProducerMessage,  file string, ip string, topic string, port string){
	t, err := tail.TailFile(file, tail.Config{
        Follow:    true,
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
        key := ip + ":" + port + "_" + t.Format("2006-01-02T15:04:05Z07:00")
        msg := &sarama.ProducerMessage{
            Topic: topic,
            Key: sarama.StringEncoder(key),
            Value: sarama.StringEncoder(line.Text),
        }
        c <- *msg
    }
}
