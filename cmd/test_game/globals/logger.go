package globals

import nblogger "github.com/banaconda/nb-logger"

var Logger nblogger.Logger

func InitLogger(path string, level int, bufferSize int, flags int) error {
	var err error
	Logger, err = nblogger.NewLogger(path, level, bufferSize, flags)
	return err
}
