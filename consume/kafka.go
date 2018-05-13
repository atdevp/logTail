package cousume

import (
    "strings"
	"os"
	"fmt"
    "github.com/Shopify/sarama"
    "github.com/log-shiper/tool"
)


func WriteToKafka(c chan sarama.ProducerMessage, brokers string){
    config := sarama.NewConfig()
    config.Producer.RequiredAcks = sarama.WaitForLocal
    config.Producer.Retry.Max = 5
    config.Producer.Partitioner = sarama.NewRandomPartitioner
    config.Producer.Return.Successes = true
    // 用来推接视频push发送消息。 联系人： 视频 亚军 & 客户端 孙友军
    // config.Net.SASL.Enable = true
    // config.Net.SASL.User = "push"
    // config.Net.SASL.Password = "push_pwd"
    ips := strings.Split(brokers, ",")
    p, err := sarama.NewSyncProducer(ips, config)
    if err != nil {
        tool.Logger.Error(err.Error())
        // os.Exit(-1)
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
                // os.Exit(-1)
			}
			msg := fmt.Sprintf("Msg is stored in parttion,offset: %d %d", partition, offset)
            tool.Logger.Info(msg)
        }else{
            tool.Logger.Error("read data form channel error.")
            // os.Exit(-1)
        }
    }
}