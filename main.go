package main

import (
	"fmt"
	"github.com/arahmandanu/sinau_go_craft/cmd"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

func main() {
	time.Local = time.UTC
	registerStackDumpReceiver()
	cmd.Execute()
}

func registerStackDumpReceiver() {
	sigChan := make(chan os.Signal)
	go func() {
		stacktrace := make([]byte, 32768)
		for range sigChan {
			length := runtime.Stack(stacktrace, true)
			fmt.Println(length)
		}
	}()
	signal.Notify(sigChan, syscall.SIGQUIT)
}
