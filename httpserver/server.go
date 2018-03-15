package httpserver

import (
    "net/http"
    "fmt"
)

// 定义HttpServer
func HsttpServer(w http.ResponseWriter, r *http.Request){
    fmt.Fprintln(w, "<h1>hello world</h1>")
}


// 启动HttpServer
func Start(httpport string){

    mux := http.NewServeMux()
    mux.HandleFunc("/", HsttpServer)
    var port = ":" + httpport
    http.ListenAndServe(port, mux) 
}