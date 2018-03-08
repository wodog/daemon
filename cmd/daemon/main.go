package main

import (
	"os"

	"github.com/wodog/daemon/fork"
)

func main() {
	fork.DaemonFork(os.Args[1:])
}
