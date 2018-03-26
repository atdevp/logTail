package tool

import (
	"github.com/op/go-logging"
	"fmt"
)

var Logger = logging.MustGetLogger("LogShiper")
var format = logging.MustStringFormatter(
    `%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
)


// 判断字符串是否为空
func Argument(s map[string]string) (string, bool){
	for k, v := range s{
		fmt.Println(k, v)
		if (v == ""){
			return k, false
		}
	}
	return "", true
}


