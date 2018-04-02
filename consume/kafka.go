package cousume

import (
    "time"
    "strings"
	"os"
	"fmt"
    "github.com/Shopify/sarama"
    "log-shiper/tool"
)


func WriteToKafka(c chan sarama.ProducerMessage, brokers string){
    config := sarama.NewConfig()
    config.Producer.RequiredAcks = sarama.WaitForAll
    config.Producer.Retry.Max = 5
    config.Producer.Return.Successes = true
    config.Producer.Timeout = 5 * time.Second
    ips := strings.Split(brokers, ",")
    p, err := sarama.NewSyncProducer(ips, config)
    if err != nil {
        tool.Logger.Error(err.Error())
        os.Exit(-1)
    }
    defer func() {
        if err := p.Close(); err != nil {
            tool.Logger.Error(err.Error())
            os.Exit(-1)
        }
    }()
    for {
        if msg, ok := <- c; ok{
            partition, offset, err := p.SendMessage(&msg)
            if err != nil {
                tool.Logger.Error(err.Error())
                os.Exit(-1)
			}
			msg := fmt.Sprintf("Msg is stored in parttion,offset: %d %d", partition, offset)
            tool.Logger.Info(msg)
        }else{
            tool.Logger.Error("read data form channel error.")
            os.Exit(-1)
        }
    }
}