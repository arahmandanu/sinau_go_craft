package main

import (
	"github.com/arahmandanu/sinau_go_craft/cmd"
	"time"
)

func main() {
	time.Local = time.UTC
	cmd.Execute()
}
