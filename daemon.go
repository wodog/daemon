package daemon

import (
	"os"

	"github.com/wodog/daemon/fork"
)

func init() {
	if os.Getppid() != 1 {
		fork.DaemonFork(os.Args)
	}
}
