package g

type MsgKey struct {
	Addr string
	Port string
}

type SystemInfo struct {
	Delay   int    `json:"delay"`
	RunTime string `json:"runTime"`
}
