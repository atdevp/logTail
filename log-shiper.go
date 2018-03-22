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
    filename  = kingpin.Arg("filename", "log file").Required().String()
    topic     = kingpin.Arg("topic", "topic name").Required().String()
    brokers   = kingpin.Arg("brokers", "kafka brokers").Required().String()
    listentip = kingpin.Arg("listentip", "listen ip address").Required().String()
    httpport  = kingpin.Arg("httpport", "http port").Required().String()
)

func init(){
    kingpin.Parse()
    tool.PathExist(*filename)
}

func main(){
    c := make(chan sarama.ProducerMessage)
    go produce.WriteToChannel(c, *filename, *listentip, *topic )
    go consume.WriteToKakc(c, *brokers)
    go consume.WriteToKakc(c, *brokers)
    httpserver.Start(*httpport)
}