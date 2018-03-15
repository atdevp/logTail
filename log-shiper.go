package main

import (
    kingpin "gopkg.in/alecthomas/kingpin.v2"
    "github.com/Shopify/sarama"
    consume "log-shiper/consume"
    "log-shiper/produce"
    "log-shiper/httpserver"
    "log-shiper/tool"
)


var (
	filename = kingpin.Arg("filename", "log file").Required().String()
	topic  = kingpin.Arg("topic", "topic name").Required().String()
    brokers = kingpin.Arg("brokers", "kafka brokers").Required().String()
    httpport = kingpin.Arg("httpport", "http port").Required().String()
    listentip = kingpin.Arg("listentip", "listen ip address").Required().String()
)

func init(){
    kingpin.Parse()
    tool.PathExist(*filename)
    go httpserver.Start(*httpport)
}

func main(){
    c := make(chan sarama.ProducerMessage, 1000)
    go produce.WriteToChannel(c, *filename, *listentip, *topic )
    consume.WriteToKakc(c, *brokers)
}