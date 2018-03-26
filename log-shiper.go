package main

import (
    "flag"
    "log-shiper/produce"
    "log-shiper/httpserver"
    "log-shiper/tool"
    "os"
    "fmt"
)


var (
    h   bool
    f   string
    t   string
    b   string
    a   string
    p   string
)
func usage(){
    fmt.Fprintf(os.Stderr, `Version: log-shiper/1.0.0
Usage: log-shiper  [-f filename] [-t topic] [-b brokers] [-a ip] [-p port] -[h]         
Options:`)
    flag.PrintDefaults()
    os.Exit(-1)
}

func init(){
    flag.BoolVar(&h, "h", false, "this help")
    flag.StringVar(&a, "a", "127.0.0.1", "log agent ip address")
    flag.StringVar(&b, "b", "", "kafka broker address")
    flag.StringVar(&f, "f", "", "log file name")
    flag.StringVar(&p, "p", "", "log agent port")
    flag.StringVar(&t, "t", "", "topic name")
    flag.Usage = usage
}

func main(){
    flag.Parse()
    if h {
        flag.Usage()
    }
    arg := map[string]string{
        "broker" : b,
        "file": f,
        "ip": a,
        "topic": t,
        "port": p,
    }
    n, ret := tool.Argument(arg)
    fmt.Println(ret)
    fmt.Println(ret)
    fmt.Println(ret)
    if !ret {
        tool.Logger.Error("%s is null", n)
        flag.Usage()
    }
    go produce.WriteToKafka(b, f, a, t )
    httpserver.Start(p)
}