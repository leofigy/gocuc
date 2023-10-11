package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	for i := 0; i < 200; i++ {
		go func() {
			time.Sleep(time.Second)
		}()
	}
	fmt.Println(runtime.NumGoroutine())
	fmt.Println(runtime.NumCPU()) // P = 8
}
