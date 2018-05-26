package cousume

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/Shopify/sarama"
	"github.com/log-shiper/g"
)

type WriteToKafka struct {
	Brokers string
	Topic   string
	MsgKey  g.MsgKey
}

func (w *WriteToKafka) Write(channel chan string) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForLocal
	config.Producer.Retry.Max = 5
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	brokers := strings.Split(w.Brokers, ",")
	client, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		log.Print("create producer fail. reason is %s.", err.Error())
	}
	defer func() {
		if err := client.Close(); err != nil {
			log.Print("close producer fail. reason is %s.", err.Error())
		}
	}()

	for {
		if line, ok := <-channel; ok {
			t := time.Now()
			key := w.MsgKey.Addr + ":" + w.MsgKey.Port + "_" + t.Format("2006-01-02T15:04:05Z07:00")

			msg := sarama.ProducerMessage{
				Topic: w.Topic,
				Value: sarama.StringEncoder(line),
				Key:   sarama.StringEncoder(key),
			}
			partition, offset, err := client.SendMessage(&msg)
			if err != nil {
				log.Print("send msg to kafka fail. reason is %s.", err.Error())
			}
			logMsg := fmt.Sprintf("Msg is stored in parttion,offset: %d %d", partition, offset)
			log.Print(logMsg)
		} else {
			log.Print("read data form channel error.")
		}
	}

}
