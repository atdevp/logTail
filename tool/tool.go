package tool

import (
	"os"
	"github.com/op/go-logging"
)

var Logger = logging.MustGetLogger("LogShiper")
var format = logging.MustStringFormatter(
    `%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
)
// 判断文件是否存在
func PathExist(_path string) {
	_, err := os.Stat(_path)
	if err != nil && err.Error() == os.ErrNotExist.Error() {
		Logger.Error("%s is not exists.", _path)
		return
	}
	return
}


