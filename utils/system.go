package utils

import (
	"bytes"
	"runtime"
	"strconv"
)

// getGID gets the current goroutine ID (copied from https://blog.sgmansfield.com/2015/12/goroutine-ids/)
func GetGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}
