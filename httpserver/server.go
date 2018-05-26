package httpserver

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/log-shiper/g"
)

type Monitor struct {
	StartTime time.Time
	Data      g.SystemInfo
}

func (m *Monitor) Start(ch chan string, port string) {
	http.HandleFunc("/monitor", func(w http.ResponseWriter, r *http.Request) {
		m.Data.RunTime = time.Now().Sub(m.StartTime).String()
		m.Data.Delay = len(ch)

		ret, _ := json.MarshalIndent(m.Data, "", "\t")
		io.WriteString(w, string(ret))
	})
	var socket = ":" + port
	http.ListenAndServe(socket, nil)

}
