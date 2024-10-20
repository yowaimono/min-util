package minutil

import (
	"runtime"
)

// RecoverPanic 恢复panic并记录日志
func RecoverPanic() {
	if err := recover(); err != nil {
		Warn("Recovered from panic: %v\n%s", err, GetStackInfo())
	}
}

// GetStackInfo 获取Panic堆栈信息
func GetStackInfo() string {
	buf := make([]byte, 4096)
	n := runtime.Stack(buf, false)
	return string(buf[:n])
}
