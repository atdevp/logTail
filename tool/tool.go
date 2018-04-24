package tool

import (
	"github.com/op/go-logging"
	"fmt"
)

var Logger = logging.MustGetLogger("LogShiper")

func Argument(s map[string]string) (string, bool){
	for k, v := range s{
		fmt.Println(k, v)
		if (v == ""){
			return k, false
		}
	}
	return "", true
}


