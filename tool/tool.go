package tool

import (
	"os"
	"github.com/op/go-logging"
)

var Logger = logging.MustGetLogger("LogShiper")

// 判断文件是否存在
func PathExist(_path string) bool {
	_, err := os.Stat(_path)
	if err != nil && err.Error() == os.ErrNotExist.Error() {
		return false
	}
	return true
}


