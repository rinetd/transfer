package main

import (
	"os"

	"github.com/rinetd/transfer/cmd"
)

func main() {
	if err := cmd.Run(); err != nil {
		os.Exit(0)
	}
}
