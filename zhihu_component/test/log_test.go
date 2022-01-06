package test

import (
	"errors"
	"github.com/huage66/zhihu_go/zhihu_component/logger"
	"runtime"
	"testing"
)

func TestLog(t *testing.T) {
	logger.Use(logger.Config{
		Level: "info",
		Path: "./log/log.txt",
	})
	AddError(errors.New("test"))
}

func AddError(err interface{}) {
	pc, file, line, boolRes := runtime.Caller(1)
	if !boolRes {
		logger.Error("AddInfo error: runtime.Caller(2)")
		return
	}
	f := runtime.FuncForPC(pc)
	logger.ErrorF("{System} {caller=%v-line:%v} {func=%v} {error=%v}", file, line, f.Name(), err)
}
