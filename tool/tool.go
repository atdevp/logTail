package tool

import (
	"fmt"
)


func Argument(s map[string]string) (string, bool){
	for k, v := range s{
		fmt.Println(k, v)
		if (v == ""){
			return k, false
		}
	}
	return "", true
}


