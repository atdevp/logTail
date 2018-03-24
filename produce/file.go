package produce

import (
    "time"
    "os"
	"github.com/Shopify/sarama"
    "github.com/hpcloud/tail"
    "log-shiper/tool"
    "strings"
    "fmt"
	
)

func WriteToKafka(brokers string, filename string, listentip string, topic string){
    config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
    config.Producer.Return.Successes = true
	config.Producer.Timeout = 5 * time.Second
	ips := strings.Split(brokers, ",")
    P, err := sarama.NewSyncProducer(ips, config)
    if err != nil {
        tool.Logger.Error(err.Error())
        os.Exit(-1)
	}
	defer func() {
		if err := P.Close(); err != nil {
            tool.Logger.Error(err.Error())
            os.Exit(-1)
		}
    }()
    
	// 读文件
	t, err := tail.TailFile(filename, tail.Config{
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
        key := listentip + "_" + t.Format("2006-01-02T15:04:05Z07:00")
        msg := &sarama.ProducerMessage{
            Topic: topic,
            Key: sarama.StringEncoder(key),
            Value: sarama.StringEncoder(line.Text),
        }
        // c <- *msg
        partition, offset, err := P.SendMessage(msg)
        if err != nil {
            tool.Logger.Error(err.Error())
            tool.Logger.Error("Send msg to kfk error.")
            os.Exit(-1)
        }
        sf := fmt.Sprintf("Msg send ok and partition,offset is %d,%d", partition ,offset)
        tool.Logger.Info(sf)
    }

}