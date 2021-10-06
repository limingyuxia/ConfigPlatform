package util

import (
	"context"
	"log"
)

// LogWriter 实现io.Writer接口，将字节流输出到log
type LogWriter struct {
	Ctx context.Context
}

// io.Writer接口实现
func (w *LogWriter) Write(p []byte) (n int, err error) {
	sz := len(p)
	if sz > 0 && p[sz-1] == '\n' {
		p = p[:sz-1]
	}

	log.Print("[sql] ", string(p))
	return sz, nil
}
