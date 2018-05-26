package g

type MsgKey struct {
	Addr string
	Port int64
}

type SystemInfo struct {
	Delay   int    `json:"delay"`
	RunTime string `json:"runTime"`
}
