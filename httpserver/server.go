package httpserver

import (
    "net/http"
    "fmt"
)

func HsttpServer(w http.ResponseWriter, r *http.Request){
    fmt.Fprintln(w, "<h1>hello world</h1>")
}

func Start(port string){

    mux := http.NewServeMux()
    mux.HandleFunc("/", HsttpServer)
    var socket = ":" + port
    http.ListenAndServe(socket, mux) 
}